module github.com/hermese-team/phoenix-multichannel-mkp/libraries/go/idempotency

go 1.26

require (
	github.com/redis/go-redis/v9 v9.21.0
)

// idempotency provides a Redis-backed deduplication key store.
// Use to prevent duplicate processing of Kafka events across services.
