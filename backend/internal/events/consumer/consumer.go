package events

import (
	"context"

	kafkasdk "github.com/tzpereira/go-kafka-sdk/kafka"
)

type Consumer struct {
    client *kafkasdk.Consumer
}

func NewConsumer(brokers []string, topic, groupID string) (*Consumer, error) {
    cfg := &kafkasdk.Config{
        Brokers: brokers,
        Topics:  []string{topic},
        GroupID: groupID,
    }

    c, err := kafkasdk.NewConsumer(cfg)
    if err != nil {
        return nil, err
    }

    return &Consumer{client: c}, nil
}

func (c *Consumer) Consume(ctx context.Context, handler func([]byte) error) error {
    return c.client.Consume(ctx, handler)
}
