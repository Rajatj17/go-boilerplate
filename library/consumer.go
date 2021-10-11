package library

import (
	"encoding/json"
	"fmt"

	"github.com/adjust/rmq/v4"
)

type emailPayload struct {
	to      string
	subject string
	body    string
}

func SendEmail(delivery rmq.Delivery) {
	var task interface{}
	json.Unmarshal([]byte(delivery.Payload()), &task)
	fmt.Print("Hey I'm here in send email function!", task)

	delivery.Ack()
}

var Consumers = []func(rmq.Delivery){SendEmail}
