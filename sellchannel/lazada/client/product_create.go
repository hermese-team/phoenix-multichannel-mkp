package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateProduct calls POST /product/create with the XML payload (the caller
// builds the <Request><Product>… document). Returns the new item id + sku list.
func (c *Client) CreateProduct(ctx context.Context, payloadXML string) (*CreateProductResponse, error) {
	_, env, err := c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/product/create", map[string]string{"payload": payloadXML})
	if err != nil {
		return nil, err
	}
	var resp CreateProductResponse
	if len(env.Detail) > 0 {
		if err := json.Unmarshal(env.Detail, &resp); err != nil {
			return nil, fmt.Errorf("decode create product: %w", err)
		}
	}
	return &resp, nil
}
