package client

import "strconv"

// UpdateStockResponse is the /product/price_quantity/update response.
// Code == "0" means success.
type UpdateStockResponse struct {
	Code      string         `json:"code"`
	Detail    []UpdateDetail `json:"detail"`
	RequestID string         `json:"request_id"`
}

// UpdateDetail is a per-item result entry.
type UpdateDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
	ItemID  int64  `json:"item_id"`
}

// MessageForItem returns the error message for the given Lazada item id, if any.
func (r *UpdateStockResponse) MessageForItem(itemID string) (string, bool) {
	if r == nil {
		return "", false
	}
	for _, d := range r.Detail {
		if strconv.FormatInt(d.ItemID, 10) == itemID {
			return d.Message, true
		}
	}
	return "", false
}
