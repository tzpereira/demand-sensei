package producer

import (
	"context"
	"log"

	kafkasdk "github.com/tzpereira/go-kafka-sdk/kafka"
)

type Producer struct {
	client *kafkasdk.Producer
}

func NewProducer(
	ctx context.Context,
	brokers []string,
) (*Producer, error) {
	producer, err := kafkasdk.NewProducer(brokers, nil)
	if err != nil {
		return nil, err
	}

	// Delivery handler global 
	producer.StartDeliveryHandler(ctx, func(m *kafkasdk.Message) {
		log.Printf(
			"message delivered | topic=%s partition=%d",
			m.Topic,
			m.Partition,
		)
	})

	return &Producer{
		client: producer,
	}, nil
}

func (p *Producer) Produce(
	ctx context.Context,
	topic string,
	key []byte,
	value []byte,
) error {
	return p.client.Produce(ctx, &kafkasdk.Message{
		Topic: topic,
		Key:   key,
		Value: value,
	})
}

func (p *Producer) Flush(timeoutMs int) {
	p.client.Flush(timeoutMs)
}

func (p *Producer) Close() {
	p.client.Close()
}
