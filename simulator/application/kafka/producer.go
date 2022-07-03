package kafka

import (
	"encoding/json"
	"log"
	"os"
	"repo_/jl-fb/imersaofsfc-simulator/application/route"
	"repo_/jl-fb/imersaofsfc-simulator/infra/kafka"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJSONPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
