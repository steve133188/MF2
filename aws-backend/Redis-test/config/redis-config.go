package config

import (
	"context"
	"crypto/tls"
	"github.com/go-redis/redis/v8"
	"log"
)

var RedisClient *redis.Client

func RedisInit() {
	var ctx = context.Background()

	RedisClient = redis.NewClient(&redis.Options{
		Addr: GoDotEnvVariable("REDISURL"),
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	})

	pong, err := RedisClient.Ping(ctx).Result()
	log.Println(pong, err)

	// set get test
	err = RedisClient.Set(ctx, "router", "success", 30).Err()
	if err != nil {
		log.Println("error in redis config,", err)
	}
	log.Println("set key router")

	val, err := RedisClient.Get(ctx, "router").Result()
	if err != nil {
		log.Println(err)
	}
	log.Println("test: ", val)
}
