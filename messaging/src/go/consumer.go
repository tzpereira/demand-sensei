package messaging

import (
	"context"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// BlockForever is a constant used to indicate that the consumer should block indefinitely when reading messages.
const BlockForever = -1

// NewConsumer creates and subscribes a Kafka consumer using the provided brokers, groupID, topics, and config.
// Returns a configured *ckafka.Consumer or an error if creation or subscription fails.
func NewConsumer(brokers []string, groupID string, topics []string, config map[string]string) (*ckafka.Consumer, error) {
	conf := &ckafka.ConfigMap{
		"bootstrap.servers": brokers[0],
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	}
	for k, v := range config {
		(*conf)[k] = v
	}
	consumer, err := ckafka.NewConsumer(conf)
	if err != nil {
		return nil, err
	}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

// Consume reads messages from the Kafka consumer and calls the handler for each message.
// It blocks until the context is cancelled or an error occurs.
// Returns ctx.Err() if the context is cancelled, or the error from consumer.ReadMessage if message consumption fails.
func Consume(ctx context.Context, consumer *ckafka.Consumer, handler func(msg *ckafka.Message)) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			msg, err := consumer.ReadMessage(-1)
			if err != nil {
				return err
			}
			handler(msg)
		}
	}
}
