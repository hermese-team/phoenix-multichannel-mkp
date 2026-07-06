package client

import (
	"context"
	"fmt"
	"time"

	"github.com/okdev/marketplace-sync/pkg/signer"
)

// AuthURL returns the Shopee partner authorization URL to redirect the merchant to.
func (c *Client) AuthURL(redirectURL string) string {
	ts := time.Now().Unix()
	path := "/api/v2/shop/auth_partner"
	sign := signer.ShopeeSignPublic(c.cfg.PartnerID, c.cfg.AppSecret, path, ts)
	return fmt.Sprintf("%s%s?partner_id=%d&timestamp=%d&sign=%s&redirect=%s",
		c.cfg.BaseURL, path, c.cfg.PartnerID, ts, sign, redirectURL)
}

func (c *Client) GetAccessToken(ctx context.Context, code string, shopID int64) (*GetAccessTokenResponse, error) {
	ts := time.Now().Unix()
	path := "/api/v2/auth/token/get"
	sign := signer.ShopeeSignPublic(c.cfg.PartnerID, c.cfg.AppSecret, path, ts)

	req := map[string]any{
		"code":       code,
		"shop_id":    shopID,
		"partner_id": c.cfg.PartnerID,
	}

	var resp GetAccessTokenResponse
	url := fmt.Sprintf("%s?partner_id=%d&timestamp=%d&sign=%s", path, c.cfg.PartnerID, ts, sign)
	if err := c.http.Post(ctx, url, req, &resp); err != nil {
		return nil, fmt.Errorf("get access token: %w", err)
	}
	return &resp, nil
}

func (c *Client) RefreshToken(ctx context.Context, shopID int64, refreshToken string) (*RefreshTokenResponse, error) {
	ts := time.Now().Unix()
	path := "/api/v2/auth/access_token/get"
	sign := signer.ShopeeSignPublic(c.cfg.PartnerID, c.cfg.AppSecret, path, ts)

	req := map[string]any{
		"refresh_token": refreshToken,
		"shop_id":       shopID,
		"partner_id":    c.cfg.PartnerID,
	}

	var resp RefreshTokenResponse
	url := fmt.Sprintf("%s?partner_id=%d&timestamp=%d&sign=%s", path, c.cfg.PartnerID, ts, sign)
	if err := c.http.Post(ctx, url, req, &resp); err != nil {
		return nil, fmt.Errorf("refresh token: %w", err)
	}
	return &resp, nil
}
