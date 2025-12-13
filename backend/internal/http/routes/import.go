package routes

import (
	"github.com/gofiber/fiber/v2"

	"demand-sensei/backend/internal/http/handlers"
)

func RegisterImportRoutes(r fiber.Router) {
	r.Post("/import", handlers.ImportHandler)
}