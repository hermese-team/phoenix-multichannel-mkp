// lazada-category browses the Lazada category tree and a category's attributes —
// the discovery step before seeding the product-create table. Category endpoints
// need no access token (app key/secret sign the request).
//
// Usage:
//
//	set -a; source .env; set +a
//	go run ./sellchannel/lazada/cmd/lazada-category            # list top-level categories
//	LAZADA_CATEGORY_ID=<leaf id> go run ./.../lazada-category  # list that category's attributes
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/okdev/marketplace-sync/config"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/client"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config: %v\n", err)
		os.Exit(1)
	}

	c := client.New(cfg.Lazada)
	lang := os.Getenv("LAZADA_LANG") // optional, e.g. th_TH

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// With a category id: list that leaf category's attributes.
	if catID := os.Getenv("LAZADA_CATEGORY_ID"); catID != "" {
		attrs, err := c.GetCategoryAttributes(ctx, catID, lang)
		if err != nil {
			fmt.Fprintf(os.Stderr, "GetCategoryAttributes: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Category %s — %d attribute(s):\n", catID, len(attrs))
		for _, a := range attrs {
			req := ""
			if a.IsMandatory == 1 {
				req = "  [MANDATORY]"
			}
			fmt.Printf("  %-30s %-14s %s%s\n", a.Name, a.InputType, a.Label, req)
		}
		return
	}

	// Otherwise: list top-level categories.
	nodes, err := c.GetCategoryTree(ctx, lang)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetCategoryTree: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Top-level categories (%d) — set LAZADA_CATEGORY_ID=<leaf id> to list its attributes:\n", len(nodes))
	for _, n := range nodes {
		leaf := ""
		if n.Leaf {
			leaf = " (leaf)"
		}
		fmt.Printf("  %-10d %s%s\n", n.CategoryID, n.Name, leaf)
	}
}
