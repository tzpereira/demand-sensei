package events

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// NewProducer creates a Kafka writer for the specified topic
func NewProducer(brokers []string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// Produce sends a message to the specified Kafka topic
func Produce(ctx context.Context, writer *kafka.Writer, key, value []byte) error {
	return writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}
