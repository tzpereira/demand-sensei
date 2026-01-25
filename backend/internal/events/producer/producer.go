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

	// Global delivery handler 
	producer.StartDeliveryHandler(ctx, func(m *kafkasdk.Message) {
		importType := string(m.Key)

		log.Printf(
			"message delivered key=%s topic=%s partition=%d value=%s\n",
			importType,
			m.Topic,
			m.Partition,
			string(m.Value),
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
