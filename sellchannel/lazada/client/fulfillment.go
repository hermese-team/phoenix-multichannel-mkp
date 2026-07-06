package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetShipmentProviders calls POST /order/shipment/providers/get and returns the
// active shipping providers (needed as Pack's shipment_provider for dropship).
//
// NOTE(mock): the request body getShipmentProvidersReq is undocumented here; we
// send "{}" (reconstructed). Verify against the live API.
func (c *Client) GetShipmentProviders(ctx context.Context) ([]ShipmentProvider, error) {
	_, env, err := c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/order/shipment/providers/get", map[string]string{"getShipmentProvidersReq": "{}"})
	if err != nil {
		return nil, err
	}
	// Reconstructed: providers under data. Try {"shipment_providers":[...]} then a bare array.
	var wrapped struct {
		ShipmentProviders []ShipmentProvider `json:"shipment_providers"`
	}
	if err := json.Unmarshal(env.Detail, &wrapped); err == nil && len(wrapped.ShipmentProviders) > 0 {
		return wrapped.ShipmentProviders, nil
	}
	var arr []ShipmentProvider
	if err := json.Unmarshal(env.Detail, &arr); err != nil {
		return nil, fmt.Errorf("decode shipment providers: %w", err)
	}
	return arr, nil
}

// Pack calls POST /order/fulfill/pack to pack order items into a package.
// deliveryType is e.g. "dropship"; shipmentProvider comes from GetShipmentProviders.
func (c *Client) Pack(ctx context.Context, orderItemIDs []string, deliveryType, shipmentProvider string) (*PackResult, error) {
	items := make([]packOrderItem, len(orderItemIDs))
	for i, id := range orderItemIDs {
		items[i] = packOrderItem{OrderItemID: id}
	}
	payload, err := json.Marshal(packRequest{PackOrderList: packOrderList{
		OrderItemList:    items,
		DeliveryType:     deliveryType,
		ShipmentProvider: shipmentProvider,
	}})
	if err != nil {
		return nil, fmt.Errorf("marshal packReq: %w", err)
	}
	_, env, err := c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/order/fulfill/pack", map[string]string{"packReq": string(payload)})
	if err != nil {
		return nil, err
	}
	var res PackResult
	if len(env.Detail) > 0 {
		if err := json.Unmarshal(env.Detail, &res); err != nil {
			return nil, fmt.Errorf("decode pack result: %w", err)
		}
	}
	return &res, nil
}

// ReadyToShip calls POST /order/package/rts to mark a packed package ready to ship.
func (c *Client) ReadyToShip(ctx context.Context, packageID, deliveryType string) error {
	payload, err := json.Marshal(readyToShipRequest{
		DeliveryType: deliveryType,
		Packages:     []rtsPackage{{PackageID: packageID}},
	})
	if err != nil {
		return fmt.Errorf("marshal readyToShipReq: %w", err)
	}
	_, _, err = c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/order/package/rts", map[string]string{"readyToShipReq": string(payload)})
	return err
}

// GetShippingDocument calls GET /order/document/get for the shipping label of the
// given order items (they must be at least packed). Returns the base64 document.
func (c *Client) GetShippingDocument(ctx context.Context, orderItemIDs []string) (*DocumentResult, error) {
	ids, err := json.Marshal(orderItemIDs)
	if err != nil {
		return nil, fmt.Errorf("marshal order_item_ids: %w", err)
	}
	raw, err := c.doSigned(ctx, "/order/document/get", map[string]string{
		"doc_type":       "shippingLabel",
		"order_item_ids": string(ids),
	})
	if err != nil {
		return nil, err
	}
	// Reconstructed: {"document":{"file":"...","mime_type":"..."}}
	var wrapped struct {
		Document DocumentResult `json:"document"`
	}
	if err := json.Unmarshal(raw, &wrapped); err != nil {
		return nil, fmt.Errorf("decode document: %w", err)
	}
	return &wrapped.Document, nil
}
