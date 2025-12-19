package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"demand-sensei/backend/internal/services"
)

func ImportHandler(svc *services.ImportService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("Incoming request: file upload")

		file, err := c.FormFile("file")
		if err != nil {
			log.Println("File upload failed: no file provided")
			return fiber.NewError(fiber.StatusBadRequest, "file is required")
		}

		log.Printf("File received: %s (%d bytes)\n", file.Filename, file.Size)

		result, err := svc.Import(file)
		if err != nil {
			log.Printf("File upload failed for %s: %v\n", file.Filename, err)
			return fiber.NewError(fiber.StatusInternalServerError, "failed to upload file")
		}

		log.Printf("File uploaded successfully: %s -> %s (%d bytes)\n", file.Filename, result.Path, result.Size)

		return c.JSON(result)
	}
}
