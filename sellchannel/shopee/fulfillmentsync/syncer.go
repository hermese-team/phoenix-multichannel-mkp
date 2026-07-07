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
	"log"

	pgAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/tokenstore"
)

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
			log.Printf("[fulfillment] order %s shop %d: %v", f.OrderSN, f.ShopID, err)
			if merr := s.repo.MarkError(ctx, f.ID, err.Error()); merr != nil {
				log.Printf("[fulfillment] mark error id %d: %v", f.ID, merr)
			}
			continue
		}
		done++
	}
	log.Printf("[fulfillment] processed %d/%d", done, len(rows))
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
	log.Printf("[fulfillment] shipped (own_fleet) order %s shop %d tracking %s", f.OrderSN, f.ShopID, f.TrackingNumber)
	return nil
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
		log.Printf("[fulfillment] get tracking number order %s: %v (order already shipped)", f.OrderSN, err)
	}

	if err := s.repo.MarkShipped(ctx, f.ID, tracking); err != nil {
		return err
	}
	log.Printf("[fulfillment] shipped (shopee_logistics) order %s shop %d tracking %s", f.OrderSN, f.ShopID, tracking)
	return nil
}
