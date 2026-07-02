package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// UpdateStockPrice calls POST /product/price_quantity/update with the XML
// payload (the caller builds the <Request><Product>… document). The response
// fields (code, detail, request_id) sit at the document root, not under "data".
// The parsed response is returned even on a non-"0" code so the caller can
// record per-item detail, with the error surfaced alongside.
func (c *Client) UpdateStockPrice(ctx context.Context, payloadXML string) (*UpdateStockResponse, error) {
	body, _, err := c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/product/price_quantity/update", map[string]string{"payload": payloadXML})

	var resp UpdateStockResponse
	if len(body) > 0 {
		// Best-effort decode even on error so per-item detail is available.
		_ = json.Unmarshal(body, &resp)
	}
	return &resp, err
}
