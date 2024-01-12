package http

import (
	"healthchecker/http/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/v1")
	api.Post("/ping-urls", handlers.PingURLs)
}