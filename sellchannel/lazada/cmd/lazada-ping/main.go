// lazada-ping is a minimal connectivity test for the Lazada Open Platform.
//
// App key/secret + base URL come from config (env: LAZADA_APP_KEY, ...).
// The access token is read from Redis (managed by the scheduler); set
// LAZADA_ACCESS_TOKEN to override. Without LAZADA_ORDER_ID it calls /seller/get
// only; with it, also fetches that order's detail + items.
//
// Usage:
//
//	set -a; source .env; set +a
//	[LAZADA_ORDER_ID=xxx] go run ./sellchannel/lazada/cmd/lazada-ping
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/okdev/marketplace-sync/config"
	redisAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/client"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config: %v\n", err)
		os.Exit(1)
	}
	lz := cfg.Lazada

	// Access token: prefer env override, else read from Redis.
	lz.AccessToken = os.Getenv("LAZADA_ACCESS_TOKEN")
	if lz.AccessToken == "" {
		rdb, err := redisAdapter.New(cfg.Redis)
		if err != nil {
			fmt.Fprintf(os.Stderr, "connect redis: %v\n", err)
			os.Exit(1)
		}
		access, _, err := tokenstore.New(rdb).Get(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "read token from redis: %v\n", err)
			os.Exit(1)
		}
		lz.AccessToken = access
	}
	if lz.AccessToken == "" {
		fmt.Fprintln(os.Stderr, "no access token — set LAZADA_ACCESS_TOKEN or seed redis (lazada-token)")
		os.Exit(1)
	}

	c := client.New(lz)
	orderID := os.Getenv("LAZADA_ORDER_ID") // optional

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 1) Connectivity check — /seller/get needs only the access token.
	seller, err := c.GetSeller(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR /seller/get: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("OK /seller/get — connection + signing + token valid.\n")
	fmt.Printf("  data: %s\n\n", string(seller))

	if orderID == "" {
		fmt.Println("No LAZADA_ORDER_ID set — skipping order fetch. Done.")
		return
	}

	// 2) Order fetch (only when an order id is provided).
	fmt.Printf("Order ID: %s\n", orderID)
	detail, err := c.GetOrderDetail(ctx, orderID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR GetOrderDetail: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("OK — OrderID: %d  Statuses: %v  PaymentMethod: %s\n", detail.OrderID, detail.Statuses, detail.PaymentMethod)

	items, err := c.GetOrderItems(ctx, orderID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR GetOrderItems: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("OK — %d item(s):\n", len(items))
	for _, it := range items {
		fmt.Printf("  SKU=%-20s  Status=%-20s  Name=%s\n", it.Sku, it.Status, it.Name)
	}
}
