package routes

import (
	"github.com/gofiber/fiber/v2"

	"demand-sensei/backend/internal/http/handlers"
)

func Register(r fiber.Router) {
	r.Get("/health", handlers.HealthHandler)
}
