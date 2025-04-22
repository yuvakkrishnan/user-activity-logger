// Kafka consumer
package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// TODO: Implement Kafka consumer
func Startconsumer(topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		GroupID: "user-activity-group",
	})

	for {

		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading messagge :%v", err)
			break
		}
		log.Printf("Message received %s", string(m.Value))
	}

}
