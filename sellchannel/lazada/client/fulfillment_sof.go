package client

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// SOF (Seller Own Fleet / DBS) status updates — the seller ships themselves and
// pushes the resulting status. Payloads mirror the reference project's
// ownfleet.go, which is known to work.

// ticksAtUnixEpoch is the number of .NET ticks (100ns intervals since year 1) at
// the Unix epoch. The SOF status/update event time is a .NET tick timestamp
// (the reference ported this from a C# integration).
const ticksAtUnixEpoch = 621355968000000000

func dotNetTicks(t time.Time) int64 { return ticksAtUnixEpoch + t.UnixNano()/100 }

// MarkOwnFleetShipped calls POST /order/package/sof/status/update to mark a
// self-shipped package as shipped, attaching its tracking number.
func (c *Client) MarkOwnFleetShipped(ctx context.Context, packageID, trackingNumber, carrierCode string) error {
	var ti sofTrackInfo
	ti.TrackInfo.LatestStatus.Status = "shipped"
	ti.TrackInfo.LatestStatus.SubStatus = "shipped"
	ti.TrackInfo.LatestEvent.EventTime = dotNetTicks(time.Now())
	body, err := json.Marshal(ti)
	if err != nil {
		return err
	}
	_, _, err = c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		"/order/package/sof/status/update", map[string]string{
			"trackingNumber": trackingNumber,
			"source":         "OPENAPI",
			"carrierCode":    carrierCode,
			"tag":            packageID,
			"trackInfo":      string(body),
		})
	return err
}

// MarkOwnFleetDelivered calls POST /order/package/sof/delivered.
func (c *Client) MarkOwnFleetDelivered(ctx context.Context, packageID string) error {
	return c.sofPackageAction(ctx, "/order/package/sof/delivered", "dbsDeliveryReq", packageID)
}

// MarkOwnFleetFailed calls POST /order/package/sof/failed_delivery.
func (c *Client) MarkOwnFleetFailed(ctx context.Context, packageID string) error {
	return c.sofPackageAction(ctx, "/order/package/sof/failed_delivery", "dbsFailedDeliveryReq", packageID)
}

// sofPackageAction posts {"packages":[{"package_id":…}]} under the given form key.
func (c *Client) sofPackageAction(ctx context.Context, apiPath, formKey, packageID string) error {
	body, err := json.Marshal(sofPackages{Packages: []sofPackage{{PackageID: packageID}}})
	if err != nil {
		return err
	}
	_, _, err = c.do(ctx, http.MethodPost, c.cfg.BaseURL, c.cfg.AccessToken,
		apiPath, map[string]string{formKey: string(body)})
	return err
}
