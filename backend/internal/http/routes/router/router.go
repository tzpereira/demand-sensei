package router

import (
	"github.com/gofiber/fiber/v2"

	routes "demand-sensei/backend/internal/http/routes"
)

func Register(app *fiber.App) {
    api := app.Group("/api/v1")

    routes.Register(api)
}
