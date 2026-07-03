package client

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/okdev/marketplace-sync/pkg/signer"
)

// UploadImage uploads a single JPG/PNG image (max 1MB) via POST /image/upload and
// returns the Lazada-hosted URL. It is a multipart upload, so it can't go through
// do(): the image bytes are a file part.
//
// NOTE(mock): the image byte part is assumed excluded from the signature (only
// system params are signed) and the response shape is reconstructed as
// {"image":{"url":...}} — verify both against the live API.
func (c *Client) UploadImage(ctx context.Context, fileName string, data []byte) (string, error) {
	sys := map[string]string{
		"app_key":     c.cfg.AppKey,
		"timestamp":   strconv.FormatInt(time.Now().UnixMilli(), 10),
		"sign_method": "sha256",
	}
	if c.cfg.AccessToken != "" {
		sys["access_token"] = c.cfg.AccessToken
	}
	signature := signer.LazadaSign("/image/upload", sys, c.cfg.AppSecret)

	endpoint := strings.TrimRight(c.cfg.BaseURL, "/") + "/image/upload"
	resp, err := c.http.R().SetContext(ctx).
		SetQueryParams(sys).
		SetQueryParam("sign", signature).
		SetMultipartField("image", fileName, imageContentType(fileName), bytes.NewReader(data)).
		Post(endpoint)
	if err != nil {
		return "", err
	}

	var env envelope
	if err := json.Unmarshal(resp.Body(), &env); err != nil {
		return "", fmt.Errorf("decode upload envelope: %w", err)
	}
	if env.Code != "0" {
		return "", &APIError{APIPath: "/image/upload", Code: env.Code, Message: env.Message}
	}
	var dto uploadImageDTO
	if err := json.Unmarshal(env.Data, &dto); err != nil {
		return "", fmt.Errorf("decode upload image: %w", err)
	}
	return dto.Image.URL, nil
}

type uploadImageDTO struct {
	Image struct {
		URL      string `json:"url"`
		HashCode string `json:"hash_code"`
	} `json:"image"`
}

// imageContentType picks the multipart part's MIME type from the file extension
// (Lazada accepts JPG/PNG); defaults to image/jpeg.
func imageContentType(fileName string) string {
	if strings.EqualFold(path.Ext(fileName), ".png") {
		return "image/png"
	}
	return "image/jpeg"
}

// SetImages calls POST /images/set to associate image URLs with a product's sku
// (max 8 images per sku).
func (c *Client) SetImages(ctx context.Context, sellerSku string, imageURLs []string) error {
	payload, err := xml.Marshal(setImagesRequest{Skus: []setImagesSku{{
		SellerSku: sellerSku,
		Images:    imageURLs,
	}}})
	if err != nil {
		return err
	}
	_, _, err = c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/images/set", map[string]string{"payload": string(payload)})
	return err
}

type setImagesSku struct {
	SellerSku string   `xml:"SellerSku"`
	Images    []string `xml:"Images>Image"`
}

type setImagesRequest struct {
	XMLName xml.Name       `xml:"Request"`
	Skus    []setImagesSku `xml:"Product>Skus>Sku"`
}
