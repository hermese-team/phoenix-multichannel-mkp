package signer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func Sign(secret, message string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

// LazadaSign reproduces the Lazada Open Platform (TOP/IOP) request signature:
//
//	base = apiPath + Σ(key + value) over all params sorted by key asc
//	sign = UPPER(hex(HMAC_SHA256(base, appSecret)))
//
// params must NOT yet contain the "sign" key.
func LazadaSign(apiPath string, params map[string]string, appSecret string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	sb.WriteString(apiPath)
	for _, k := range keys {
		sb.WriteString(k)
		sb.WriteString(params[k])
	}
	return strings.ToUpper(Sign(appSecret, sb.String()))
}

func ShopeeSign(partnerID int64, appSecret, path string, timestamp int64, accessToken string, shopID int64) string {
	base := fmt.Sprintf("%d%s%d%s%d", partnerID, path, timestamp, accessToken, shopID)
	return Sign(appSecret, base)
}

func ShopeeSignPublic(partnerID int64, appSecret, path string, timestamp int64) string {
	base := fmt.Sprintf("%d%s%d", partnerID, path, timestamp)
	return Sign(appSecret, base)
}
