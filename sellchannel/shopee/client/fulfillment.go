package client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/okdev/marketplace-sync/pkg/signer"
)

// GetShippingParameter returns which fulfillment branch the order uses:
// pickup, dropoff, or non_integrated (own-fleet).
func (c *Client) GetShippingParameter(ctx context.Context, shopID int64, accessToken, orderSN string) (*ShippingParameterResponse, error) {
	ts := time.Now().Unix()
	path := "/api/v2/logistics/get_shipping_parameter"
	sign := signer.ShopeeSign(c.cfg.PartnerID, c.cfg.AppSecret, path, ts, accessToken, shopID)

	url := fmt.Sprintf(
		"%s?partner_id=%d&timestamp=%d&sign=%s&shop_id=%d&access_token=%s&order_sn=%s",
		path, c.cfg.PartnerID, ts, sign, shopID, accessToken, orderSN,
	)

	var resp ShippingParameterResponse
	if err := c.http.Get(ctx, url, &resp); err != nil {
		return nil, fmt.Errorf("get shipping parameter: %w", err)
	}
	if resp.Error != "" {
		return nil, fmt.Errorf("get shipping parameter: %s: %s", resp.Error, resp.Message)
	}
	return &resp, nil
}

// ShipOrderNonIntegrated ships an order using own logistics (seller provides tracking number).
func (c *Client) ShipOrderNonIntegrated(ctx context.Context, shopID int64, accessToken, orderSN, trackingNumber string) error {
	ts := time.Now().Unix()
	path := "/api/v2/logistics/ship_order"
	sign := signer.ShopeeSign(c.cfg.PartnerID, c.cfg.AppSecret, path, ts, accessToken, shopID)

	body := map[string]any{
		"order_sn": orderSN,
		"non_integrated": map[string]string{
			"tracking_number": trackingNumber,
		},
	}
	url := fmt.Sprintf(
		"%s?partner_id=%d&timestamp=%d&sign=%s&shop_id=%d&access_token=%s",
		path, c.cfg.PartnerID, ts, sign, shopID, accessToken,
	)

	var resp ShipOrderResponse
	if err := c.http.Post(ctx, url, body, &resp); err != nil {
		return fmt.Errorf("ship order non-integrated: %w", err)
	}
	if resp.Error != "" {
		return fmt.Errorf("ship order non-integrated: %s: %s", resp.Error, resp.Message)
	}
	return nil
}

// ShipOrderPickup ships an order via Shopee logistics using address + timeslot pickup.
func (c *Client) ShipOrderPickup(ctx context.Context, shopID int64, accessToken, orderSN string, addressID int64, pickupTimeID string) error {
	ts := time.Now().Unix()
	path := "/api/v2/logistics/ship_order"
	sign := signer.ShopeeSign(c.cfg.PartnerID, c.cfg.AppSecret, path, ts, accessToken, shopID)

	body := map[string]any{
		"order_sn": orderSN,
		"pickup": map[string]any{
			"address_id":     addressID,
			"pickup_time_id": pickupTimeID,
		},
	}
	url := fmt.Sprintf(
		"%s?partner_id=%d&timestamp=%d&sign=%s&shop_id=%d&access_token=%s",
		path, c.cfg.PartnerID, ts, sign, shopID, accessToken,
	)

	var resp ShipOrderResponse
	if err := c.http.Post(ctx, url, body, &resp); err != nil {
		return fmt.Errorf("ship order pickup: %w", err)
	}
	if resp.Error != "" {
		return fmt.Errorf("ship order pickup: %s: %s", resp.Error, resp.Message)
	}
	return nil
}

// GetTrackingNumber fetches the tracking number assigned by Shopee after ship_order.
func (c *Client) GetTrackingNumber(ctx context.Context, shopID int64, accessToken, orderSN string) (string, error) {
	ts := time.Now().Unix()
	path := "/api/v2/logistics/get_tracking_number"
	sign := signer.ShopeeSign(c.cfg.PartnerID, c.cfg.AppSecret, path, ts, accessToken, shopID)

	url := fmt.Sprintf(
		"%s?partner_id=%d&timestamp=%d&sign=%s&shop_id=%d&access_token=%s&order_sn=%s",
		path, c.cfg.PartnerID, ts, sign, shopID, accessToken, orderSN,
	)

	var resp TrackingNumberResponse
	if err := c.http.Get(ctx, url, &resp); err != nil {
		return "", fmt.Errorf("get tracking number: %w", err)
	}
	if resp.Error != "" {
		return "", fmt.Errorf("get tracking number: %s: %s", resp.Error, resp.Message)
	}

	b, _ := json.Marshal(resp.Response)
	_ = b
	return resp.Response.TrackingNumber, nil
}
