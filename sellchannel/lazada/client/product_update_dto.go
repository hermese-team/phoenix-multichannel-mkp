package client

// UpdateProductResponse is the `data` block of a successful /product/update.
type UpdateProductResponse struct {
	ItemID int64 `json:"item_id"`
}
