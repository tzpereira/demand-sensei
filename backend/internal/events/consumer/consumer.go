package events

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// NewConsumer creates a Kafka reader for the specified topic
func NewConsumer(brokers []string, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
	})
}

// Consume reads messages from the topic and executes the handler
func Consume(ctx context.Context, reader *kafka.Reader, handler func(msg kafka.Message)) error {
	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		handler(m)
	}
}
