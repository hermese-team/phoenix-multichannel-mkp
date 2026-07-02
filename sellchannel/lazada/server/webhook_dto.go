package server

// OrderWebhookRequest is the Lazada order-status push payload. Verified against
// a real Lazada callback sample (UA Apache-HttpClient):
//
//	{"seller_id":"9999","message_type":0,"timestamp":1782979023,"site":"lazada_sg",
//	 "data":{"order_status":"...","trade_order_id":"123456",
//	         "trade_order_line_id":"12345","status_update_time":1782979023}}
//
// Note: seller_id is a STRING and the order fields are nested under "data".
type OrderWebhookRequest struct {
	SellerID    string           `json:"seller_id"`
	MessageType int              `json:"message_type"`
	Timestamp   int64            `json:"timestamp"`
	Site        string           `json:"site"`
	Data        OrderWebhookData `json:"data"`
}

type OrderWebhookData struct {
	OrderStatus      string `json:"order_status"`
	TradeOrderID     string `json:"trade_order_id"`
	TradeOrderLineID string `json:"trade_order_line_id"`
	StatusUpdateTime int64  `json:"status_update_time"`
}
