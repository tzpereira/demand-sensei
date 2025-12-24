package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"demand-sensei/backend/internal/events/consumer"

	kafkasdk "github.com/tzpereira/go-kafka-sdk/kafka"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	brokers := []string{"localhost:9092"}
	groupID := "imports-processor"
	topics := []string{"imports.created"}

	c, err := consumer.NewConsumer(
		brokers,
		groupID,
		topics,
	)
	if err != nil {
		log.Fatalf("failed to init consumer: %v", err)
	}
	defer c.Close()

	log.Println("consumer started")

	err = c.Consume(ctx, func(msg *kafkasdk.Message) error {
		log.Printf(
			"received | topic=%s partition=%d value=%s",
			msg.Topic,
			msg.Partition,
			string(msg.Value),
		)
		return nil
	})

	if err != nil && err != context.Canceled {
		log.Fatalf("consumer stopped with error: %v", err)
	}
}
