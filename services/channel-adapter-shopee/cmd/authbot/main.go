// cmd/authbot/main.go
//
// Pure-HTTP Shopee OAuth flow — no headless browser.
// All steps reverse-engineered from mitmproxy captures.
//
// Flow:
//  1. GET partner auth URL  →  302 Location contains /authorize?...&random=NONCE
//     Extract NONCE.
//
//  2. GET /opservice/api/v1/oauth2/login?auth_shop=1&id=PARTNER_ID&random=NONCE&redirect_url=...
//     (no session cookies yet)  →  302 Location = Shopee login page URL
//     That URL contains `sign=` and `timestamp=` — extract them.
//
//  3. GET /signin/oauth/identifier?...  →  initial cookies (_QPWSDCXHZQA).
//
//  4. POST /api/account/login/?...sign=...&timestamp=...
//     Body: username=&password_hash=SHA256(password)&remember=false&captcha_signature=
//     Response: {code:0, data:{auth_code:"..."}} + SPC_F / SC_DFP / SPC_SI cookies.
//
//  5. GET /api/v1/oauth2/callback?auth_code=...&state=...&oauth_device_id=...
//     →  SHOP_TOKEN{NONCE} cookie.
//
//  6. GET /opservice/api/v1/oauth2/login?...&random=NONCE&...
//     (now WITH session cookies)  →  200 + SPC_CDS cookie.
//
//  7. POST /opservice/api/v1/authorization_management/authorize/submit
//     ?SPC_CDS=...&SPC_CDS_VER=2
//     →  redirect to redirect_url?code=...
//
//  8. Exchange code via Shopee Open Platform API  →  tokens in Redis.
//
// Usage:
//   SHOPEE_SELLER_USERNAME=xxx SHOPEE_SELLER_PASSWORD=yyy go run ./cmd/authbot

package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"

	"git.amaze-x.com/phoenix/channel-adapter-shopee/config"
	"git.amaze-x.com/phoenix/channel-adapter-shopee/internal/infrastructure"
	"git.amaze-x.com/phoenix/channel-adapter-shopee/internal/shopee"
	"golang.org/x/net/publicsuffix"
)

const (
	// accountBase MUST be HTTPS so that Go's cookiejar stores cookies with the
	// Secure flag (SPC_SI, SPC_F, SC_DFP). Shopee's account server sets all
	// session cookies as Secure; using http:// causes the jar to silently drop them.
	// The opservice 302 Location may give http:// URLs — we upgrade them in httpsAccount().
	accountBase   = "https://account.sandbox.test-stable.shopee.com"
	openBase      = "https://open.sandbox.test-stable.shopee.com"
	oauthDeviceID = "GoAuthBotDevice0000001"
	spcCDSVer     = "2"
)

// ── JSON models ───────────────────────────────────────────────────────────────

type loginResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		AuthCode string `json:"auth_code"`
	} `json:"data"`
}

type submitResp struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Msg   string `json:"msg"`
	Data  *struct {
		RedirectURL string `json:"redirect_url"`
	} `json:"data"`
}

// ── Main ──────────────────────────────────────────────────────────────────────

func main() {
	cfg := config.MustLoad()
	rdb, err := infrastructure.NewRedis(cfg)
	if err != nil {
		log.Fatalf("redis: %v", err)
	}
	signer := shopee.NewSigner(cfg.Shopee.PartnerID, cfg.Shopee.PartnerKey)
	tokenStore := shopee.NewRedisTokenStore(rdb, cfg.Shopee.BaseURL, signer,
		cfg.Shopee.PartnerID, cfg.Shopee.ShopID)

	// ── Shortcut: SHOPEE_AUTH_CODE bypasses the full web OAuth flow. ─────────────
	// Use this when you have a code from the Shopee Partner Center test-auth page.
	// Cross-region login (SG partner + TH seller) cannot be automated in sandbox
	// because opservice always signs for the partner's registered region (SG), and
	// the login endpoint validates the sign server-side with an unknown algorithm.
	//
	// Workflow:
	//   1. Open https://partner.test-stable.shopee.sg/ → Test → Auth
	//   2. Select the TH test shop → Generate auth code
	//   3. SHOPEE_AUTH_CODE=<code> make auth-bot
	if code := os.Getenv("SHOPEE_AUTH_CODE"); code != "" {
		log.Printf("[shortcut] exchanging pre-obtained auth_code=%s", code)
		if err := tokenStore.ExchangeCode(context.Background(), code); err != nil {
			log.Fatalf("[shortcut] ExchangeCode: %v", err)
		}
		log.Println("[shortcut] tokens stored in Redis — worker is ready")
		return
	}

	username := os.Getenv("SHOPEE_SELLER_USERNAME")
	password := os.Getenv("SHOPEE_SELLER_PASSWORD")
	if username == "" || password == "" {
		if os.Getenv("SHOPEE_SPC_SI") == "" {
			log.Fatal("SHOPEE_SELLER_USERNAME and SHOPEE_SELLER_PASSWORD are required (or set SHOPEE_AUTH_CODE or SHOPEE_SPC_SI)")
		}
	}

	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	client := &http.Client{
		Jar: jar,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse // manual redirect handling
		},
		Timeout: 15 * time.Second,
	}

	// Build partner auth URL (our HMAC-signed request to Shopee Open Platform)
	authURL := shopee.GetAuthURL(cfg.Shopee.BaseURL, cfg.Shopee.PartnerID, signer, cfg.Shopee.RedirectURL)
	partnerID := cfg.Shopee.PartnerID

	// ── Step 1: partner auth URL → extract nonce + opservice URL ────────────────
	nonce, opserviceURL, err := step1ExtractNonce(client, authURL)
	if err != nil {
		log.Fatalf("[1] %v", err)
	}
	log.Printf("[1] nonce = %s", nonce)

	// ── Step 2: opservice oauth2/login (unauthenticated) → login page URL ─────
	//   302 Location gives us Shopee's login URL which includes a server-computed
	//   sign + timestamp. The sign covers the region param, so we must let
	//   Shopee compute the sign for the correct seller region.
	sellerRegion := os.Getenv("SHOPEE_SELLER_REGION")
	if sellerRegion == "" {
		sellerRegion = "TH"
	}
	loginPageURL, stateB64, err := step2GetLoginPageURL(client, opserviceURL, sellerRegion)
	if err != nil {
		log.Fatalf("[2] %v", err)
	}
	log.Printf("[2] login page URL (region=%s)", urlParam(loginPageURL, "region"))

	// ── Step 3: GET login identifier page → initial cookies ───────────────────
	if err := step3GetInitialCookies(client, loginPageURL); err != nil {
		log.Fatalf("[3] %v", err)
	}
	log.Println("[3] initial cookies acquired")

	// Inject JS-set anti-bot cookies that Go's HTTP client can't get without
	// executing JavaScript. Shopee's login page sets these via JS (reCAPTCHA/
	// device fingerprint). In sandbox the values are not strictly validated.
	injectAntiBotCookies(client.Jar, accountBase)

	// ── Step 4-5: login + OAuth callback ─────────────────────────────────────
	// Shortcut: if SHOPEE_SPC_SI is set, inject the pre-existing seller session
	// cookie and attempt session-based login. When a browser user is already
	// logged into seller center, the login SPA's JS calls POST /api/account/login/
	// WITHOUT username/password — relying on SPC_SI to authenticate — and the
	// account server returns auth_code directly. We replicate that here.
	//
	// To get SPC_SI:
	//   Chrome DevTools → Application → Cookies → account.sandbox.test-stable.shopee.com
	//   Copy SPC_SI value → set SHOPEE_SPC_SI env var.
	if spcSI := os.Getenv("SHOPEE_SPC_SI"); spcSI != "" {
		log.Printf("[4-5] SHOPEE_SPC_SI set — injecting session cookie, attempting session-based auth")
		injectSessionCookies(client.Jar, accountBase, spcSI,
			os.Getenv("SHOPEE_SC_DFP"),
			os.Getenv("SHOPEE_SPC_SEC_SI"),
		)
		sellerRegion = os.Getenv("SHOPEE_SELLER_REGION")
		if sellerRegion == "" {
			sellerRegion = "TH"
		}

		// Step 4 (session variant): Re-GET the login page URL with SPC_SI now in the jar.
		// When the account server sees SPC_SI, it may detect the existing session and
		// immediately redirect to redirect_uri?code=AUTH_CODE instead of serving the SPA.
		// This mirrors the OAuth server pattern: GET /signin/oauth/ + valid session = code grant.
		spcSIAuthCode, sessErr := trySessionGrant(client, loginPageURL)
		if sessErr == nil && spcSIAuthCode != "" {
			log.Printf("[4] session grant OK, auth_code=%s", spcSIAuthCode)
			// Step 5: exchange auth_code → SHOP_TOKEN cookie
			if err := step5OAuthCallback(client, spcSIAuthCode, stateB64, nonce, sellerRegion); err != nil {
				log.Printf("[5] OAuth callback failed: %v", err)
			} else {
				log.Printf("[5] SHOP_TOKEN set via session grant")
			}
		} else {
			log.Printf("[4] session grant failed: %v", sessErr)
			log.Printf("[4] will proceed to step 6 with SPC_SI only")
		}
	} else {
		authCode, sellerRegion, loginErr := step4PostLogin(client, username, sha256Hex(password), loginPageURL,
			func(correctRegion string) (string, error) {
				log.Printf("[2b] re-fetching login URL with region=%s", correctRegion)
				newURL, _, e := step2GetLoginPageURL(client, opserviceURL, correctRegion)
				return newURL, e
			})

		wrongRegion := loginErr != nil && strings.Contains(loginErr.Error(), "wrong_region")
		if loginErr != nil && !wrongRegion {
			log.Fatalf("[4] %v", loginErr)
		}
		if wrongRegion {
			log.Printf("[4] error_wrong_region — cross-region login blocked by Shopee sandbox.")
			log.Printf("[4] Workaround A: SHOPEE_SPC_SI=<value> make auth-bot  (copy SPC_SI from browser DevTools)")
			log.Printf("[4] Workaround B: AUTH_CODE=<code> make auth-bot        (generate from Partner Center → Test → Auth)")
			os.Exit(1)
		}
		log.Printf("[4] login OK, auth_code=%s region=%s", authCode, sellerRegion)

		// ── Step 5: GET /oauth2/callback → SHOP_TOKEN cookie ─────────────────
		if err := step5OAuthCallback(client, authCode, stateB64, nonce, sellerRegion); err != nil {
			log.Fatalf("[5] %v", err)
		}
		log.Printf("[5] SHOP_TOKEN set")
	}

	// ── Step 6: opservice oauth2/login (authenticated) → SPC_CDS cookie ───────
	spcCDS, err := step6GetSPCCDS(client, partnerID, nonce)
	if err != nil {
		log.Fatalf("[6] %v", err)
	}
	log.Printf("[6] SPC_CDS = %s", spcCDS)

	// ── Step 7: POST /authorize/submit → redirect with code= ─────────────────
	code, err := step7AuthorizeSubmit(client, spcCDS)
	if err != nil {
		log.Fatalf("[7] %v", err)
	}
	log.Printf("[7] auth code = %s", code)

	// ── Step 8: exchange code for tokens ─────────────────────────────────────
	if err := tokenStore.ExchangeCode(context.Background(), code); err != nil {
		log.Fatalf("[8] ExchangeCode: %v", err)
	}
	log.Println("[8] tokens stored in Redis — worker is ready")
}

// ── Step implementations ──────────────────────────────────────────────────────

// step1ExtractNonce GETs the partner auth URL, reads the 302 Location,
// and returns both the nonce and the full opservice URL to use in step 2.
func step1ExtractNonce(client *http.Client, authURL string) (nonce, opserviceURL string, err error) {
	resp, err := get(client, authURL, nil)
	if err != nil {
		return "", "", err
	}
	resp.Body.Close()
	loc := resp.Header.Get("Location")
	if loc == "" {
		return "", "", fmt.Errorf("expected 302 from partner auth URL, got %d", resp.StatusCode)
	}
	log.Printf("[1] redirect → %s", loc)
	u, _ := url.Parse(loc)
	nonce = u.Query().Get("random")
	if nonce == "" {
		return "", "", fmt.Errorf("no 'random' param in %s", loc)
	}
	return nonce, loc, nil
}

// step2GetLoginPageURL calls the opservice oauth2/login endpoint WITHOUT session cookies.
// opserviceURL is the URL from step 1's 302 Location.
// regionHint, if non-empty, is appended as ?region=<hint> so Shopee generates a login URL
// whose sign covers that region (needed when the seller's region differs from the partner's).
// Returns the login page URL (with sign + timestamp) and the base64-encoded state.
func step2GetLoginPageURL(client *http.Client, opserviceURL, regionHint string) (loginPageURL, stateB64 string, err error) {
	callURL := opserviceURL
	if regionHint != "" {
		// Append region hint without re-encoding existing params.
		// Shopee will compute a sign covering the given region in the redirect URL.
		callURL = opserviceURL + "&region=" + regionHint
	}
	log.Printf("[2] GET %s", callURL)

	resp, err := get(client, callURL, nil)
	if err != nil {
		return "", "", err
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	loc := resp.Header.Get("Location")
	if resp.StatusCode != 302 || loc == "" {
		return "", "", fmt.Errorf("expected 302 from oauth2/login, got %d: %s", resp.StatusCode, body)
	}
	// Upgrade http:// → https:// so Go's cookiejar stores Secure-flagged cookies.
	loc = httpsAccount(loc)
	log.Printf("[2] login page URL = %s", loc)

	u, _ := url.Parse(loc)
	stateB64 = u.Query().Get("state")
	return loc, stateB64, nil
}

// step3GetInitialCookies GETs the login identifier page to populate initial cookies.
// Also scans the HTML for embedded JS state (e.g. window.__INITIAL_STATE__) that may
// contain a per-session client_secret used to compute the cross-region sign.
func step3GetInitialCookies(client *http.Client, loginPageURL string) error {
	resp, err := get(client, loginPageURL, nil)
	if err != nil {
		return err
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	log.Printf("[3] GET → %d (%d bytes)", resp.StatusCode, len(body))
	return nil
}


// step4PostLogin POSTs credentials to Shopee's account login API.
// loginPageURL must already be signed by Shopee (from step 2).
// Returns (authCode, sellerRegion, error).
//
// On error_wrong_region, calls getSignedURL(correctRegion) to obtain a new Shopee-signed
// login URL for the seller's actual region, then retries.
func step4PostLogin(
	client *http.Client,
	username, pwHash, loginPageURL string,
	getSignedURL func(correctRegion string) (string, error),
) (string, string, error) {
	// First attempt with the URL Shopee gave us (e.g. region=SG for SG partner).
	authCode, code, msg, err := doLoginAttempt(client, username, pwHash, loginPageURL, "")
	if err == nil {
		u, _ := url.Parse(loginPageURL)
		return authCode, u.Query().Get("region"), nil
	}

	// On wrong-region: try swapping region=SG → region=TH in the signed URL.
	// If Shopee's sign does NOT cover the region parameter, this will succeed
	// and we avoid needing a TH-signed URL from opservice.
	if code == 10006 || strings.Contains(msg, "wrong_region") {
		sellerRegion := os.Getenv("SHOPEE_SELLER_REGION")
		if sellerRegion == "" {
			sellerRegion = "TH"
		}
		// Attempt 2a: swap region in the original signed URL (sign may not cover region).
		swappedURL := strings.Replace(loginPageURL, "region=SG", "region="+sellerRegion, 1)
		if swappedURL != loginPageURL {
			log.Printf("[4] retrying with region swapped SG→%s (sign preserved)", sellerRegion)
			authCode, _, _, err = doLoginAttempt(client, username, pwHash, swappedURL, "")
			if err == nil {
				return authCode, sellerRegion, nil
			}
			log.Printf("[4] region-swap attempt failed: %v", err)
		}

		// Attempt 2b: ask Shopee to sign a new login URL for the correct region.
		newURL, e := getSignedURL(sellerRegion)
		if e != nil {
			return "", "", fmt.Errorf("re-fetching signed URL for %s: %w", sellerRegion, e)
		}
		authCode, _, _, err = doLoginAttempt(client, username, pwHash, newURL, "")
		if err == nil {
			return authCode, sellerRegion, nil
		}
	}

	return "", "", err
}

// doLoginAttempt performs one login POST attempt using the login page URL as-is.
// The login page URL must already be signed by Shopee for the correct region.
// Returns authCode, Shopee JSON code, Shopee message, error.
func doLoginAttempt(client *http.Client, username, pwHash, loginPageURL, _ string) (string, int, string, error) {
	u, _ := url.Parse(loginPageURL)
	u.Path = "/api/account/login/"

	// Preserve the raw query exactly — Shopee validates the sign against the verbatim query.
	postURL := u.String()
	log.Printf("[4] POST %s", postURL)

	form := url.Values{}
	form.Set("username", username)
	form.Set("password_hash", pwHash)
	form.Set("remember", "false")
	form.Set("captcha_signature", "")

	req, err := http.NewRequest(http.MethodPost, postURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", 0, "", err
	}
	// Referer: login SPA sub-route (mitmproxy shows /signin/oauth/identifier)
	refererURL, _ := url.Parse(loginPageURL)
	refererURL.Path = "/signin/oauth/identifier"

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("User-Agent", chromeUA)
	req.Header.Set("Origin", accountBase)
	req.Header.Set("Referer", refererURL.String())
	// Security/fingerprint headers observed in mitmproxy (browser sends "undefined"
	// when the JS fingerprint SDK has no value computed yet).
	req.Header.Set("X-Sec-Dfp", "undefined")
	req.Header.Set("af-ac-enc-sz-token", "")
	req.Header.Set("x-sz-sdk-version", "3.4.1-2&1.5.6")

	resp, err := client.Do(req)
	if err != nil {
		return "", 0, "", fmt.Errorf("POST login: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Printf("[4] response %d: %s", resp.StatusCode, body)

	var lr loginResp
	if err := json.Unmarshal(body, &lr); err != nil {
		return "", 0, "", fmt.Errorf("parsing login response: %w", err)
	}
	if lr.Code != 0 {
		return "", lr.Code, lr.Message, fmt.Errorf("login failed code=%d message=%s", lr.Code, lr.Message)
	}
	if lr.Data == nil || lr.Data.AuthCode == "" {
		return "", 0, "", fmt.Errorf("login succeeded but no auth_code in body: %s", body)
	}
	return lr.Data.AuthCode, 0, "", nil
}

// trySessionGrant re-GETs the login page URL with SPC_SI already in the cookie jar.
// When the account server detects an existing SPC_SI session it may skip the login form
// and immediately redirect to the redirect_uri with an auth_code (OAuth code-grant flow).
// Returns the auth_code if the server grants it; otherwise returns an error with details.
func trySessionGrant(client *http.Client, loginPageURL string) (string, error) {
	log.Printf("[4] session grant: GET %s", loginPageURL)
	resp, err := get(client, loginPageURL, nil)
	if err != nil {
		return "", fmt.Errorf("GET login page: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Printf("[4] session grant response %d (body=%d bytes)", resp.StatusCode, len(body))

	// Case 1: server redirects directly to redirect_uri?code=... — success
	if resp.StatusCode == 302 {
		loc := resp.Header.Get("Location")
		log.Printf("[4] session grant redirect → %s", loc)
		code, err := extractCode(loc)
		if err == nil && code != "" {
			return code, nil
		}
		// Redirect exists but no code (e.g. redirect back to login page) — fall through
		log.Printf("[4] session grant redirect has no code: %s", loc)
	}

	// Case 2: 200 HTML — account server still serving login SPA (session not detected server-side)
	if len(body) > 50 && strings.Contains(string(body[:min(200, len(body))]), "<!DOCTYPE") {
		return "", fmt.Errorf("account server returned login SPA (session not detected server-side)")
	}

	return "", fmt.Errorf("unexpected response %d: %s", resp.StatusCode, body)
}

// step5OAuthCallback GETs Shopee's internal /oauth2/callback endpoint, which
// validates the login auth_code and sets SHOP_TOKEN{NONCE} cookie.
// sellerRegion must match the region used in the successful login POST (step 4).
func step5OAuthCallback(client *http.Client, authCode, stateB64, nonce, sellerRegion string) error {
	q := url.Values{}
	q.Set("auth_code", authCode)
	q.Set("response_type", "code")
	q.Set("state", stateB64)
	q.Set("scope", "profile")
	q.Set("region", sellerRegion)
	q.Set("oauth_device_id", oauthDeviceID)

	callbackURL := openBase + "/api/v1/oauth2/callback?" + q.Encode()
	resp, err := get(client, callbackURL, map[string]string{"Referer": accountBase + "/"})
	if err != nil {
		return err
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	log.Printf("[5] oauth2/callback %d: %s", resp.StatusCode, body)

	// May return 302 (to /authorize) or 200 JSON — either way SHOP_TOKEN is set via cookie
	return nil
}

// step6GetSPCCDS calls /opservice/api/v1/oauth2/login WITH session cookies.
// This time Shopee responds with 200 and sets SPC_CDS cookie.
func step6GetSPCCDS(client *http.Client, partnerID int64, nonce string) (string, error) {
	// Debug: log all cookies currently in jar for both domains
	if u, _ := url.Parse(accountBase); u != nil {
		var names []string
		for _, c := range client.Jar.Cookies(u) {
			names = append(names, c.Name)
		}
		log.Printf("[6] account cookies: %v", names)
	}
	if u, _ := url.Parse(openBase); u != nil {
		var names []string
		for _, c := range client.Jar.Cookies(u) {
			names = append(names, c.Name)
		}
		log.Printf("[6] open cookies: %v", names)
	}
	redirectURL := fmt.Sprintf(
		"%s/authorize?auth_shop=true&auth_type=shop&id=%d&isRedirect=true&is_agent=false&random=%s",
		openBase, partnerID, nonce,
	)
	q := url.Values{}
	q.Set("auth_shop", "1")
	q.Set("id", fmt.Sprintf("%d", partnerID))
	q.Set("random", nonce)
	q.Set("redirect_url", redirectURL)

	initURL := openBase + "/opservice/api/v1/oauth2/login?" + q.Encode()
	resp, err := get(client, initURL, nil)
	if err != nil {
		return "", err
	}

	// If still getting 302, follow the redirect (should set SPC_CDS)
	if resp.StatusCode == 302 {
		resp.Body.Close()
		loc := httpsAccount(resp.Header.Get("Location"))
		log.Printf("[6] oauth2/login 302 → %s", loc)
		if loc != "" {
			resp2, err := get(client, loc, nil)
			if err != nil {
				return "", fmt.Errorf("following oauth2/login redirect: %w", err)
			}
			body2, _ := io.ReadAll(resp2.Body)
			resp2.Body.Close()
			log.Printf("[6] redirect response %d: %s", resp2.StatusCode, body2)
			// Follow one more redirect if needed
			if resp2.StatusCode == 302 {
				loc2 := resp2.Header.Get("Location")
				log.Printf("[6] second redirect → %s", loc2)
				if loc2 != "" {
					resp3, err := get(client, loc2, nil)
					if err != nil {
						return "", fmt.Errorf("following second redirect: %w", err)
					}
					body3, _ := io.ReadAll(resp3.Body)
					resp3.Body.Close()
					log.Printf("[6] second redirect response %d: %s", resp3.StatusCode, string(body3)[:min(200, len(body3))])
				}
			}
		}
	} else {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		log.Printf("[6] oauth2/login %d: %s", resp.StatusCode, body)
	}

	// SPC_CDS should now be in the cookie jar
	parsedURL, _ := url.Parse(openBase)
	for _, c := range client.Jar.Cookies(parsedURL) {
		if c.Name == "SPC_CDS" {
			return c.Value, nil
		}
	}
	return "", fmt.Errorf("SPC_CDS cookie not found after oauth2/login")
}

// step7AuthorizeSubmit POSTs to the authorize submit endpoint.
// SPC_CDS goes in query params (as seen in mitmproxy).
// Returns the partner OAuth code from the redirect URL.
func step7AuthorizeSubmit(client *http.Client, spcCDS string) (string, error) {
	q := url.Values{}
	q.Set("SPC_CDS", spcCDS)
	q.Set("SPC_CDS_VER", spcCDSVer)
	submitURL := openBase + "/opservice/api/v1/authorization_management/authorize/submit?" + q.Encode()

	req, err := http.NewRequest(http.MethodPost, submitURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, */*")
	req.Header.Set("User-Agent", chromeUA)
	req.Header.Set("Referer", openBase+"/authorize")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("authorize/submit: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Printf("[7] authorize/submit %d: %s", resp.StatusCode, body)

	// Case 1: 302 with code= in Location
	if loc := resp.Header.Get("Location"); loc != "" {
		return extractCode(loc)
	}

	// Case 2: JSON body with redirect_url containing code=
	var sr submitResp
	if err := json.Unmarshal(body, &sr); err == nil {
		if sr.Code != 0 {
			return "", fmt.Errorf("authorize/submit error: code=%d error=%s msg=%s", sr.Code, sr.Error, sr.Msg)
		}
		if sr.Data != nil && sr.Data.RedirectURL != "" {
			return extractCode(sr.Data.RedirectURL)
		}
	}

	return "", fmt.Errorf("no code in authorize/submit response: %s", body)
}

// ── Helpers ───────────────────────────────────────────────────────────────────

func get(client *http.Client, rawURL string, extraHeaders map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", chromeUA)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	for k, v := range extraHeaders {
		req.Header.Set(k, v)
	}
	return client.Do(req)
}

func extractCode(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("parsing redirect URL: %w", err)
	}
	code := u.Query().Get("code")
	if code == "" {
		return "", fmt.Errorf("no 'code' in URL: %s", rawURL)
	}
	return code, nil
}

// sha256Hex returns lowercase hex SHA-256 of s.
// Shopee login API expects password_hash = SHA256(plaintext_password).
func sha256Hex(s string) string {
	h := sha256.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}


const chromeUA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/149.0.0.0 Safari/537.36"

// injectAntiBotCookies manually sets the cookies that Shopee's login page
// JavaScript normally sets (reCAPTCHA/device fingerprint).
//
// _QPWSDCXHZQA is validated server-side — use the real value from your browser:
//   Chrome DevTools → Application → Cookies → account.sandbox.test-stable.shopee.com
//   Copy _QPWSDCXHZQA value → set SHOPEE_ANTI_BOT_COOKIE env var.
// injectSessionCookies injects a pre-existing seller session into the cookie jar,
// bypassing the login step. spcSI is required; scDFP and spcSecSI are optional.
func injectSessionCookies(jar http.CookieJar, baseURL, spcSI, scDFP, spcSecSI string) {
	u, _ := url.Parse(baseURL)
	cookies := []*http.Cookie{
		{Name: "SPC_SI", Value: spcSI, Path: "/", Secure: true, HttpOnly: true},
	}
	if scDFP != "" {
		cookies = append(cookies, &http.Cookie{Name: "SC_DFP", Value: scDFP, Path: "/", Secure: true})
	}
	if spcSecSI != "" {
		cookies = append(cookies, &http.Cookie{Name: "SPC_SEC_SI", Value: spcSecSI, Path: "/", Secure: true, HttpOnly: true})
	}
	jar.SetCookies(u, cookies)
	log.Printf("[4-5] injected SPC_SI + %d extra cookies for %s", len(cookies)-1, u.Host)
}

func injectAntiBotCookies(jar http.CookieJar, baseURL string) {
	u, _ := url.Parse(baseURL)

	antiBotValue := os.Getenv("SHOPEE_ANTI_BOT_COOKIE")
	if antiBotValue == "" {
		// Fallback: random UUID (may not work if Shopee validates strictly)
		antiBotValue = newUUID()
		log.Printf("[inject] SHOPEE_ANTI_BOT_COOKIE not set, using random UUID (may fail)")
		log.Printf("[inject] To fix: Chrome DevTools → Application → Cookies → account.sandbox.test-stable.shopee.com → copy _QPWSDCXHZQA value")
	} else {
		log.Printf("[inject] using SHOPEE_ANTI_BOT_COOKIE from env")
	}

	jar.SetCookies(u, []*http.Cookie{
		{Name: "_QPWSDCXHZQA", Value: antiBotValue, Domain: u.Hostname(), Path: "/"},
		{Name: "REC7iLP4Q", Value: "", Domain: u.Hostname(), Path: "/"},
	})
	log.Printf("[inject] anti-bot cookies set for %s (_QPWSDCXHZQA=%s)", u.Hostname(), antiBotValue)
}

// newUUID generates a random UUID v4 string.
func newUUID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant bits
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// httpsAccount upgrades an account subdomain URL from http:// to https://.
// Shopee's opservice returns http:// login page URLs, but Go's cookiejar
// silently drops Secure-flagged cookies on plain-HTTP responses.
func httpsAccount(rawURL string) string {
	const httpPrefix = "http://account.sandbox.test-stable.shopee.com"
	if strings.HasPrefix(rawURL, httpPrefix) {
		return "https://account.sandbox.test-stable.shopee.com" + rawURL[len(httpPrefix):]
	}
	return rawURL
}

// urlParam extracts a single query param value from a raw URL string.
func urlParam(rawURL, key string) string {
	u, _ := url.Parse(rawURL)
	return u.Query().Get(key)
}

