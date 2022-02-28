package main

import (
	"ChatbotAPI/config"
	"ChatbotAPI/hanlder"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"net/http"
	"strconv"
)


func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	config.RedisInit()

	go func() {

		messagesSub := config.ChatBotDB.Subscribe(context.Background(),"messages.received.WABA" ,"bot")

		payload := make(map[string]interface{})

		actions :=make([]interface{} ,0)

		//botPip := config.ChatBotDB.Pipeline()

		for msg := range messagesSub.Channel() {
			flow := make( map[string]interface{})

			// #1 Unmarshall data
			if err := json.Unmarshal([]byte(msg.Payload), &payload);err != nil{

				fmt.Printf("Marshall failed=%s\n",err)

			}

			if msg.Channel =="bot"{
				fmt.Println("bot : " , msg.Payload)
			}
			//--------------------------------------- messages.received------------------------------------------------------------
			if msg.Channel =="messages.received.WABA"{
				chatListItems := make(map[string]string)
				//flow := make(map[string]string)

				//#2 check chatlist
				searchKey := fmt.Sprintf("botBucket:%s:%s",payload["channel"],payload["room_id"])

				chatListItems ,err :=config.ChatBotDB.HGetAll(context.Background() , searchKey ).Result()
				if err != nil || len(chatListItems)==0{
					//#2.1 if chatlist not exist check the channel's flow
					log.Println("get ChatList err : " ,err)

					flowKey := fmt.Sprintf("flows:*:%s:default"  , payload["channel"])

					flows ,err :=config.ChatBotDB.Keys(context.Background() , flowKey).Result()

					if err != nil {
						fmt.Println("no chat bot exist")
						continue
					}

					flowKey = flows[0]

					val ,err :=config.ChatBotDB.Get(context.Background() , flowKey ).Result()
					if err!= nil{
						fmt.Println("no chat bot exist")
						continue
					}
					if err := json.Unmarshal([]byte(val), &flow);err != nil{

						fmt.Printf("Marshal flow failed=%s\n",err)

					}
					fmt.Println("flow" , flow)

					action := make(map[string]interface{})

					actionKey := fmt.Sprintf("%s" , flow["default"])

					fmt.Println(flow["default"])

					fmt.Println("action key = " , actionKey)

					val ,err =config.ChatBotDB.Get(context.Background() , actionKey ).Result()
					if err!= nil{
						fmt.Println("the flow no default action exist")
						continue
					}

					if err := json.Unmarshal([]byte(val), &action);err != nil{

						fmt.Printf("Marshal action failed=%s\n",err)
						continue

					}

					mP := action["payload"].(map[string]interface{})
					pP := payload["sender"].(string)
					mP["recipientId"] = pP
					fmt.Println("message payload : " ,mP)

					botPayload , err := json.Marshal(mP)
					config.ChatBotDB.Publish(context.Background() ,"bot" , botPayload)


					resp , err :=sendMessage(botPayload)
					if err != nil{
						fmt.Println(err)
					}
					fmt.Println(resp)

					chatListKey := fmt.Sprintf("botBucket:%s:%s" , payload["channel"] , payload["room_id"])
					
					chatListPayload := make(map[string]interface{})
					
					chatListPayload["stage"] = 1
					chatListPayload["parentKey"] = "0"
					chatListPayload["flowKey"] = flowKey
					chatListPayload["flowLength"] = flow["length"]
					fmt.Println("flow : " , flow)

					config.ChatBotDB.HMSet(context.Background(),chatListKey ,chatListPayload )

					chatListItems2 := make(map[string]string)

					chatListItems2 , err = config.ChatBotDB.HGetAll(context.Background(),chatListKey  ).Result()

					if err !=nil{
						fmt.Println("cannot get the chatList")
					}
					
					fmt.Println("chatList items : " ,chatListItems2)

					config.ChatBotDB.Publish(context.Background() ,"bot" , botPayload)

					continue
				}
				fmt.Println("chatlist 1 " , chatListItems)

				optionKey := fmt.Sprintf("automations:%s*#%s*" ,chatListItems["stage"],payload["body"] )

				fmt.Println("opt key L ",optionKey)

				optionKeys ,err :=config.ChatBotDB.Keys(context.Background() , optionKey).Result()

				if err != nil {
					fmt.Println("no chat bot exist")
					continue
				}
				fmt.Println("optionkeys : ", optionKeys)
				optionKey = optionKeys[0]

				option := make([]string , 0)

				val ,err := config.ChatBotDB.Get(context.Background() , optionKey).Result()

				if err = json.Unmarshal([]byte(val) , &option);err != nil{
					fmt.Println("unmarshal option fail" ,err )
				}

				fmt.Println("action :" , option)

				for  _,v  :=range option{
					action := make(map[string]interface{})

					val , err := config.ChatBotDB.Get(context.Background() , v).Result()
					if err != nil || len(val) ==0{
						fmt.Println("get action fail or something went wrong")
					}
					if err := json.Unmarshal([]byte(val), &action);err != nil{

						fmt.Printf("Marshal action failed=%s\n",err)
						continue

					}

					mP := action["payload"].(map[string]interface{})
					pP := payload["sender"].(string)
					mP["recipientId"] = pP
					fmt.Println("message payload : " ,mP)

					botPayload , err := json.Marshal(mP)
					config.ChatBotDB.Publish(context.Background() ,"bot" , botPayload)

					resp , err :=sendMessage(botPayload)
					if err != nil{
						fmt.Println(err)
					}
					fmt.Println(resp)
				}

				fmt.Println("chatList" , chatListItems)


				if err != nil {

					log.Println("error in redis config,", err)

				}
				// TODO #4  publish the actions force
				for v := range actions{
					fmt.Println(v)
					//	TODO send out the message
				}
				// TODO #5  update the ChatList
				stage , err := config.ChatBotDB.HIncrBy(context.Background() ,searchKey,"stage", 1) .Result()
				if err != nil {
					fmt.Println("add stage fail")
				}
				fmt.Println(searchKey)
				fmt.Println("updated stage to :" , stage)
				fmt.Println("flow length:" , chatListItems["flowLength"])

				flowLength , err := strconv.ParseInt(chatListItems["flowLength"] , 32 , 64)
				if err != nil {
					fmt.Println("flow")
				}

				if stage > flowLength{
					_ ,err := config.ChatBotDB.Del(context.Background() , searchKey).Result()
					if err != nil{
						fmt.Println("delete chatlist fail :" , searchKey)
					}
				}

				chatListItems , err = config.ChatBotDB.HGetAll(context.Background(),searchKey  ).Result()
				if err !=nil{
					fmt.Println("cannot get the chatList")
				}

				fmt.Println("chatList items : " ,chatListItems)
			}



		}

	}()


	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Chat Bot API Server is running")
	})

	api := app.Group("/api")

	api.Post("/flow", hanlder.CreateFlow)
	api.Post("/action", hanlder.CreateAction)
	api.Post("/option", hanlder.CreateOption)

	app.Post("/" , func(c *fiber.Ctx) error {

		//data := new(map[string]interface{})
		//
		//
		//
		//if err := c.BodyParser(data);err !=nil{
		//	panic(err)
		//}
		//
		//payload , err := json.Marshal(data)
		//if err != nil {
		//	panic(err)
		//}

		config.ChatBotDB.Publish(context.Background() ,"messages.received.WABA" , c.Body())

		return c.Status(http.StatusOK).JSON(c.Body())
	})

	app.Listen(":3010")

}

func sendMessage( value[]byte) (resp *http.Response , err error) {

	//botPayload , err := json.Marshal()

	url :="http://k8s-channelr-channelr-90ec0dd324-309508086.ap-east-1.elb.amazonaws.com/ch-router/send-message?cid=tiffany&cname=WABA"

	resp, err = http.Post(url, "application/json", bytes.NewBuffer(value))

	return resp ,err
}