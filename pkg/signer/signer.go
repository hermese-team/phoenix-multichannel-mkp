package signer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Sign(secret, message string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

func ShopeeSign(partnerID int64, appSecret, path string, timestamp int64, accessToken string, shopID int64) string {
	base := fmt.Sprintf("%d%s%d%s%d", partnerID, path, timestamp, accessToken, shopID)
	return Sign(appSecret, base)
}

func ShopeeSignPublic(partnerID int64, appSecret, path string, timestamp int64) string {
	base := fmt.Sprintf("%d%s%d", partnerID, path, timestamp)
	return Sign(appSecret, base)
}
