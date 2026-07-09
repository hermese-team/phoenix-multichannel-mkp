package server

// OrderWebhookRequest matches Shopee's push notification payload for order events.
// https://open.shopee.com/documents/v2/v2.push.order_status
type OrderWebhookRequest struct {
	Code      int   `json:"code"`
	Timestamp int64 `json:"timestamp"`
	ShopID    int64 `json:"shop_id"`
	Data      struct {
		OrderSN string `json:"ordersn"`
		Status  string `json:"status"`
	} `json:"data"`
}

type ProductWebhookRequest struct {
	Code      int   `json:"code"`
	Timestamp int64 `json:"timestamp"`
	ShopID    int64 `json:"shop_id"`
	Data      struct {
		ItemID int64  `json:"item_id"`
		Status string `json:"status"`
	} `json:"data"`
}

// OrderRawEvent is published to raw.accepted.shopee.v1.dev immediately after a webhook push.
// The consumer will enrich it by calling GetOrderDetail.
type OrderRawEvent struct {
	ShopID    int64  `json:"shop_id"`
	OrderSN   string `json:"order_sn"`
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Timestamp int64  `json:"timestamp"`
}
