package consumer

import (
	"context"
	"log"

	kafkasdk "github.com/tzpereira/go-kafka-sdk/kafka"
)

type Consumer struct {
	client *kafkasdk.Consumer
}

func NewConsumer(
	brokers []string,
	groupID string,
	topics []string,
) (*Consumer, error) {
	c, err := kafkasdk.NewConsumer(brokers, groupID, topics, nil)
	if err != nil {
		return nil, err
	}

	return &Consumer{client: c}, nil
}

func (c *Consumer) Consume(
	ctx context.Context,
	handler func(msg *kafkasdk.Message) error,
) error {
	return c.client.Consume(ctx, func(msg *kafkasdk.Message) {
		if err := handler(msg); err != nil {
			log.Printf(
				"consumer handler error | topic=%s partition=%d err=%v",
				msg.Topic,
				msg.Partition,
				err,
			)
		}
	})
}

func (c *Consumer) Close() {
	c.client.Close()
}
