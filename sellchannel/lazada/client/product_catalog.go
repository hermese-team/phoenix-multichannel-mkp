package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// GetProducts calls GET /products/get for one page. filter is the status filter
// (all, live, inactive, deleted, pending, rejected, sold-out); limit max 50.
// Returns the total product count and this page's products.
func (c *Client) GetProducts(ctx context.Context, filter string, offset, limit int) (int, []CatalogProduct, error) {
	raw, err := c.doSigned(ctx, "/products/get", map[string]string{
		"filter": filter,
		"offset": strconv.Itoa(offset),
		"limit":  strconv.Itoa(limit),
	})
	if err != nil {
		return 0, nil, err
	}
	var d productsDTO
	if err := json.Unmarshal(raw, &d); err != nil {
		return 0, nil, fmt.Errorf("decode products: %w", err)
	}
	return d.TotalProducts, d.Products, nil
}

// GetProductItem calls GET /product/item/get for a single product by item id.
func (c *Client) GetProductItem(ctx context.Context, itemID int64) (*CatalogProduct, error) {
	raw, err := c.doSigned(ctx, "/product/item/get", map[string]string{
		"item_id": strconv.FormatInt(itemID, 10),
	})
	if err != nil {
		return nil, err
	}
	var p CatalogProduct
	if err := json.Unmarshal(raw, &p); err != nil {
		return nil, fmt.Errorf("decode product item: %w", err)
	}
	return &p, nil
}
