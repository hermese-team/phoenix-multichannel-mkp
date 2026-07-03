package client

// CreateProductResponse is the `data` block of a successful /product/create.
type CreateProductResponse struct {
	ItemID  int64 `json:"item_id"`
	SkuList []struct {
		SellerSku string `json:"seller_sku"`
		SkuID     int64  `json:"sku_id"`
		ShopSku   string `json:"shop_sku"`
	} `json:"sku_list"`
}
