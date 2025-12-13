package events

import (
	"context"

	kafkasdk "github.com/tzpereira/go-kafka-sdk/kafka"
)

type Producer struct {
	client *kafkasdk.Producer
	topic  string
}

func NewProducer(brokers []string, topic string) (*Producer, error) {
	producer, err := kafkasdk.NewProducer(brokers, nil)
	if err != nil {
		return nil, err
	}
	return &Producer{client: producer, topic: topic}, nil
}

func (p *Producer) Produce(ctx context.Context, value []byte) error {
	msg := &kafkasdk.Message{
		Topic: p.topic,
		Value: value,
	}
	return p.client.Produce(ctx, msg)
}

func (p *Producer) StartDeliveryHandler(ctx context.Context, handler func(m *kafkasdk.Message)) {
	p.client.StartDeliveryHandler(ctx, handler)
}

func (p *Producer) Flush(timeoutMs int) {
	p.client.Flush(timeoutMs)
}

func (p *Producer) Close() {
	p.client.Close()
}
