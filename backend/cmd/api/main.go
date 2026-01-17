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
	"demand-sensei/backend/internal/storage"

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

	s3Storage, err := storage.NewS3CompatibleStorage(
		os.Getenv("S3_ENDPOINT"),
		os.Getenv("S3_ACCESS_KEY"),
		os.Getenv("S3_SECRET_KEY"),
		os.Getenv("S3_BUCKET"),
		os.Getenv("S3_BASE_PATH"),
		os.Getenv("S3_USE_SSL") == "true",
	)
	if err != nil {
		log.Fatalf("failed to init s3 storage: %v", err)
	}

	app := fiber.New(fiber.Config{
		StrictRouting: false,
		CaseSensitive: false,
	})

	router.Register(app, deps.Deps{
		Producer: kafkaProducer,
		Storage:  s3Storage,
	})

	log.Println("API running on :9001")
	if err := app.Listen(":9001"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
