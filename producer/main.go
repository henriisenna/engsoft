package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	evt "event"
	"github.com/go-redis/redis/v7"
	"utils"
)

var (
	streamName string = os.Getenv("STREAM")
	client     *redis.Client
)

func init() {
	var err error
	client, err = utils.NewRedisClient()
	if err != nil {
		panic(err)
	}
}

func main() {
	generateEvent()
}

func generateEvent() {
	var userID uint64 = 0
	for i := 0; i < 10; i++ {

		userID++ //uint64(rand.Intn(1000))

		eventType := []evt.Type{evt.InvestimentoType, evt.InvestidorType}[rand.Intn(2)]

		if eventType == evt.InvestimentoType {

			newID, err := produceMsg(map[string]interface{}{
				"type": string(eventType),
				"data": &evt.InvestimentoEvent{
					Base: &evt.Base{
						Type:     eventType,
						DateTime: time.Now(),
					},
					UserID: userID,
				},
			})

			checkError(err, newID, string(eventType), userID)

		} else {

			investidor := []string{"José", "Henrique", "Thays"}[rand.Intn(3)]

			newID, err := produceMsg(map[string]interface{}{
				"type": string(eventType),
				"data": &evt.InvestidorEvent{
					Base: &evt.Base{
						Type:     eventType,
						DateTime: time.Now(),
					},
					UserID:     userID,
					investidor: investidor,
				},
			})

			checkError(err, newID, string(eventType), userID, investidor)
		}

	}
}

func produceMsg(event map[string]interface{}) (string, error) {

	return client.XAdd(&redis.XAddArgs{
		Stream: streamName,
		Values: event,
	}).Result()
}

func checkError(err error, newID, eventType string, userID uint64, investidor ...string) {
	if err != nil {
		fmt.Printf("produce event error:%v\n", err)
	} else {

		if len(investidor) > 0 {
			fmt.Printf("produce event success Type:%v UserID:%v investidor:%v offset:%v\n",
				string(eventType), userID, investidor, newID)
		} else {
			fmt.Printf("produce event success Type:%v UserID:%v offset:%v\n",
				string(eventType), userID, newID)
		}

	}
}
