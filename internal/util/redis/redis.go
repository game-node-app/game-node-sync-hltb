package redis

import (
	"context"
	"game-node-sync-hltb/internal/util"
	"github.com/redis/go-redis/v9"
	"time"
)

func Url() string {
	return util.GetEnv("REDIS_ADDR", "localhost:9112")
}

func CreateClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: Url(),
	})
}

func Get(key string) (string, error) {
	client := CreateClient()
	defer client.Close()

	ctx := context.Background()

	return client.Get(ctx, key).Result()
}

func Set(key string, value string, exp *time.Duration) error {
	client := CreateClient()
	defer client.Close()
	ctx := context.Background()
	var duration time.Duration
	if exp != nil {
		duration = *exp
	} else {
		duration = 0
	}
	return client.Set(ctx, key, value, duration).Err()
}
