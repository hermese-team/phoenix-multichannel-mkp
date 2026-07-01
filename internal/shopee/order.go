package shopee

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const (
	pathGetOrderList   = "/api/v2/order/get_order_list"
	pathGetOrderDetail = "/api/v2/order/get_order_detail"
)

// OrderAPI exposes Shopee order-related API calls.
type OrderAPI struct {
	client *Client
}

// NewOrderAPI creates an OrderAPI backed by the given Client.
func NewOrderAPI(c *Client) *OrderAPI {
	return &OrderAPI{client: c}
}

// GetOrderList fetches a page of order SNs using cursor-based pagination.
// Returns the list, next cursor, and whether there are more pages.
func (a *OrderAPI) GetOrderList(ctx context.Context, req GetOrderListRequest) ([]OrderBrief, string, bool, error) {
	params := url.Values{}
	params.Set("time_range_field", req.TimeRangeField)
	params.Set("time_from", strconv.FormatInt(req.TimeFrom, 10))
	params.Set("time_to", strconv.FormatInt(req.TimeTo, 10))
	params.Set("page_size", strconv.Itoa(req.PageSize))
	if req.Cursor != "" {
		params.Set("cursor", req.Cursor)
	}
	if req.OrderStatus != "" {
		params.Set("order_status", req.OrderStatus)
	}

	var resp GetOrderListResponse
	if err := a.client.get(ctx, pathGetOrderList, params, &resp); err != nil {
		return nil, "", false, err
	}
	if resp.IsTokenExpired() {
		return nil, "", false, fmt.Errorf("shopee: access token expired — trigger refresh")
	}
	if resp.IsError() {
		return nil, "", false, fmt.Errorf("shopee get_order_list: %s – %s (request_id=%s)",
			resp.Error, resp.Message, resp.RequestID)
	}
	if resp.Response == nil {
		return nil, "", false, nil
	}
	return resp.Response.OrderList, resp.Response.NextCursor, resp.Response.More, nil
}

// GetOrderDetail fetches full order details for up to 50 order SNs at once.
// Shopee allows max 50 per call — caller must chunk larger lists.
func (a *OrderAPI) GetOrderDetail(ctx context.Context, orderSNs []string, optionalFields ...string) ([]Order, error) {
	if len(orderSNs) == 0 {
		return nil, nil
	}
	if len(orderSNs) > 50 {
		return nil, fmt.Errorf("shopee: get_order_detail accepts max 50 SNs, got %d", len(orderSNs))
	}

	params := url.Values{}
	params.Set("order_sn_list", strings.Join(orderSNs, ","))
	if len(optionalFields) > 0 {
		params.Set("response_optional_fields", strings.Join(optionalFields, ","))
	}

	var resp GetOrderDetailResponse
	if err := a.client.get(ctx, pathGetOrderDetail, params, &resp); err != nil {
		return nil, err
	}
	if resp.IsTokenExpired() {
		return nil, fmt.Errorf("shopee: access token expired — trigger refresh")
	}
	if resp.IsError() {
		return nil, fmt.Errorf("shopee get_order_detail: %s – %s (request_id=%s)",
			resp.Error, resp.Message, resp.RequestID)
	}
	if resp.Response == nil {
		return nil, nil
	}
	return resp.Response.OrderList, nil
}

// chunk splits a slice into sub-slices of at most size n.
func chunk(s []string, n int) [][]string {
	var out [][]string
	for len(s) > 0 {
		if len(s) < n {
			n = len(s)
		}
		out = append(out, s[:n])
		s = s[n:]
	}
	return out
}

// GetOrderDetailAll fetches full detail for any number of SNs, chunking as needed.
func (a *OrderAPI) GetOrderDetailAll(ctx context.Context, orderSNs []string, optionalFields ...string) ([]Order, error) {
	var all []Order
	for _, batch := range chunk(orderSNs, 50) {
		orders, err := a.GetOrderDetail(ctx, batch, optionalFields...)
		if err != nil {
			return all, err
		}
		all = append(all, orders...)
	}
	return all, nil
}
