package shopee

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"time"
)

// Signer builds signed query parameters for every Shopee API v2 call.
//
// Shopee V2 signing rules:
//   - Shop-level API:   base = "{partnerID}{path}{timestamp}{accessToken}{shopID}"
//   - Partner-level API: base = "{partnerID}{path}{timestamp}"
//
// sign = HMAC-SHA256(base, partnerKey) → hex lowercase
// Required query params: partner_id, timestamp, sign, [shop_id, access_token]
type Signer struct {
	partnerID  int64
	partnerKey string
}

// NewSigner creates a Signer with the app's partner credentials.
func NewSigner(partnerID int64, partnerKey string) *Signer {
	return &Signer{partnerID: partnerID, partnerKey: partnerKey}
}

// SignShop returns query params for a shop-level API call.
// shopID and accessToken are required; they are part of the base string.
func (s *Signer) SignShop(path string, shopID int64, accessToken string) url.Values {
	ts := time.Now().Unix()
	base := fmt.Sprintf("%d%s%d%s%d", s.partnerID, path, ts, accessToken, shopID)
	sign := s.hmac(base)

	q := url.Values{}
	q.Set("partner_id", fmt.Sprintf("%d", s.partnerID))
	q.Set("timestamp", fmt.Sprintf("%d", ts))
	q.Set("sign", sign)
	q.Set("shop_id", fmt.Sprintf("%d", shopID))
	q.Set("access_token", accessToken)
	return q
}

// SignPartner returns query params for a partner-level API call (no shop context).
func (s *Signer) SignPartner(path string) url.Values {
	ts := time.Now().Unix()
	base := fmt.Sprintf("%d%s%d", s.partnerID, path, ts)
	sign := s.hmac(base)

	q := url.Values{}
	q.Set("partner_id", fmt.Sprintf("%d", s.partnerID))
	q.Set("timestamp", fmt.Sprintf("%d", ts))
	q.Set("sign", sign)
	return q
}

func (s *Signer) hmac(base string) string {
	h := hmac.New(sha256.New, []byte(s.partnerKey))
	h.Write([]byte(base))
	return hex.EncodeToString(h.Sum(nil))
}
