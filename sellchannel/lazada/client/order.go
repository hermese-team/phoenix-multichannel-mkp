package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// orderPageLimit is the page size requested from /orders/get.
const orderPageLimit = 100

// GetOrderList calls GET /orders/get for one page. The window bounds are
// formatted as UTC "yyyy-MM-ddTHH:mm:ssZ". Returns the total count and the
// order numbers on this page.
func (c *Client) GetOrderList(ctx context.Context, updateAfter, updateBefore time.Time, offset int) (int, []string, error) {
	raw, err := c.doSigned(ctx, "/orders/get", map[string]string{
		"limit":         strconv.Itoa(orderPageLimit),
		"offset":        strconv.Itoa(offset),
		"update_after":  updateAfter.UTC().Format("2006-01-02T15:04:05Z"),
		"update_before": updateBefore.UTC().Format("2006-01-02T15:04:05Z"),
	})
	if err != nil {
		return 0, nil, err
	}
	var data orderListDTO
	if err := json.Unmarshal(raw, &data); err != nil {
		return 0, nil, fmt.Errorf("decode order list: %w", err)
	}
	numbers := make([]string, len(data.Orders))
	for i, o := range data.Orders {
		numbers[i] = strconv.FormatInt(o.OrderNumber, 10)
	}
	return data.CountTotal, numbers, nil
}

// GetOrderDetail calls GET /order/get.
func (c *Client) GetOrderDetail(ctx context.Context, orderID string) (*OrderDetail, error) {
	raw, err := c.doSigned(ctx, "/order/get", map[string]string{"order_id": orderID})
	if err != nil {
		return nil, err
	}
	var dto orderDetailDTO
	if err := json.Unmarshal(raw, &dto); err != nil {
		return nil, fmt.Errorf("decode order detail: %w", err)
	}
	return dto.toDomain(), nil
}

// GetOrderItems calls GET /order/items/get.
func (c *Client) GetOrderItems(ctx context.Context, orderID string) ([]OrderItem, error) {
	raw, err := c.doSigned(ctx, "/order/items/get", map[string]string{"order_id": orderID})
	if err != nil {
		return nil, err
	}
	var dtos []orderItemDTO
	if err := json.Unmarshal(raw, &dtos); err != nil {
		return nil, fmt.Errorf("decode order items: %w", err)
	}
	items := make([]OrderItem, len(dtos))
	for i, d := range dtos {
		items[i] = d.toDomain()
	}
	return items, nil
}
