package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/okdev/marketplace-sync/pkg/signer"
)

// GetSeller calls GET /seller/get — a minimal authenticated request used as a
// connectivity / credential smoke test. It returns the raw data blob.
func (c *Client) GetSeller(ctx context.Context) (json.RawMessage, error) {
	return c.doSigned(ctx, "/seller/get", nil)
}

// doSigned performs a signed GET against the API gateway and returns the
// envelope `data` element, erroring on a non-"0" code.
func (c *Client) doSigned(ctx context.Context, apiPath string, apiParams map[string]string) (json.RawMessage, error) {
	_, env, err := c.do(ctx, http.MethodGet, c.cfg.BaseURL, c.cfg.AccessToken, apiPath, apiParams)
	if err != nil {
		return nil, err
	}
	return env.Detail, nil
}

// do signs and performs the request against `gateway`, reads the full body,
// validates the envelope code, and returns both the raw body and the parsed
// envelope. accessToken is included in the signature only when non-empty (the
// token-refresh call signs without it). For POST the system params + sign go in
// the query string and the business params in the x-www-form-urlencoded body
// (the Lazada/LazopClient convention). A non-"0" code is surfaced as an error
// together with the body so callers can still inspect per-item detail.
func (c *Client) do(ctx context.Context, method, gateway, accessToken, apiPath string, apiParams map[string]string) ([]byte, envelope, error) {
	sys := map[string]string{
		"app_key":     c.cfg.AppKey,
		"timestamp":   strconv.FormatInt(time.Now().UnixMilli(), 10),
		"sign_method": "sha256",
	}
	if accessToken != "" {
		sys["access_token"] = accessToken
	}

	all := make(map[string]string, len(sys)+len(apiParams))
	for k, v := range sys {
		all[k] = v
	}
	for k, v := range apiParams {
		all[k] = v
	}
	signature := signer.LazadaSign(apiPath, all, c.cfg.AppSecret)

	req := c.http.R().SetContext(ctx)
	if method == http.MethodPost {
		// System params + sign in the query string, business params in the
		// x-www-form-urlencoded body (Lazada/LazopClient convention).
		req.SetQueryParams(sys)
		req.SetFormData(apiParams)
	} else {
		req.SetQueryParams(all)
	}
	req.SetQueryParam("sign", signature)

	endpoint := strings.TrimRight(gateway, "/") + apiPath
	resp, err := req.Execute(method, endpoint)
	if err != nil {
		return nil, envelope{}, err
	}

	respBody := resp.Body()
	var env envelope
	if err := json.Unmarshal(respBody, &env); err != nil {
		return nil, envelope{}, fmt.Errorf("decode envelope (%s): %w", apiPath, err)
	}
	if env.Code != "0" {
		return respBody, env, newAPIError(apiPath, env)
	}
	return respBody, env, nil
}
