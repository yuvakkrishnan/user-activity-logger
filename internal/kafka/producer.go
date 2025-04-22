// Kafka producer
package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// TODO: Implement Kafka producer

func SendMessage(topic, message string) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{Value: []byte(message)},
	)

	if err != nil {
		log.Printf("Kafka write error :%v", err)
	}
}
