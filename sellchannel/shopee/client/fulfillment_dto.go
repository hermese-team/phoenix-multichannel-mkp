package client

// ShipOrderResponse is the response from /api/v2/logistics/ship_order.
type ShipOrderResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

// ShippingParameterResponse is returned by /api/v2/logistics/get_shipping_parameter.
// It tells us which branch (pickup / dropoff / non_integrated) the order uses.
type ShippingParameterResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Response  struct {
		Pickup         *PickupInfo         `json:"pickup"`
		Dropoff        *DropoffInfo        `json:"dropoff"`
		NonIntegrated  *NonIntegratedInfo  `json:"non_integrated"`
	} `json:"response"`
}

type PickupInfo struct {
	AddressList []struct {
		AddressID     int64  `json:"address_id"`
		AddressType   []string `json:"address_type"`
		TimeSlotList  []struct {
			PickupTimeID string `json:"pickup_time_id"`
		} `json:"time_slot_list"`
	} `json:"address_list"`
}

type DropoffInfo struct {
	BranchList []struct {
		BranchID int64 `json:"branch_id"`
	} `json:"branch_list"`
}

type NonIntegratedInfo struct {
	// empty struct — just the presence of this key means non-integrated
}

// TrackingNumberResponse is returned by /api/v2/logistics/get_tracking_number.
type TrackingNumberResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Response  struct {
		TrackingNumber string `json:"tracking_number"`
	} `json:"response"`
}
