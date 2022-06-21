package main

import (
	"fmt"
	"log"

	kafka2 "github.com/LeonardoGregoriocs/code_delivery/application/kafka"
	"github.com/LeonardoGregoriocs/code_delivery/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		go kafka2.Produce(msg)
		fmt.Println(string(msg.Value))
	}

	// route := route2.Route{
	// 	ID:       "1",
	// 	ClientID: "2",
	// }

	// route.LoadPositions()
	// stringjson, _ := route.ExportJsonPositions()

	// for i, j := range stringjson {
	// 	fmt.Println(i, j)
	// }
}
