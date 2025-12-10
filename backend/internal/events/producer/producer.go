package events

import (
	"context"

	kafkasdk "github.com/tzpereira/go-kafka-sdk/kafka"
)

type Producer struct {
	client *kafkasdk.Producer
}

func NewProducer(brokers []string, topic string) (*Producer, error) {
	p, err := kafkasdk.NewProducer(brokers, topic)
	if err != nil {
		return nil, err
	}

	return &Producer{client: p}, nil
}

func (p *Producer) Produce(ctx context.Context, key, value []byte) error {
	return p.client.Produce(ctx, key, value)
}
