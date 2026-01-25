package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"demand-sensei/backend/internal/http/validators"
	"demand-sensei/backend/internal/services"
)

func ImportHandler(svc *services.ImportService, importType string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("Incoming request: file upload")

		file, err := c.FormFile("file")
		if err != nil {
			log.Println("File upload failed: no file provided")
			return fiber.NewError(fiber.StatusBadRequest, "file is required")
		}

		log.Printf("File received: %s (%d bytes)\n", file.Filename, file.Size)

		validator := validators.GetValidator(importType)
		if err := validator(file); err != nil {
			log.Printf("Validation failed for %s: %v\n", file.Filename, err)
			return fiber.NewError(fiber.StatusBadRequest, "invalid file format: "+err.Error())
		}

		result, err := svc.Import(file, importType)
		if err != nil {
			log.Printf("File upload failed for %s: %v\n", file.Filename, err)
			return fiber.NewError(fiber.StatusInternalServerError, "failed to upload file")
		}

		log.Printf("File uploaded successfully: %s -> %s (%d bytes)\n", file.Filename, result.Path, result.Size)

		return c.JSON(result)
	}
}
