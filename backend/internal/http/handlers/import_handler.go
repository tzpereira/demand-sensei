package handlers

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const TEMP_UPLOAD_DIR = "/app/data/uploads"

func ImportHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File is required",
		})
	}

	uploadDir := TEMP_UPLOAD_DIR
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create upload directory",
		})
	}

	filePath := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	return c.JSON(fiber.Map{
		"message":  "file uploaded successfully",
		"filename": file.Filename,
		"size":     file.Size,
		"path":     filePath,
	})
}
