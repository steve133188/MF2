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
//"matrixsense:automations:1:0:#1"
var dummyOption1 = []string{
	"matrixsense:actions:00000000002",
}
//"matrixsense:automations:1:0:#2"
var dummyOption2 = []string{
	"matrixsense:actions:00000000003",
}
//"matrixsense:automations:1:0:#3"
var dummyOption3 = []string{
	"matrixsense:actions:00000000004",
}

//"matrixsense:actions:00000000004"
var dummyReplyAction2 = map[string]interface{}{
	"type": "REPLY",
	"consumer": "customer or message ",
	"payload":  map[string]interface{}{"body": "Lun Height :183"},
}
//"matrixsense:actions:00000000003"
var dummyReplyAction3 = map[string]interface{}{
	"type": "REPLY",
	"consumer": "customer or message ",
	"payload":  map[string]interface{}{"body": "Ben Height : 170"},
}
//"matrixsense:actions:00000000002"
var dummyReplyAction4 = map[string]interface{}{
	"type": "REPLY",
	"consumer": "customer or message ",
	"payload":  map[string]interface{}{"body": "Steve Height :180"},
}
//"matrixsense:actions:00000000000"
var errorHandleAction = map[string]interface{}{
	"type": "REPLY",
	"consumer": "message ",
	"payload":  map[string]interface{}{"body": "Sorry I dont understand please try agein"},
}

//"matrixsense:actions:0000000000001"
var defaultAction = map[string]interface{}{
	"type": "REPLY",
	"consumer": "message ",
	"payload":  map[string]interface{}{"body": "Hi I'm MF Assistant \n What can I help you? \n  Please type and rely following number of items \n 1. How tall is Steve\n 1. How tall is Ben\n1. How tall is Lun\n"},
}

//"matrixsense:flows:0000000000000"
var dummyFlow = map[string]interface{}{
	"flowName": "flow1",
	"companyId": "matrixsense",
	"length": 1,
	"flow":[][]string{
		[]string{
			"matrixsense:automations:1:0:#1",
			"matrixsense:automations:1:0:#2",
			"matrixsense:automations:1:0:#3",
		},
	},
	"create_at": "0000000000",
	"update_at": "0000000000",
	"create_by": "matrixsense",
	"update_by": "tiffany",
	"default": "matrixsense:actions:0000000000001",
	"timeout": map[string]interface{}{"duration": "10s" , "action": []string{
		"tiffany:actions:00000000000",
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
