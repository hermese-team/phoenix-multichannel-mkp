package shopee

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

const tracer = "shopee.client"

// Client is the base HTTP client for Shopee Open Platform API v2.
// All outbound calls are signed and traced.
type Client struct {
	baseURL    string
	httpClient *http.Client
	signer     *Signer
	shopID     int64
	tokenStore TokenStore
}

// NewClient constructs a Client. tokenStore is used to read/refresh access tokens.
func NewClient(baseURL string, timeout time.Duration, signer *Signer, shopID int64, ts TokenStore) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: timeout},
		signer:     signer,
		shopID:     shopID,
		tokenStore: ts,
	}
}

// get issues a signed GET request for a shop-level API path.
// Query params in extraParams are merged after the signature params.
func (c *Client) get(ctx context.Context, path string, extraParams url.Values, out any) error {
	return c.do(ctx, http.MethodGet, path, extraParams, nil, out)
}

// post issues a signed POST request for a shop-level API path.
func (c *Client) post(ctx context.Context, path string, body any, out any) error {
	return c.do(ctx, http.MethodPost, path, nil, body, out)
}

func (c *Client) do(ctx context.Context, method, path string, params url.Values, body any, out any) error {
	t := otel.Tracer(tracer)
	ctx, span := t.Start(ctx, fmt.Sprintf("shopee %s %s", method, path))
	defer span.End()
	span.SetAttributes(
		attribute.String("shopee.path", path),
		attribute.Int64("shopee.shop_id", c.shopID),
	)

	accessToken, err := c.tokenStore.AccessToken(ctx)
	if err != nil {
		return fmt.Errorf("getting access token: %w", err)
	}

	// Build signed query params
	q := c.signer.SignShop(path, c.shopID, accessToken)
	for k, vs := range params {
		for _, v := range vs {
			q.Add(k, v)
		}
	}

	rawURL := c.baseURL + path + "?" + q.Encode()

	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshalling request body: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, rawURL, bodyReader)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return fmt.Errorf("shopee request %s %s: %w", method, path, err)
	}
	defer resp.Body.Close()

	span.SetAttributes(attribute.Int("http.status_code", resp.StatusCode))

	if resp.StatusCode != http.StatusOK {
		raw, _ := io.ReadAll(resp.Body)
		span.SetStatus(codes.Error, fmt.Sprintf("status %d", resp.StatusCode))
		return fmt.Errorf("shopee %s %s returned %d: %s", method, path, resp.StatusCode, string(raw))
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}
	return nil
}
