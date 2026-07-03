package client

import (
	"context"
	"encoding/json"
	"fmt"
)

// GetCategoryTree calls GET /category/tree/get and returns the category forest.
// languageCode is optional (e.g. "en_US", "th_TH"); empty defaults to en_US.
func (c *Client) GetCategoryTree(ctx context.Context, languageCode string) ([]CategoryNode, error) {
	params := map[string]string{}
	if languageCode != "" {
		params["language_code"] = languageCode
	}
	raw, err := c.doSigned(ctx, "/category/tree/get", params)
	if err != nil {
		return nil, err
	}
	var nodes []CategoryNode
	if err := json.Unmarshal(raw, &nodes); err != nil {
		return nil, fmt.Errorf("decode category tree: %w", err)
	}
	return nodes, nil
}

// GetCategoryAttributes calls GET /category/attributes/get for a leaf category,
// returning the attributes it accepts (with IsMandatory flags).
func (c *Client) GetCategoryAttributes(ctx context.Context, primaryCategoryID, languageCode string) ([]CategoryAttribute, error) {
	params := map[string]string{"primary_category_id": primaryCategoryID}
	if languageCode != "" {
		params["language_code"] = languageCode
	}
	raw, err := c.doSigned(ctx, "/category/attributes/get", params)
	if err != nil {
		return nil, err
	}
	var attrs []CategoryAttribute
	if err := json.Unmarshal(raw, &attrs); err != nil {
		return nil, fmt.Errorf("decode category attributes: %w", err)
	}
	return attrs, nil
}
