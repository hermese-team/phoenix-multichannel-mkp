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
