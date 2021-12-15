package redis

import (
	"context"
	"fmt"

	"github.com/Telinga-Digital/go-structure/config"
	"github.com/go-redis/redis/v8"
)

func Publish(channel string, payload interface{}) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.GetRedisConfig().Host, config.GetRedisConfig().Port),
		Password: config.GetRedisConfig().Password,
		DB:       0,
	})

	ctx := context.Background()

	err := rdb.Publish(ctx, channel, payload).Err()

	if err != nil {
		panic(err)
	}
}
