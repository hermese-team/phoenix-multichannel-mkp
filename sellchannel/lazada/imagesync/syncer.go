// Package imagesync uploads product images to Lazada and attaches them to a
// product's sku. It is a resumable two-step flow: after UploadImage succeeds the
// Lazada URL is persisted before SetImages, so a later failure resumes at
// SetImages and never re-uploads.
package imagesync

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/client"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

// maxImageBytes guards the 1MB upload limit before hitting the API.
const maxImageBytes = 1 << 20

// imageDownloader fetches source images from http(s) URLs with a bounded timeout
// so a slow/hanging URL can't block the sync.
var imageDownloader = &http.Client{Timeout: 30 * time.Second}

type Syncer struct {
	cfg    client.Config
	tokens *tokenstore.Store
	repo   *postgres.LazadaProductImageRepository
	http   *resty.Client // shared across per-call clients to reuse connections
}

func New(cfg client.Config, tokens *tokenstore.Store, repo *postgres.LazadaProductImageRepository) *Syncer {
	return &Syncer{cfg: cfg, tokens: tokens, repo: repo, http: client.NewHTTPClient()}
}

// SyncPending processes up to limit in-progress image rows. A failure records the
// error but keeps the progress status, so the row resumes next run.
func (s *Syncer) SyncPending(ctx context.Context, limit int) error {
	rows, err := s.repo.ListPending(ctx, limit)
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		logger.Info("no pending images")
		return nil
	}

	done := 0
	for _, im := range rows {
		if err := s.process(ctx, im); err != nil {
			logger.Error("image", "id", im.ID, "seller_sku", im.SellerSku, "error", err)
			if merr := s.repo.MarkError(ctx, im.ID, err.Error()); merr != nil {
				logger.Error("mark image error", "id", im.ID, "error", merr)
			}
			continue
		}
		done++
	}
	logger.Info("image sync done", "processed", len(rows), "done", done)
	return nil
}

// process uploads the source image (unless already uploaded) then attaches the
// Lazada URL to the sku. Each step persists progress so it is resumable.
func (s *Syncer) process(ctx context.Context, im postgres.LazadaProductImage) error {
	url := im.UploadedURL

	if url == "" {
		data, name, err := readSource(ctx, im.Source)
		if err != nil {
			return err
		}
		if len(data) > maxImageBytes {
			return fmt.Errorf("image %d exceeds 1MB (%d bytes)", im.ID, len(data))
		}
		if err := s.call(ctx, func(c *client.Client) error {
			u, e := c.UploadImage(ctx, name, data)
			url = u
			return e
		}); err != nil {
			return err
		}
		if url == "" {
			return fmt.Errorf("upload returned empty url")
		}
		if err := s.repo.MarkUploaded(ctx, im.ID, url); err != nil {
			return err
		}
		logger.Info("image uploaded", "id", im.ID, "url", url)
	}

	if err := s.call(ctx, func(c *client.Client) error {
		return c.SetImages(ctx, im.SellerSku, []string{url})
	}); err != nil {
		return err
	}
	if err := s.repo.MarkDone(ctx, im.ID); err != nil {
		return err
	}
	logger.Info("image set", "id", im.ID, "seller_sku", im.SellerSku)
	return nil
}

// readSource loads the image bytes from a local path or an http(s) URL, plus a
// file name for the multipart part.
func readSource(ctx context.Context, source string) ([]byte, string, error) {
	name := baseName(source)
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, source, nil)
		if err != nil {
			return nil, "", err
		}
		resp, err := imageDownloader.Do(req)
		if err != nil {
			return nil, "", err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("download image: status %d", resp.StatusCode)
		}
		data, err := io.ReadAll(io.LimitReader(resp.Body, maxImageBytes+1))
		return data, name, err
	}
	data, err := os.ReadFile(source)
	return data, name, err
}

func baseName(source string) string {
	if i := strings.IndexByte(source, '?'); i >= 0 {
		source = source[:i]
	}
	return path.Base(source)
}

// call runs fn with a token-bearing client, refreshing the token once on an auth
// error. Because upload is persisted before set, a refresh never re-uploads.
func (s *Syncer) call(ctx context.Context, fn func(*client.Client) error) error {
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return err
	}
	err = fn(c)
	if err != nil && client.IsAuthError(err) {
		logger.Warn("image auth error, refreshing token", "error", err)
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
