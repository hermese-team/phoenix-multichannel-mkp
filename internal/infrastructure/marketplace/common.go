// Package marketplace contains outbound HTTP gateways to external sales
// channels (TikTok, Lazada, Shopee), built on a shared resty client.
package marketplace

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"

	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
	domainerrors "github.com/ascend/phoenix-multichannel-mkp/pkg/errors"
)

// baseGateway holds the wiring common to every marketplace gateway. Channel
// specifics (payload shape, paths) live in each concrete gateway.
type baseGateway struct {
	channel entity.Channel
	rc      *resty.Client
	baseURL string
	apiKey  string
}

func (g baseGateway) Channel() entity.Channel { return g.channel }

// pushResponse is the minimal shape we read an external id from.
type pushResponse struct {
	ID string `json:"id"`
}

// post sends body to baseURL+path with bearer auth and returns the external id.
func (g baseGateway) post(ctx context.Context, path string, body any) (string, error) {
	var out pushResponse
	resp, err := g.rc.R().
		SetContext(ctx).
		SetAuthToken(g.apiKey).
		SetBody(body).
		SetResult(&out).
		Post(g.baseURL + path)
	if err != nil {
		return "", err
	}
	if resp.IsError() {
		return "", domainerrors.New(
			"MARKETPLACE_ERROR",
			fmt.Sprintf("%s returned status %d", g.channel, resp.StatusCode()),
			domainerrors.ErrInternalServer,
		)
	}
	return out.ID, nil
}
