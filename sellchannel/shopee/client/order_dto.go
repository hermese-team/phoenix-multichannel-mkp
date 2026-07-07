package client

type GetOrderDetailResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Response  struct {
		OrderList []OrderDetail `json:"order_list"`
	} `json:"response"`
}

type OrderDetail struct {
	OrderSN          string      `json:"order_sn"`
	OrderStatus      string      `json:"order_status"`
	ShopID           int64       `json:"shop_id"`
	BuyerUsername    string      `json:"buyer_username"`
	TotalAmount      float64     `json:"total_amount"`
	Currency         string      `json:"currency"`
	PaymentMethod    string      `json:"payment_method"`
	CreateTime       int64       `json:"create_time"`
	UpdateTime       int64       `json:"update_time"`
	ShipByDate       int64       `json:"ship_by_date"`
	ItemList         []OrderItem `json:"item_list"`
	RecipientAddress struct {
		Name        string `json:"name"`
		Phone       string `json:"phone"`
		Town        string `json:"town"`
		District    string `json:"district"`
		City        string `json:"city"`
		State       string `json:"state"`
		Region      string `json:"region"`
		Zipcode     string `json:"zipcode"`
		FullAddress string `json:"full_address"`
	} `json:"recipient_address"`
}

type GetOrderListResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Response  struct {
		OrderList  []OrderListItem `json:"order_list"`
		More       bool            `json:"more"`
		NextCursor string          `json:"next_cursor"`
	} `json:"response"`
}

type OrderListItem struct {
	OrderSN     string `json:"order_sn"`
	OrderStatus string `json:"order_status"`
}

type OrderItem struct {
	ItemID        int64  `json:"item_id"`
	ItemName      string `json:"item_name"`
	ItemSku       string `json:"item_sku"`
	ModelID       int64  `json:"model_id"`
	ModelName     string `json:"model_name"`
	ModelSku      string `json:"model_sku"`
	ModelQuantity int    `json:"model_quantity_purchased"`
	ModelPrice    float64 `json:"model_discounted_price"`
}
