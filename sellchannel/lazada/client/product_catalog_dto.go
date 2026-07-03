package client

// Catalog DTOs for the reconcile (pull) flow.
//
// NOTE(mock): GetProducts' sku fields use PascalCase keys (SkuId, SellerSku,
// ShopSku, Status) — unlike order/items/get which is snake_case. The exact set
// of fields is reconstructed; verify against a live GetProducts response.

// CatalogSku is one sku of a product returned by GetProducts / GetProductItem.
type CatalogSku struct {
	SkuID     int64   `json:"SkuId"`
	SellerSku string  `json:"SellerSku"`
	ShopSku   string  `json:"ShopSku"`
	Status    string  `json:"Status"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

// CatalogProduct is one product (its attributes + skus).
type CatalogProduct struct {
	ItemID     int64 `json:"item_id"`
	Attributes struct {
		Name string `json:"name"`
	} `json:"attributes"`
	Skus []CatalogSku `json:"skus"`
}

// productsDTO is the `data` block of /products/get.
type productsDTO struct {
	TotalProducts int              `json:"total_products"`
	Products      []CatalogProduct `json:"products"`
}
