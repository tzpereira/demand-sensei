package routes

import (
	"github.com/gofiber/fiber/v2"

	"demand-sensei/backend/internal/http/deps"
	"demand-sensei/backend/internal/http/handlers"
	"demand-sensei/backend/internal/services"
)

func RegisterImportRoutes(r fiber.Router, d deps.Deps) {
	service := services.NewImportService(d.Storage, d.Producer)

	r.Post("/import", handlers.ImportHandler(service))
}
