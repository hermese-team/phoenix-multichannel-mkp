// Package fulfillmentsync drives order fulfillment on Shopee.
// It handles two delivery models:
//   - shopee_logistics: call ship_order with pickup parameters (Shopee arranges logistics)
//   - own_fleet: call ship_order with seller-provided tracking number
//
// Each model is a single API call (simpler than Lazada's Pack→RTS two-step).
// On a token warning or API error the row is left as "pending" so the scheduler retries.
package fulfillmentsync

import (
	"context"
	"fmt"
	"strings"

	pgAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/tokenstore"
)

// permanentFulfillmentErrors are Shopee error codes that indicate a configuration
// or business problem that will not resolve itself — retrying wastes quota and floods logs.
// Orders with these errors are marked 'failed' and require ops intervention.
var permanentFulfillmentErrors = []string{
	"logistics.no_supported_pickup_address", // seller has no pickup address configured
	"logistics.pickup_address_not_found",    // address was deleted
	"error_param",                           // bad request — code bug, not transient
	"logistics.order_cannot_ship",           // order not in shippable state
}

// Syncer processes pending Shopee fulfillments from the DB.
type Syncer struct {
	shopeeClient *client.Client
	tokens       *tokenstore.Store
	repo         *pgAdapter.ShopeeFulfillmentRepository
}

func New(shopeeClient *client.Client, tokens *tokenstore.Store, repo *pgAdapter.ShopeeFulfillmentRepository) *Syncer {
	return &Syncer{shopeeClient: shopeeClient, tokens: tokens, repo: repo}
}

// SyncPending processes up to limit pending fulfillments.
// Failures record the error and leave the row in "pending" for the next run.
func (s *Syncer) SyncPending(ctx context.Context, limit int) error {
	rows, err := s.repo.ListPending(ctx, limit)
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		return nil
	}

	done := 0
	for _, f := range rows {
		if err := s.fulfill(ctx, f); err != nil {
			permanent := isPermanentError(err)
			logger.ErrorContext(ctx, "fulfill order failed",
				"event", "fulfillment.ship.error",
				"order_sn", f.OrderSN,
				"shop_id", f.ShopID,
				"permanent", permanent,
				"error", err,
			)
			if permanent {
				if merr := s.repo.MarkFailed(ctx, f.ID, err.Error()); merr != nil {
					logger.ErrorContext(ctx, "mark failed error",
						"event", "fulfillment.mark_failed.error",
						"id", f.ID,
						"error", merr,
					)
				}
			} else {
				if merr := s.repo.MarkError(ctx, f.ID, err.Error()); merr != nil {
					logger.ErrorContext(ctx, "mark error failed",
						"event", "fulfillment.mark_error.error",
						"id", f.ID,
						"error", merr,
					)
				}
			}
			continue
		}
		done++
	}
	logger.InfoContext(ctx, "fulfillment batch done",
		"event", "fulfillment.batch.done",
		"processed", done,
		"total", len(rows),
	)
	return nil
}

func (s *Syncer) fulfill(ctx context.Context, f pgAdapter.ShopeeFulfillment) error {
	accessToken, _, err := s.tokens.Get(ctx, f.ShopID)
	if err != nil || accessToken == "" {
		return fmt.Errorf("no token for shop %d: %v", f.ShopID, err)
	}

	switch f.DeliveryType {
	case pgAdapter.ShopeeDeliveryOwnFleet:
		return s.fulfillOwnFleet(ctx, f, accessToken)
	default:
		return s.fulfillShopeeLogistics(ctx, f, accessToken)
	}
}

// fulfillOwnFleet ships with seller-provided tracking number.
func (s *Syncer) fulfillOwnFleet(ctx context.Context, f pgAdapter.ShopeeFulfillment, accessToken string) error {
	if f.TrackingNumber == "" {
		return fmt.Errorf("own_fleet requires tracking_number for order %s", f.OrderSN)
	}
	if err := s.shopeeClient.ShipOrderNonIntegrated(ctx, f.ShopID, accessToken, f.OrderSN, f.TrackingNumber); err != nil {
		return err
	}
	if err := s.repo.MarkShipped(ctx, f.ID, f.TrackingNumber); err != nil {
		return err
	}
	logger.InfoContext(ctx, "order shipped",
		"event", "fulfillment.shipped",
		"delivery_type", "own_fleet",
		"order_sn", f.OrderSN,
		"shop_id", f.ShopID,
		"tracking_number", f.TrackingNumber,
	)
	return nil
}

// isPermanentError returns true when the error is a known Shopee business/config error
// that will not resolve itself on retry. These orders are marked 'failed' immediately.
func isPermanentError(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	for _, code := range permanentFulfillmentErrors {
		if strings.Contains(msg, code) {
			return true
		}
	}
	return false
}

// fulfillShopeeLogistics ships via Shopee logistics using the first available pickup slot.
func (s *Syncer) fulfillShopeeLogistics(ctx context.Context, f pgAdapter.ShopeeFulfillment, accessToken string) error {
	params, err := s.shopeeClient.GetShippingParameter(ctx, f.ShopID, accessToken, f.OrderSN)
	if err != nil {
		return fmt.Errorf("get shipping parameter: %w", err)
	}

	// Pick the first available pickup address + timeslot.
	if params.Response.Pickup == nil || len(params.Response.Pickup.AddressList) == 0 {
		return fmt.Errorf("no pickup address for order %s (try own_fleet instead)", f.OrderSN)
	}
	addr := params.Response.Pickup.AddressList[0]
	if len(addr.TimeSlotList) == 0 {
		return fmt.Errorf("no pickup timeslot for order %s", f.OrderSN)
	}
	pickupTimeID := addr.TimeSlotList[0].PickupTimeID

	if err := s.shopeeClient.ShipOrderPickup(ctx, f.ShopID, accessToken, f.OrderSN, addr.AddressID, pickupTimeID); err != nil {
		return err
	}

	// Fetch the tracking number Shopee assigned.
	tracking, err := s.shopeeClient.GetTrackingNumber(ctx, f.ShopID, accessToken, f.OrderSN)
	if err != nil {
		// Non-fatal: order is shipped, tracking fetch failure is logged only.
		logger.WarnContext(ctx, "get tracking number failed (order already shipped)",
			"event", "fulfillment.tracking.error",
			"order_sn", f.OrderSN,
			"error", err,
		)
	}

	if err := s.repo.MarkShipped(ctx, f.ID, tracking); err != nil {
		return err
	}
	logger.InfoContext(ctx, "order shipped",
		"event", "fulfillment.shipped",
		"delivery_type", "shopee_logistics",
		"order_sn", f.OrderSN,
		"shop_id", f.ShopID,
		"tracking_number", tracking,
	)
	return nil
}
