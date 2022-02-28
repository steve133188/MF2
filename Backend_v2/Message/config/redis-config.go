package config

import (
	"context"
	"crypto/tls"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var ClusterClient *redis.ClusterClient

var TestClient *redis.Client

func RedisInit() {
	var ctx = context.Background()

	ClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{GoDotEnvVariable("REDISURL")},
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	})
	pong, err := ClusterClient.Ping(ctx).Result()
	log.Println(pong, err)

	// set get test
	err = ClusterClient.Set(ctx, "message", "success", 30*time.Second).Err()
	if err != nil {
		log.Println("error in redis config,", err)
	}
	log.Println("set key message")

	val, err := ClusterClient.Get(ctx, "message").Result()
	if err != nil {
		log.Println(err)
	}
	log.Println("message: ", val)
}

func TestRedis() {
	var ctx = context.Background()

	TestClient = redis.NewClient(&redis.Options{
		// Addr:     "redis-master-sr.default.svc.cluster.local:6379",
		Addr: "localhost:6379",
	})
	pong, err := TestClient.Ping(ctx).Result()
	log.Println(pong, err)

}
