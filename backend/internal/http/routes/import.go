package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"demand-sensei/backend/internal/http/handlers"
	"demand-sensei/backend/internal/services"
	"demand-sensei/backend/internal/storage"
)

func RegisterImportRoutes(r fiber.Router) {
	var st storage.Storage
	var err error

	st, err = storage.NewS3CompatibleStorage(
		os.Getenv("S3_ENDPOINT"),
		os.Getenv("S3_ACCESS_KEY"),
		os.Getenv("S3_SECRET_KEY"),
		os.Getenv("S3_BUCKET"),
		os.Getenv("S3_BASE_PATH"),
		os.Getenv("S3_USE_SSL") == "true",
	)

	if err != nil {
		panic(err)
	}

	service := services.NewImportService(st)
	r.Post("/import", handlers.ImportHandler(service))
}