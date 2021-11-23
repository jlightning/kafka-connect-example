package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type simpleMessage struct {
	Payload struct {
		Before map[string]interface{} `json:"before"`
		After  map[string]interface{} `json:"after"`
	} `json:"payload"`
}

func main() {
	fmt.Println("SERVER STARTED")

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "foo",
		"auto.offset.reset": "smallest",
	})
	checkErr(err)

	defer consumer.Close()

	err = consumer.Subscribe("customers", nil)
	checkErr(err)

	run := true
	for run == true {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			var sm simpleMessage
			err := json.Unmarshal(e.Value, &sm)
			checkErr(err)
			fmt.Println(sm.Payload)
		case kafka.Error:
			if e.Code() != kafka.ErrUnknownTopicOrPart {
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				run = false
			}
		default:
			//fmt.Printf("Ignored %v\n", e)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
