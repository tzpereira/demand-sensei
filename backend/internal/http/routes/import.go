package routes

import (
	"github.com/gofiber/fiber/v2"

	"demand-sensei/backend/internal/http/handlers"
	"demand-sensei/backend/internal/services"
	"demand-sensei/backend/internal/storage"
)

func RegisterImportRoutes(r fiber.Router) {
	storage := storage.NewLocalStorage()
	service := services.NewImportService(storage)

	r.Post("/import", handlers.ImportHandler(service))
}