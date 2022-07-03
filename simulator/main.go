package main

import (
	"fmt"
	"log"
	kafkaProduce "repo_/jl-fb/imersaofsfc-simulator/application/kafka"
	"repo_/jl-fb/imersaofsfc-simulator/infra/kafka"

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
	/* 	r := route.Route{
	   		ID:       "1",
	   		ClientID: "1",
	   	}
	   	r.LoadPositions()
	   	jsonroute, err := r.ExportJSONPositions()
	   	if err != nil {
	   		fmt.Println(err)
	   	}
	   	fmt.Println(jsonroute[0]) */

	//	producer := kafka.NewKafkaProducer()
	//	kafka.Publish("Olá", "readtest", producer)

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		go kafkaProduce.Produce(msg)
		fmt.Println(string(msg.Value))
	}
}
