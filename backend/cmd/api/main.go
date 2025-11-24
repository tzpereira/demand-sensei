package main

import (
	"log"

	"demand-sensei/backend/internal/http/routes/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.Register(app)

	log.Println("API Gateway running on :9001")
	log.Fatal(app.Listen(":9001"))
}
