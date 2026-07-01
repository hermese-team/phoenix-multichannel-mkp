package shopee

// ---------------------------------------------------------------------------
// Common
// ---------------------------------------------------------------------------

// BaseResponse is embedded in every Shopee API response.
type BaseResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func (r BaseResponse) IsError() bool { return r.Error != "" && r.Error != "error_auth_token_expired" }
func (r BaseResponse) IsTokenExpired() bool { return r.Error == "error_auth_token_expired" }

// ---------------------------------------------------------------------------
// Auth
// ---------------------------------------------------------------------------

type GetAccessTokenRequest struct {
	Code      string `json:"code"`
	ShopID    int64  `json:"shop_id"`
	PartnerID int64  `json:"partner_id"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
	ShopID       int64  `json:"shop_id"`
	PartnerID    int64  `json:"partner_id"`
}

type TokenResponse struct {
	BaseResponse
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn     int64  `json:"expire_in"` // seconds from now
	ShopID       int64  `json:"shop_id"`
	PartnerID    int64  `json:"partner_id"`
}

// ---------------------------------------------------------------------------
// Order — get_order_list
// ---------------------------------------------------------------------------

type GetOrderListRequest struct {
	TimeRangeField string `json:"time_range_field"` // "create_time" | "update_time"
	TimeFrom       int64  `json:"time_from"`        // Unix timestamp
	TimeTo         int64  `json:"time_to"`          // Unix timestamp
	PageSize       int    `json:"page_size"`        // 1–100
	Cursor         string `json:"cursor,omitempty"`
	OrderStatus    string `json:"order_status,omitempty"`
}

type GetOrderListResponse struct {
	BaseResponse
	Response *OrderListData `json:"response"`
}

type OrderListData struct {
	More       bool         `json:"more"`
	NextCursor string       `json:"next_cursor"`
	OrderList  []OrderBrief `json:"order_list"`
}

type OrderBrief struct {
	OrderSN string `json:"order_sn"`
}

// ---------------------------------------------------------------------------
// Order — get_order_detail
// ---------------------------------------------------------------------------

type GetOrderDetailRequest struct {
	OrderSNList            []string `json:"order_sn_list"`
	OptionalFields         string   `json:"optional_fields,omitempty"`
	ResponseOptionalFields string   `json:"response_optional_fields,omitempty"`
}

type GetOrderDetailResponse struct {
	BaseResponse
	Response *OrderDetailData `json:"response"`
}

type OrderDetailData struct {
	OrderList []Order `json:"order_list"`
}

type Order struct {
	OrderSN          string      `json:"order_sn"`
	Region           string      `json:"region"`
	Currency         string      `json:"currency"`
	COD              bool        `json:"cod"`
	TotalAmount      float64     `json:"total_amount"`
	OrderStatus      string      `json:"order_status"`
	ShippingCarrier  string      `json:"shipping_carrier"`
	PaymentMethod    string      `json:"payment_method"`
	EstShippingFee   float64     `json:"estimated_shipping_fee"`
	MessageToSeller  string      `json:"message_to_seller"`
	CreateTime       int64       `json:"create_time"`
	UpdateTime       int64       `json:"update_time"`
	DaysToShip       int         `json:"days_to_ship"`
	ShipByDate       int64       `json:"ship_by_date"`
	BuyerUserID      int64       `json:"buyer_user_id"`
	BuyerUsername    string      `json:"buyer_username"`
	RecipientAddress Address     `json:"recipient_address"`
	ItemList         []OrderItem `json:"item_list"`
	InvoiceData      *Invoice    `json:"invoice_data"`
}

type Address struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Town        string `json:"town"`
	District    string `json:"district"`
	City        string `json:"city"`
	State       string `json:"state"`
	Region      string `json:"region"`
	Zipcode     string `json:"zipcode"`
	FullAddress string `json:"full_address"`
}

type OrderItem struct {
	ItemID           int64  `json:"item_id"`
	ItemName         string `json:"item_name"`
	ItemSKU          string `json:"item_sku"`
	ModelID          int64  `json:"model_id"`
	ModelName        string `json:"model_name"`
	ModelSKU         string `json:"model_sku"`
	ModelQuantity    int    `json:"model_quantity_purchased"`
	ModelOrigPrice   float64 `json:"model_original_price"`
	ModelDiscPrice   float64 `json:"model_discounted_price"`
	IsBundle         bool   `json:"is_bundle_deal_item"`
	ImageInfo        *Image `json:"image_info"`
}

type Image struct {
	ImageURL string `json:"image_url"`
}

type Invoice struct {
	Number string `json:"number"`
}
