package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"demand-sensei/backend/internal/events/producer"
	"demand-sensei/backend/internal/http/deps"
	"demand-sensei/backend/internal/http/routes/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()

	brokers := os.Getenv("KAFKA_BROKERS")

	kafkaProducer, err := producer.NewProducer(
		ctx,
		[]string{brokers},
	)

	if err != nil {
		log.Fatalf("failed to init kafka producer: %v", err)
	}
	defer kafkaProducer.Close()

	app := fiber.New()

	router.Register(app, deps.Deps{
		Producer: kafkaProducer,
	})

	log.Println("API running on :9001")
	if err := app.Listen(":9001"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
