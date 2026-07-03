// Package vouchersync creates seller vouchers on Lazada from pending rows. It is
// an outbound loop like productcreatesync, calling /promotion/voucher/create (one
// voucher per call) and recording the returned promotion id.
package vouchersync

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/client"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

type Syncer struct {
	cfg    client.Config
	tokens *tokenstore.Store
	repo   *postgres.LazadaSellerVoucherRepository
	http   *resty.Client // shared across per-call clients to reuse connections
}

func New(cfg client.Config, tokens *tokenstore.Store, repo *postgres.LazadaSellerVoucherRepository) *Syncer {
	return &Syncer{cfg: cfg, tokens: tokens, repo: repo, http: client.NewHTTPClient()}
}

// CreatePending creates up to limit pending vouchers on Lazada, one call each,
// recording each voucher's outcome (promotion id or error).
func (s *Syncer) CreatePending(ctx context.Context, limit int) error {
	vouchers, err := s.repo.ListPending(ctx, limit)
	if err != nil {
		return err
	}
	if len(vouchers) == 0 {
		logger.Info("no pending vouchers to create")
		return nil
	}

	created := 0
	for _, v := range vouchers {
		promoID, err := s.createOne(ctx, v)
		if err != nil {
			logger.Error("create voucher", "voucher_name", v.VoucherName, "error", err)
			if merr := s.repo.MarkFailed(ctx, v.VoucherName, err.Error()); merr != nil {
				logger.Error("mark voucher failed", "voucher_name", v.VoucherName, "error", merr)
			}
			continue
		}
		if merr := s.repo.MarkCreated(ctx, v.VoucherName, promoID); merr != nil {
			logger.Error("mark voucher created", "voucher_name", v.VoucherName, "error", merr)
			continue
		}
		created++
		logger.Info("voucher created", "voucher_name", v.VoucherName, "promotion_id", promoID)
	}
	logger.Info("voucher create done", "processed", len(vouchers), "created", created)
	return nil
}

// createOne validates then creates one voucher, refreshing the token once on an
// auth error and retrying.
func (s *Syncer) createOne(ctx context.Context, v postgres.LazadaSellerVoucher) (int64, error) {
	req, err := buildCreateRequest(v)
	if err != nil {
		return 0, err
	}
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return 0, err
	}
	promoID, err := c.CreateSellerVoucher(ctx, req)
	if err != nil {
		if !client.IsAuthError(err) {
			return 0, err
		}
		logger.Warn("create voucher auth error, refreshing token", "voucher_name", v.VoucherName, "error", err)
		if rerr := s.refreshTokens(ctx); rerr != nil {
			return 0, err
		}
		if c, err = s.clientWithStoredToken(ctx); err != nil {
			return 0, err
		}
		if promoID, err = c.CreateSellerVoucher(ctx, req); err != nil {
			return 0, err
		}
	}
	return promoID, nil
}

// buildCreateRequest maps a row to the client request and validates the fields
// the API requires (surfacing a clear error instead of a remote rejection).
func buildCreateRequest(v postgres.LazadaSellerVoucher) (client.VoucherCreateRequest, error) {
	voucherType := v.VoucherType
	if voucherType == "" {
		voucherType = "COLLECTIBLE_VOUCHER" // create only supports this type
	}

	switch {
	case v.VoucherName == "":
		return client.VoucherCreateRequest{}, fmt.Errorf("voucher_name is required")
	case v.Apply == "":
		return client.VoucherCreateRequest{}, fmt.Errorf("apply is required (ENTIRE_SHOP / SPECIFIC_PRODUCTS)")
	case v.DisplayArea == "":
		return client.VoucherCreateRequest{}, fmt.Errorf("display_area is required")
	case v.CriteriaOverMoney == "":
		return client.VoucherCreateRequest{}, fmt.Errorf("criteria_over_money is required")
	case v.PeriodStartTime == 0 || v.PeriodEndTime == 0:
		return client.VoucherCreateRequest{}, fmt.Errorf("period_start_time and period_end_time are required")
	case v.PeriodEndTime <= v.PeriodStartTime:
		return client.VoucherCreateRequest{}, fmt.Errorf("period_end_time must be after period_start_time")
	}

	// Discount-type-specific requirement.
	switch v.DiscountType {
	case client.VoucherDiscountMoney:
		if v.OfferingMoneyValueOff == "" {
			return client.VoucherCreateRequest{}, fmt.Errorf("offering_money_value_off is required for %s", client.VoucherDiscountMoney)
		}
	case client.VoucherDiscountPercentage:
		if v.OfferingPercentageDiscountOff <= 0 {
			return client.VoucherCreateRequest{}, fmt.Errorf("offering_percentage_discount_off is required for %s", client.VoucherDiscountPercentage)
		}
	default:
		return client.VoucherCreateRequest{}, fmt.Errorf("unknown discount_type %q (want %s or %s)",
			v.DiscountType, client.VoucherDiscountMoney, client.VoucherDiscountPercentage)
	}

	return client.VoucherCreateRequest{
		VoucherName:                   v.VoucherName,
		VoucherType:                   voucherType,
		Apply:                         v.Apply,
		DisplayArea:                   v.DisplayArea,
		DiscountType:                  v.DiscountType,
		CriteriaOverMoney:             v.CriteriaOverMoney,
		PeriodStartTime:               v.PeriodStartTime,
		PeriodEndTime:                 v.PeriodEndTime,
		CollectStartTime:              v.CollectStartTime,
		PerCustomerLimit:              v.PerCustomerLimit,
		Issued:                        v.Issued,
		OfferingMoneyValueOff:         v.OfferingMoneyValueOff,
		OfferingPercentageDiscountOff: v.OfferingPercentageDiscountOff,
		MaxDiscountOfferingMoneyValue: v.MaxDiscountOfferingMoneyValue,
	}, nil
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
