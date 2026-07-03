// Package fulfillmentsync drives order fulfillment on Lazada. It handles both
// models: dropship (Pack → ReadyToShip via Lazada logistics) and own-fleet
// (push shipped status for seller-shipped packages).
//
// Dropship is a resumable two-step state machine: after Pack succeeds the row is
// persisted as "packed" before ReadyToShip is attempted, so a later failure (or
// a token refresh mid-way) resumes at RTS and never packs twice.
package fulfillmentsync

import (
	"context"
	"strings"

	"github.com/go-resty/resty/v2"

	"github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/client"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

type Syncer struct {
	cfg    client.Config
	tokens *tokenstore.Store
	repo   *postgres.LazadaFulfillmentRepository
	http   *resty.Client // shared across per-call clients to reuse connections
}

func New(cfg client.Config, tokens *tokenstore.Store, repo *postgres.LazadaFulfillmentRepository) *Syncer {
	return &Syncer{cfg: cfg, tokens: tokens, repo: repo, http: client.NewHTTPClient()}
}

// SyncPending processes up to limit in-progress fulfillments, each per its
// delivery model. A failure records the error but keeps the progress status, so
// the row is retried (resumed) next run instead of stranded.
func (s *Syncer) SyncPending(ctx context.Context, limit int) error {
	rows, err := s.repo.ListPending(ctx, limit)
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		logger.Info("no pending fulfillments")
		return nil
	}

	done := 0
	for _, f := range rows {
		if err := s.fulfill(ctx, f); err != nil {
			logger.Error("fulfill", "id", f.ID, "order", f.OrderID, "error", err)
			if merr := s.repo.MarkError(ctx, f.ID, err.Error()); merr != nil {
				logger.Error("mark fulfillment error", "id", f.ID, "error", merr)
			}
			continue
		}
		done++
	}
	logger.Info("fulfillment sync done", "processed", len(rows), "done", done)
	return nil
}

// fulfill runs the right steps for the row's delivery model. Each step persists
// its progress, so a returned error leaves the row resumable at the failed step.
func (s *Syncer) fulfill(ctx context.Context, f postgres.LazadaFulfillment) error {
	if f.DeliveryType == postgres.DeliveryOwnFleet {
		return s.fulfillOwnFleet(ctx, f)
	}
	return s.fulfillDropship(ctx, f)
}

// fulfillOwnFleet pushes the shipped status for a seller-shipped package. The row
// supplies package id + tracking (the seller's own shipment).
func (s *Syncer) fulfillOwnFleet(ctx context.Context, f postgres.LazadaFulfillment) error {
	if err := s.call(ctx, func(c *client.Client) error {
		return c.MarkOwnFleetShipped(ctx, f.PackageID, f.TrackingNumber, f.CarrierCode)
	}); err != nil {
		return err
	}
	if err := s.repo.MarkStatus(ctx, f.ID, postgres.FulfillmentShipped, f.PackageID, f.TrackingNumber); err != nil {
		return err
	}
	logger.Info("fulfillment shipped (own-fleet)", "id", f.ID, "order", f.OrderID, "package", f.PackageID)
	return nil
}

// fulfillDropship packs the items (unless already packed) then marks the package
// ready to ship. Pack's result is persisted before RTS to make the flow resumable.
func (s *Syncer) fulfillDropship(ctx context.Context, f postgres.LazadaFulfillment) error {
	packageID, tracking := f.PackageID, f.TrackingNumber

	if f.SyncStatus != postgres.FulfillmentPacked {
		var res *client.PackResult
		if err := s.call(ctx, func(c *client.Client) error {
			r, e := c.Pack(ctx, splitIDs(f.OrderItemIDs), postgres.DeliveryDropship, f.ShipmentProvider)
			res = r
			return e
		}); err != nil {
			return err
		}
		if len(res.OrderItems) > 0 {
			packageID = res.OrderItems[0].PackageID
			tracking = res.OrderItems[0].TrackingNumber
		}
		// Persist "packed" before RTS so a later failure resumes here, never re-packs.
		if err := s.repo.MarkStatus(ctx, f.ID, postgres.FulfillmentPacked, packageID, tracking); err != nil {
			return err
		}
		logger.Info("fulfillment packed (dropship)", "id", f.ID, "order", f.OrderID, "package", packageID)
	}

	if err := s.call(ctx, func(c *client.Client) error {
		return c.ReadyToShip(ctx, packageID, postgres.DeliveryDropship)
	}); err != nil {
		return err
	}
	if err := s.repo.MarkStatus(ctx, f.ID, postgres.FulfillmentReadyToShip, packageID, tracking); err != nil {
		return err
	}
	logger.Info("fulfillment ready to ship (dropship)", "id", f.ID, "order", f.OrderID, "package", packageID)
	return nil
}

// call runs fn with a token-bearing client, refreshing the token once on an auth
// error and retrying. Because each API step calls this separately (and Pack is
// persisted before RTS), a refresh mid-flow never repeats an already-done step.
func (s *Syncer) call(ctx context.Context, fn func(*client.Client) error) error {
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return err
	}
	err = fn(c)
	if err != nil && client.IsAuthError(err) {
		logger.Warn("fulfill auth error, refreshing token", "error", err)
		if rerr := s.refreshTokens(ctx); rerr != nil {
			return err
		}
		if c, err = s.clientWithStoredToken(ctx); err != nil {
			return err
		}
		err = fn(c)
	}
	return err
}

func splitIDs(csv string) []string {
	parts := strings.Split(csv, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

// clientWithStoredToken / refreshTokens mirror ordersync: tokens live in Redis.
func (s *Syncer) clientWithStoredToken(ctx context.Context) (*client.Client, error) {
	access, refresh, err := s.tokens.Get(ctx)
	if err != nil {
		return nil, err
	}
	cfg := s.cfg
	cfg.AccessToken = access
	cfg.RefreshToken = refresh
	return client.NewWithHTTP(cfg, s.http), nil
}

func (s *Syncer) refreshTokens(ctx context.Context) error {
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return err
	}
	res, err := c.RefreshToken(ctx)
	if err != nil {
		return err
	}
	logger.Info("lazada token refreshed", "expires_in", res.ExpiresIn)
	return s.tokens.Set(ctx, res.AccessToken, res.RefreshToken, 0)
}
