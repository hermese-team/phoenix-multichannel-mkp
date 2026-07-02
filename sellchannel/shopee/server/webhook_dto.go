package server

type OrderWebhookRequest struct {
	ShopID  int64  `json:"shop_id"`
	OrderSN string `json:"ordersn"`
	Status  string `json:"status"`
}

type ProductWebhookRequest struct {
	ShopID int64  `json:"shop_id"`
	ItemID int64  `json:"item_id"`
	Status string `json:"status"`
}
