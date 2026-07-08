package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/okdev/marketplace-sync/pkg/signer"
)

// GetOrderList fetches all orders updated within [timeFrom, timeTo] (unix seconds).
// Paginates automatically using next_cursor until more=false.
func (c *Client) GetOrderList(ctx context.Context, shopID int64, accessToken string, timeFrom, timeTo int64) ([]OrderListItem, error) {
	const pageSize = 50

	var all []OrderListItem
	cursor := ""

	for {
		ts := time.Now().Unix()
		path := "/api/v2/order/get_order_list"
		sign := signer.ShopeeSign(c.cfg.PartnerID, c.cfg.AppSecret, path, ts, accessToken, shopID)

		url := fmt.Sprintf(
			"%s?partner_id=%d&timestamp=%d&sign=%s&shop_id=%d&access_token=%s&time_range_field=update_time&time_from=%d&time_to=%d&page_size=%d",
			path, c.cfg.PartnerID, ts, sign, shopID, accessToken, timeFrom, timeTo, pageSize,
		)
		if cursor != "" {
			url += "&cursor=" + cursor
		}

		var resp GetOrderListResponse
		if err := c.http.Get(ctx, url, &resp); err != nil {
			return nil, fmt.Errorf("get order list: %w", err)
		}
		if resp.Error != "" {
			return nil, fmt.Errorf("get order list: %s: %s", resp.Error, resp.Message)
		}

		all = append(all, resp.Response.OrderList...)

		if !resp.Response.More || resp.Response.NextCursor == "" {
			break
		}
		cursor = resp.Response.NextCursor
	}

	return all, nil
}

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
