package config

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"strings"
	"time"
)

var MemoryDB *redis.ClusterClient
var ChatBotDB *redis.ClusterClient

type RedisClient struct {
	MemoryDB *redis.ClusterClient
	ChatBotDB *redis.ClusterClient
}

type RedisServices interface {
	RedisInit()

}


func RedisInit() {
	//MemoryDBInit()
	ChatBotDBInit()
	SaveDummy()
}

var MEMORYDBURL = []string{"clustercfg.mf2-redis.2j4s5t.memorydb.ap-east-1.amazonaws.com:6379"}

var CHATBOTDBURL = "127.0.0.1:6379"

//"automations:1:0:#1"
var dummyOption1 = []string{
	"actions:00000000002",
}

//"automations:1:0:#2"
var dummyOption2 = []string{
	"actions:00000000003",
}

//"automations:1:0:#3"
var dummyOption3 = []string{
	"actions:00000000004",
}

//"actions:00000000004"
var dummyReplyAction2 = map[string]interface{}{
	"type":     "REPLY",
	"consumer": "customer or message ",
	"payload":  map[string]interface{}{"text": "Lun Height :183" , "type":"TEXT"},
}

//"actions:00000000003"
var dummyReplyAction3 = map[string]interface{}{
	"type":     "REPLY",
	"consumer": "customer or message ",
	"payload":  map[string]interface{}{"text": "Ben Height : 170", "type":"TEXT"},
}

//"actions:00000000002"
var dummyReplyAction4 = map[string]interface{}{
	"type":     "REPLY",
	"consumer": "customer or message ",
	"payload":  map[string]interface{}{"text": "Steve Height :180", "type":"TEXT"},
}

//"actions:00000000000"
var errorHandleAction = map[string]interface{}{
	"type":     "REPLY",
	"consumer": "message ",
	"payload":  map[string]interface{}{"text": "Sorry I dont understand please try again" , "type":"TEXT"},
}

//"actions:0000000000001"
var defaultAction = map[string]interface{}{
	"type":     "REPLY",
	"consumer": "message ",
	"payload":  map[string]interface{}{"text": "Hi I'm MF Assistant \nWhat can I help you? \nPlease type and rely following number of items\n1. How tall is Steve\n2. How tall is Ben\n3. How tall is Lun\n", "type":"TEXT"},
}

//"flows:0000000000000"
var dummyFlow = map[string]interface{}{
	"flowName":  "flow1",
	"companyId": "matrixsense",
	"length":    1,
	"flow": [][]string{
		[]string{
		"automations:1:0:#1",
		"automations:1:0:#2",
		"automations:1:0:#3",
		},
	},
	"create_at": "0000000000",
	"update_at": "0000000000",
	"create_by": "matrixsense",
	"update_by": "tiffany",
	"default":   "actions:0000000000001",
	"timeout": map[string]interface{}{"duration": "10s", "actions": []string{
		"tiffany:actions:00000000000",
	}},
}

func SaveDummy() {

	pipe := ChatBotDB.Pipeline()
	//flow
	flow, _ := json.Marshal(dummyFlow)
	pipe.Set(context.Background(), "flows:0000000000000:WABA:default", flow , 0)

	//option
	option1, _ := json.Marshal(dummyOption1)
	pipe.Set(context.Background(), "automations:1:0:#1", option1, 0)
	option2, _ := json.Marshal(dummyOption2)
	pipe.Set(context.Background(), "automations:1:0:#2", option2, 0)
	option3, _ := json.Marshal(dummyOption3)
	pipe.Set(context.Background(), "automations:1:0:#3", option3, 0)

	//action
	defaultaction, _ := json.Marshal(defaultAction)
	pipe.Set(context.Background(), "actions:0000000000001", defaultaction, 0)

	errcation, _ := json.Marshal(errorHandleAction)
	pipe.Set(context.Background(), "actions:00000000000", errcation, 0)

	action2, _ := json.Marshal(dummyReplyAction2)
	pipe.Set(context.Background(), "actions:00000000004", action2, 0)
	action3, _ := json.Marshal(dummyReplyAction3)
	pipe.Set(context.Background(), "actions:00000000003", action3, 0)
	action4, _ := json.Marshal(dummyReplyAction4)
	pipe.Set(context.Background(), "actions:00000000002", action4, 0)

	_, err := pipe.Exec(context.Background())
	if err != nil {
		fmt.Println(err)
	}

}

func MemoryDBInit() {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()

	ctx.Done()
	addrs := strings.Split(os.Getenv("MEMORYDBURL"), " ")

	if len(addrs) == 0 || addrs == nil {
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

func ChatBotDBInit() {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	//addr := os.Getenv("CHATBOTDBURL")
	addrs := MEMORYDBURL


	//ChatBotDB = redis.NewClient(&redis.Options{
	//	Addr: addr,
	//})
	//addrs := strings.Split(os.Getenv("MEMORYDBURL"), " ")

	//if len(addrs) == 0 || addrs == nil {
	//	addrs = MEMORYDBURL
	//}

	ChatBotDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
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
