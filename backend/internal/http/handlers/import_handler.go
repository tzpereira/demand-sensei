package handlers

import (
	"github.com/gofiber/fiber/v2"

	"demand-sensei/backend/internal/services"
)

func ImportHandler(svc *services.ImportService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "file is required")
		}

		result, err := svc.Import(file)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(result)
	}
}
