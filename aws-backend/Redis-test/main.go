package main

import (
	"ChatMessage_Fiber/config"
	"context"
	"fmt"
)

func main() {
	config.RedisInit()
	var ctx = context.Background()

	redisSub := config.RedisClient.Subscribe(ctx, "chatrooms")

	for {
		err := config.RedisClient.Publish(ctx, "chatrooms", "testing").Err()
		if err != nil {
			fmt.Println(err)
		}

		message, err := redisSub.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(message)

	}

}
