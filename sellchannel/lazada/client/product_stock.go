package client

import (
	"context"
	"net/http"
)

// AdjustSellableQuantity calls POST /product/stock/sellable/adjust to increase or
// decrease (delta) the sellable quantity of one or more SKUs. The caller builds
// the multi-warehouse XML payload.
func (c *Client) AdjustSellableQuantity(ctx context.Context, payloadXML string) error {
	_, _, err := c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/product/stock/sellable/adjust", map[string]string{"payload": payloadXML})
	return err
}

// UpdateSellableQuantity calls POST /product/stock/sellable/update to set the
// absolute sellable quantity of one or more SKUs.
func (c *Client) UpdateSellableQuantity(ctx context.Context, payloadXML string) error {
	_, _, err := c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/product/stock/sellable/update", map[string]string{"payload": payloadXML})
	return err
}
