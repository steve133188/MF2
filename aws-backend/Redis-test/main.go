package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Channel struct {
	ID      string `json:"id" bson:"id"`
	Address string `json:"address" bson:"address"`
}

func main() {
	var ctx = context.Background()
	var RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	redisPubsub := RedisClient.PSubscribe(ctx, "__keyspace@0__:*")

	for {
		message, err := redisPubsub.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(message)

	}

}
