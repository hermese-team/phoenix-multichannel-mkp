package client

// Fulfillment request/response shapes.
//
// NOTE(mock): the STANDARD (dropship) bodies — packReq, readyToShipReq,
// getShipmentProvidersReq and their responses — are NOT in docs/lazada-api-docs.md
// (which lists only error codes) nor in the reference project. They are
// reconstructed from Lazada's public spec and MUST be verified against the live
// API before production use. The SOF (own-fleet) shapes mirror the reference
// project's ownfleet.go, which is known to work.

// ── Standard / dropship ─────────────────────────────────────────────────────

type packOrderItem struct {
	OrderItemID string `json:"order_item_id"`
}

type packOrderList struct {
	OrderItemList    []packOrderItem `json:"order_item_list"`
	DeliveryType     string          `json:"delivery_type"`
	ShipmentProvider string          `json:"shipment_provider,omitempty"`
}

// packRequest is the packReq body.
type packRequest struct {
	PackOrderList packOrderList `json:"pack_order_list"`
}

// PackResult is Pack's per-item outcome (reconstructed). The first item's
// package id + tracking number drive the next step (ReadyToShip).
type PackResult struct {
	OrderItems []struct {
		OrderItemID    int64  `json:"order_item_id"`
		PackageID      string `json:"package_id"`
		TrackingNumber string `json:"tracking_number"`
	} `json:"order_items"`
}

type rtsPackage struct {
	PackageID string `json:"package_id"`
}

// readyToShipRequest is the readyToShipReq body.
type readyToShipRequest struct {
	DeliveryType string       `json:"delivery_type"`
	Packages     []rtsPackage `json:"packages"`
}

// ShipmentProvider is one active provider from GetShipmentProviders.
type ShipmentProvider struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	ProviderType string `json:"provider_type"`
}

// DocumentResult is the shipping-label document from GetShippingDocument.
type DocumentResult struct {
	File     string `json:"file"` // base64-encoded PDF/HTML
	MimeType string `json:"mime_type"`
}

// ── SOF / own-fleet (mirrors reference ownfleet.go) ─────────────────────────

type sofPackage struct {
	PackageID string `json:"package_id"`
}

type sofPackages struct {
	Packages []sofPackage `json:"packages"`
}

type sofTrackInfo struct {
	TrackInfo struct {
		LatestStatus struct {
			Status    string `json:"status"`
			SubStatus string `json:"subStatus"`
		} `json:"latestStatus"`
		LatestEvent struct {
			EventTime int64 `json:"eventTime"` // .NET ticks (see fulfillment_sof.go)
		} `json:"latestEvent"`
	} `json:"trackInfo"`
}
