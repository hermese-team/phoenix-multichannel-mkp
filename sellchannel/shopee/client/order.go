package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/okdev/marketplace-sync/pkg/signer"
)

// GetOrderDetail fetches order details from Shopee.
// accessToken must be the shop-level token stored via OAuth.
func (c *Client) GetOrderDetail(ctx context.Context, shopID int64, accessToken string, orderSNs []string) ([]OrderDetail, error) {
	ts := time.Now().Unix()
	path := "/api/v2/order/get_order_detail"
	sign := signer.ShopeeSign(c.cfg.PartnerID, c.cfg.AppSecret, path, ts, accessToken, shopID)

	// Shopee get_order_detail is a GET with all params in the query string.
	url := fmt.Sprintf("%s?partner_id=%d&timestamp=%d&sign=%s&shop_id=%d&access_token=%s&order_sn_list=%s&response_optional_fields=%s",
		path,
		c.cfg.PartnerID,
		ts,
		sign,
		shopID,
		accessToken,
		strings.Join(orderSNs, ","),
		"buyer_username,recipient_address,item_list,pay_time,ship_by_date",
	)

	var resp GetOrderDetailResponse
	if err := c.http.Get(ctx, url, &resp); err != nil {
		return nil, fmt.Errorf("get order detail: %w", err)
	}
	if resp.Error != "" {
		return nil, fmt.Errorf("shopee error %s: %s", resp.Error, resp.Message)
	}
	return resp.Response.OrderList, nil
}
