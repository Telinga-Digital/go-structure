package redis

import (
	"context"
	"fmt"

	"github.com/Telinga-Digital/go-structure/config"
	ExampleJob "github.com/Telinga-Digital/go-structure/jobs/example_job"
	"github.com/go-redis/redis/v8"
)

func Subscribe() {
	if config.GetRedisConfig().Enabled == "true" {
		rdb := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", config.GetRedisConfig().Host, config.GetRedisConfig().Port),
			Password: config.GetRedisConfig().Password,
			DB:       0,
		})

		ctx := context.Background()

		subscribe := rdb.Subscribe(ctx, "go-structure", ExampleJob.Channel)

		// defer subscribe.Close()

		ch := subscribe.Channel()

		go func() {
			for msg := range ch {
				if msg.Channel == ExampleJob.Channel {
					ExampleJob.Handle(msg.Payload)
				}
			}
		}()

		fmt.Println("Subscribed to redis")
	}
}
