package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"strings"
	"time"
)

var MemoryDB *redis.ClusterClient
var ChatBotDB *redis.Client

func RedisInit() {
	MemoryDBInit()
	ChatBotDBInit()
}
var MEMORYDBURL = []string{"clustercfg.mf2-redis.2j4s5t.memorydb.ap-east-1.amazonaws.com:6379"}

var CHATBOTDBURL ="127.0.0.1:6379"

var dummyOption1 = map[string]interface{}{

}
var dummyOption2 = map[string]interface{}{

}
var dummyAction = map[string]interface{}{
	"type": "REPLY",
	"consumer": []string{"customer or message "},
	"payload":  map[string]interface{}{"body": "why why tell me why"},
}

var defualtAction = map[string]interface{}{

}


var dummyFlow = map[string]interface{}{
	"flowName": "flow1",
	"companyId": "matrixsense",
	"length": 1,
	"flow":[]string{
		"actions:tiffany:automations:0:flowKey:0:0:ANY:REPLY",
	},
	"create_at": "0000000000",
	"update_at": "0000000000",
	"create_by": "matrixsense",
	"update_by": "tiffany",
	"default": "actions:tiffany:automations:0:flowKey:0:0:ANY",
	"timeout": map[string]interface{}{"duration": "10s" , "action": []string{
		"actions:tiffany:automations:0:0:0:ANY:REPLY",
	}},
}


func MemoryDBInit(){
	var ctx,cancel = context.WithTimeout(context.Background() ,time.Second*3)

	defer cancel()

	ctx.Done()
	addrs  := strings.Split(os.Getenv("MEMORYDBURL")," ")

	if len(addrs) ==0 || addrs == nil{
		addrs = MEMORYDBURL
	}

	MemoryDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	})

	pong, err := MemoryDB.Ping(ctx).Result()

	log.Println(pong, err)

	err = MemoryDB.Set(ctx, "router", "success", 30).Err()
	if err != nil {
		log.Println("error in redis config,", err)
	}

	val, err := MemoryDB.Get(ctx, "router").Result()
	if err != nil {
		log.Println(err)
	}
	log.Println("MemoryDB: ", val)
	select {
	case <-ctx.Done():
		fmt.Println("MemoryDB timeout")
		fmt.Println(ctx.Err())

	default:
		fmt.Println("MemoryDB connected")
	}

	// set get test

}

func ChatBotDBInit(){
	var ctx,cancel = context.WithTimeout(context.Background() ,time.Second*5)

	defer cancel()

	addr :=os.Getenv("CHATBOTDBURL")

	if  addr == "" {
		addr = CHATBOTDBURL
	}

	ChatBotDB = redis.NewClient(&redis.Options{
		Addr: addr,
	})


	pong, err := ChatBotDB.Ping(ctx).Result()
	log.Println(pong, err)

	// set get test
	err = ChatBotDB.Set(ctx, "router", "success", 30).Err()
	if err != nil {
		log.Println("error in redis config,", err)
	}
	log.Println("set key router")

	val, err := ChatBotDB.Get(ctx, "router").Result()
	if err != nil {
		log.Println(err)
	}
	log.Println("ChatBotDB: ", val)
}
