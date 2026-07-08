package redis

import (
	"context"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type Client struct {
	rdb *goredis.Client
}

func New(cfg Config) (*Client, error) {
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		_ = rdb.Close()
		return nil, fmt.Errorf("ping redis: %w", err)
	}
	return &Client{rdb: rdb}, nil
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.rdb.Get(ctx, key).Result()
}

// GetString returns the value and whether the key exists. A missing key is not
// an error (found == false).
func (c *Client) GetString(ctx context.Context, key string) (string, bool, error) {
	v, err := c.rdb.Get(ctx, key).Result()
	if err == goredis.Nil {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return v, true, nil
}

func (c *Client) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return c.rdb.Set(ctx, key, value, ttl).Err()
}

// SetMany atomically sets several key/value pairs with the same ttl (0 = no
// expiry) in a single MULTI/EXEC transaction, so callers never observe a
// partially-updated set.
func (c *Client) SetMany(ctx context.Context, ttl time.Duration, kv map[string]string) error {
	pipe := c.rdb.TxPipeline()
	for k, v := range kv {
		pipe.Set(ctx, k, v, ttl)
	}
	_, err := pipe.Exec(ctx)
	return err
}

// SetNX sets the key only if it does not exist. It returns true when the key
// was newly set (first time seen), false when it already existed.
func (c *Client) SetNX(ctx context.Context, key, value string, ttl time.Duration) (bool, error) {
	return c.rdb.SetNX(ctx, key, value, ttl).Result()
}

func (c *Client) Delete(ctx context.Context, key string) error {
	return c.rdb.Del(ctx, key).Err()
}

// SAddWithTTL adds members to a Redis Set and (re-)sets the TTL on the key.
// Used for safety-net sets where membership and window expiry are coupled:
//
//	SADD  safety-net:{channel}:processed {member}
//	EXPIRE safety-net:{channel}:processed {ttl}
//
// The EXPIRE is refreshed on every write so the window slides with traffic.
func (c *Client) SAddWithTTL(ctx context.Context, key string, ttl time.Duration, members ...string) error {
	args := make([]interface{}, len(members))
	for i, m := range members {
		args[i] = m
	}
	pipe := c.rdb.TxPipeline()
	pipe.SAdd(ctx, key, args...)
	pipe.Expire(ctx, key, ttl)
	_, err := pipe.Exec(ctx)
	return err
}

// SIsMember returns true if member belongs to the set at key.
func (c *Client) SIsMember(ctx context.Context, key, member string) (bool, error) {
	return c.rdb.SIsMember(ctx, key, member).Result()
}

// SMIsMember checks whether each member belongs to the set at key and returns
// a bool slice in the same order as members.
// Requires Redis 6.2+ (SMISMEMBER command).
func (c *Client) SMIsMember(ctx context.Context, key string, members ...string) ([]bool, error) {
	args := make([]interface{}, len(members))
	for i, m := range members {
		args[i] = m
	}
	return c.rdb.SMIsMember(ctx, key, args...).Result()
}
