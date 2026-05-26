package store_test

import (
	"context"
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
)

// redisClient creates a test Redis client.
// It skips the test if Redis is unavailable.
func redisClient(t *testing.T) *redis.Client {
	t.Helper()

	addr := os.Getenv("REDIS_URL")
	if addr == "" {
		addr = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Skipf("Redis not available at %s — skipping integration test", addr)
	}

	t.Cleanup(func() {
		client.Close()
	})

	return client
}

func TestRedisConnection_Ping(t *testing.T) {
	client := redisClient(t)

	err := client.Ping(context.Background()).Err()
	if err != nil {
		t.Fatalf("PING failed: %v", err)
	}
}

func TestRedisConnection_SetAndGet(t *testing.T) {
	client := redisClient(t)

	ctx := context.Background()
	key := "test:day4:setget"

	t.Cleanup(func() {
		client.Del(ctx, key)
	})

	if err := client.Set(ctx, key, "hello", 0).Err(); err != nil {
		t.Fatalf("SET failed: %v", err)
	}

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		t.Fatalf("GET failed: %v", err)
	}

	if val != "hello" {
		t.Errorf("got %q, want hello", val)
	}
}

func TestRedisConnection_IncrAndExpire(t *testing.T) {
	client := redisClient(t)

	ctx := context.Background()
	key := "test:day4:incr"

	t.Cleanup(func() {
		client.Del(ctx, key)
	})

	count, err := client.Incr(ctx, key).Result()
	if err != nil {
		t.Fatalf("INCR failed: %v", err)
	}

	if count != 1 {
		t.Errorf("got %d, want 1", count)
	}

	client.Expire(ctx, key, 60000000000)

	count, err = client.Incr(ctx, key).Result()
	if err != nil {
		t.Fatalf("second INCR failed: %v", err)
	}

	if count != 2 {
		t.Errorf("got %d, want 2", count)
	}
}
