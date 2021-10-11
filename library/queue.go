package library

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/adjust/rmq/v4"
	"main.go/config"
)

var Queue rmq.Queue

func CreateQueue() {
	connection, err1 := rmq.OpenConnection(
		config.AppConfig.QueueConfig.ConnectionName,
		config.AppConfig.QueueConfig.Protocol,
		config.AppConfig.QueueConfig.Url,
		1,
		make(chan<- error),
	)
	if err1 != nil {
		panic(fmt.Sprintf("Error in establishing queue connection %s", err1.Error()))
	}

	defaultQueue, err2 := connection.OpenQueue("Default")
	if err2 != nil {
		panic("Error in creating queue")
	}

	defaultQueue.StartConsuming(10, time.Second)

	for i := 0; i < len(Consumers); i++ {
		defaultQueue.AddConsumerFunc(fmt.Sprintf("consumer=%d", i), Consumers[i])
		print(i)
	}

	Queue = defaultQueue
}

func Dispatch(payload interface{}) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic("Hey something is wrong!")
	}

	Queue.PublishBytes(payloadBytes)
}
