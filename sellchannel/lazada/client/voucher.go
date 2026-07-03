package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Voucher discount types (voucher_discount_type).
const (
	VoucherDiscountMoney      = "MONEY_VALUE_OFF"
	VoucherDiscountPercentage = "PERCENTAGE_DISCOUNT_OFF"
)

// VoucherCreateRequest is the input to CreateSellerVoucher. Epoch fields are in
// milliseconds (Lazada's convention for the voucher timestamps). Optional
// discount fields depend on DiscountType and are omitted from the form when empty.
type VoucherCreateRequest struct {
	VoucherName       string // voucher_name
	VoucherType       string // voucher_type — create only supports COLLECTIBLE_VOUCHER
	Apply             string // ENTIRE_SHOP / SPECIFIC_PRODUCTS
	DisplayArea       string // REGULAR_CHANNEL / STORE_FOLLOWER / OFFLINE / LIVE_STREAM / CEM_SELLER
	DiscountType      string // MONEY_VALUE_OFF / PERCENTAGE_DISCOUNT_OFF
	CriteriaOverMoney string // min order value that triggers the discount
	PeriodStartTime   int64  // usable-from (epoch ms)
	PeriodEndTime     int64  // usable-until (epoch ms)
	CollectStartTime  int64  // optional — when customers can collect it (epoch ms)
	PerCustomerLimit  int    // limit — vouchers per customer
	Issued            int    // issued — total quantity

	// Discount detail — set the pair matching DiscountType.
	OfferingMoneyValueOff         string // for MONEY_VALUE_OFF
	OfferingPercentageDiscountOff int    // for PERCENTAGE_DISCOUNT_OFF
	MaxDiscountOfferingMoneyValue string // optional cap, PERCENTAGE_DISCOUNT_OFF only
}

// formParams builds the x-www-form-urlencoded body, dropping empty optionals so
// Lazada doesn't reject blank values.
func (r VoucherCreateRequest) formParams() map[string]string {
	p := map[string]string{
		"voucher_name":          r.VoucherName,
		"voucher_type":          r.VoucherType,
		"apply":                 r.Apply,
		"display_area":          r.DisplayArea,
		"voucher_discount_type": r.DiscountType,
		"criteria_over_money":   r.CriteriaOverMoney,
		"period_start_time":     strconv.FormatInt(r.PeriodStartTime, 10),
		"period_end_time":       strconv.FormatInt(r.PeriodEndTime, 10),
		"limit":                 strconv.Itoa(r.PerCustomerLimit),
		"issued":                strconv.Itoa(r.Issued),
	}
	if r.CollectStartTime > 0 {
		p["collect_start"] = strconv.FormatInt(r.CollectStartTime, 10)
	}
	if r.OfferingMoneyValueOff != "" {
		p["offering_money_value_off"] = r.OfferingMoneyValueOff
	}
	if r.OfferingPercentageDiscountOff > 0 {
		p["offering_percentage_discount_off"] = strconv.Itoa(r.OfferingPercentageDiscountOff)
	}
	if r.MaxDiscountOfferingMoneyValue != "" {
		p["max_discount_offering_money_value"] = r.MaxDiscountOfferingMoneyValue
	}
	return p
}

// CreateSellerVoucher calls POST /promotion/voucher/create and returns the new
// promotion ID.
//
// NOTE(mock): the promotion endpoints report business results with top-level
// success/error_code/error_msg fields (siblings of the standard gateway
// code/message), and `data` carries the promotion ID on success. do() only
// validates the gateway code, so success is checked here. Verify this envelope
// shape against the live API.
func (c *Client) CreateSellerVoucher(ctx context.Context, req VoucherCreateRequest) (int64, error) {
	rawBody, _, err := c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/promotion/voucher/create", req.formParams())
	if err != nil {
		return 0, err
	}
	var vr voucherCreateEnvelope
	if err := json.Unmarshal(rawBody, &vr); err != nil {
		return 0, fmt.Errorf("decode voucher create: %w", err)
	}
	// Check success before decoding data: on failure `data` may be null or an
	// object carrying error detail, so surface error_msg rather than a decode error.
	if !vr.Success {
		return 0, &APIError{
			APIPath: "/promotion/voucher/create",
			Code:    strconv.Itoa(vr.ErrorCode),
			Message: vr.ErrorMsg,
		}
	}
	promotionID, err := strconv.ParseInt(strings.Trim(string(vr.Data), `"`), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("decode voucher promotion id from %q: %w", string(vr.Data), err)
	}
	return promotionID, nil
}

type voucherCreateEnvelope struct {
	Success   bool            `json:"success"`
	ErrorCode int             `json:"error_code"`
	ErrorMsg  string          `json:"error_msg"`
	Data      json.RawMessage `json:"data"` // promotion ID (Number, sometimes quoted)
}
