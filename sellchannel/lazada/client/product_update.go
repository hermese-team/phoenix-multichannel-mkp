package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// UpdateProduct calls POST /product/update with the XML payload (the caller
// builds the <Request><Product>… document, identifying the product by ItemId).
func (c *Client) UpdateProduct(ctx context.Context, payloadXML string) (*UpdateProductResponse, error) {
	_, env, err := c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/product/update", map[string]string{"payload": payloadXML})
	if err != nil {
		return nil, err
	}
	// Best-effort decode: the caller may ignore the body, so a successful update
	// must not look failed just because `data` had an unexpected shape.
	var resp UpdateProductResponse
	_ = json.Unmarshal(env.Data, &resp)
	return &resp, nil
}
