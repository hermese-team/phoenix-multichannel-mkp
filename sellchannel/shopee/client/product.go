package client

import (
	"context"
	"fmt"
	"time"

	"github.com/okdev/marketplace-sync/internal/domain/product"
	"github.com/okdev/marketplace-sync/pkg/signer"
)

func (c *Client) PublishProduct(ctx context.Context, shopID int64, p *product.Product) (int64, error) {
	ts := time.Now().Unix()
	path := "/api/v2/product/add_item"
	sign := signer.ShopeeSign(c.cfg.PartnerID, c.cfg.AppSecret, path, ts, "", shopID)

	req := AddItemRequest{
		ItemName:    p.Name,
		Description: p.Description,
		SellerSku:   p.SKU,
	}

	var resp AddItemResponse
	url := fmt.Sprintf("%s?partner_id=%d&timestamp=%d&sign=%s&shop_id=%d", path, c.cfg.PartnerID, ts, sign, shopID)
	if err := c.http.Post(ctx, url, req, &resp); err != nil {
		return 0, fmt.Errorf("add item: %w", err)
	}
	if resp.Error != "" {
		return 0, fmt.Errorf("shopee error: %s", resp.Message)
	}
	return resp.Response.ItemID, nil
}

//

func (c *Client) UpdateProduct(ctx context.Context, shopID int64, p *product.Product) error {
	ts := time.Now().Unix()
	path := "/api/v2/product/update_item"
	sign := signer.ShopeeSign(c.cfg.PartnerID, c.cfg.AppSecret, path, ts, "", shopID)

	req := UpdateItemRequest{
		ItemName:    p.Name,
		Description: p.Description,
		SellerSku:   p.SKU,
	}

	var resp BaseResponse
	url := fmt.Sprintf("%s?partner_id=%d&timestamp=%d&sign=%s&shop_id=%d", path, c.cfg.PartnerID, ts, sign, shopID)
	if err := c.http.Post(ctx, url, req, &resp); err != nil {
		return fmt.Errorf("update item: %w", err)
	}
	if resp.Error != "" {
		return fmt.Errorf("shopee error: %s", resp.Message)
	}
	return nil
}
