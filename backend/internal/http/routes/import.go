package routes

import (
	"demand-sensei/backend/internal/http/deps"
	"demand-sensei/backend/internal/http/handlers"
	"demand-sensei/backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterImportRoutes(r fiber.Router, d deps.Deps) {
    service := services.NewImportService(d.Storage, d.Producer)

    importGroup := r.Group("/import")
    importGroup.Post("/sales", handlers.ImportHandler(service, "sales"))
}