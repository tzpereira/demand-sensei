package events

import (
	"context"

	kafkasdk "github.com/tzpereira/go-kafka-sdk/kafka"
)

type Consumer struct {
	client *kafkasdk.Consumer
}

func NewConsumer(brokers []string, groupID string, topics []string) (*Consumer, error) {
	c, err := kafkasdk.NewConsumer(brokers, groupID, topics, nil)
	if err != nil {
		return nil, err
	}
	return &Consumer{client: c}, nil
}

func (c *Consumer) Consume(ctx context.Context, handler func(msg *kafkasdk.Message)) error {
	return c.client.Consume(ctx, handler)
}

func (c *Consumer) Close() {
	c.client.Close()
}
