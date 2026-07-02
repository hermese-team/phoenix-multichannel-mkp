package client

type BaseResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type AddItemRequest struct {
	ItemName    string `json:"item_name"`
	Description string `json:"description"`
	SellerSku   string `json:"item_sku"`
}

type AddItemResponse struct {
	BaseResponse
	Response struct {
		ItemID int64 `json:"item_id"`
	} `json:"response"`
}

type UpdateItemRequest struct {
	ItemID      int64  `json:"item_id"`
	ItemName    string `json:"item_name"`
	Description string `json:"description"`
	SellerSku   string `json:"item_sku"`
}
