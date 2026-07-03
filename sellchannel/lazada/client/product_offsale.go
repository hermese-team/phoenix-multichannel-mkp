package client

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// DeactivateProduct calls POST /product/deactivate to take a whole product
// offline by its ItemId.
//
// NOTE(mock): apiRequestBody is only documented as "ItemId mandatory, Skus
// optional" — reconstructed as {"itemId":<id>}. Verify against the live API.
func (c *Client) DeactivateProduct(ctx context.Context, itemID int64) error {
	body, err := json.Marshal(map[string]int64{"itemId": itemID})
	if err != nil {
		return err
	}
	_, _, err = c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/product/deactivate", map[string]string{"apiRequestBody": string(body)})
	return err
}

// RemoveProduct calls POST /product/remove to delete the given seller skus
// (max 50 per request).
//
// NOTE: seller_sku_list is deprecated in favour of sku_id_list; we use it here
// because the mock table keys by seller_sku.
func (c *Client) RemoveProduct(ctx context.Context, sellerSkus []string) error {
	list, err := json.Marshal(sellerSkus)
	if err != nil {
		return err
	}
	_, _, err = c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/product/remove", map[string]string{"seller_sku_list": string(list)})
	return err
}

// RemoveSku calls POST /product/sku/remove to delete specific SKUs (and their
// sales attributes) of a product. variationName is the sale-attribute name of
// the removed variation (e.g. "color_family").
func (c *Client) RemoveSku(ctx context.Context, itemID int64, variationName string, sellerSkus []string) error {
	skus := make([]removeSkuSku, len(sellerSkus))
	for i, s := range sellerSkus {
		skus[i] = removeSkuSku{SellerSku: s}
	}
	payload, err := xml.Marshal(removeSkuRequest{ItemID: itemID, Variation: variationName, Skus: skus})
	if err != nil {
		return err
	}
	_, _, err = c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/product/sku/remove", map[string]string{"payload": string(payload)})
	return err
}

type removeSkuSku struct {
	SellerSku string `xml:"SellerSku"`
}

// removeSkuRequest marshals to (per the docs example):
//
//	<Request><Product><ItemId>…</ItemId>
//	  <variation><variation1><name>…</name></variation1></variation>
//	  <Skus><Sku><SellerSku>…</SellerSku></Sku>…</Skus>
//	</Product></Request>
type removeSkuRequest struct {
	XMLName   xml.Name       `xml:"Request"`
	ItemID    int64          `xml:"Product>ItemId"`
	Variation string         `xml:"Product>variation>variation1>name"`
	Skus      []removeSkuSku `xml:"Product>Skus>Sku"`
}
