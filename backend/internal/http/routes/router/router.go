package router

import (
	"demand-sensei/backend/internal/http/deps"
	"demand-sensei/backend/internal/http/routes"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, d deps.Deps) {
    api := app.Group("/api/v1")

    routes.RegisterHealthRoutes(api)
    routes.RegisterImportRoutes(api, d)
}