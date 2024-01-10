package http

import "github.com/gofiber/fiber/v2"

func setupRoutes(app *fiber.App) {
	api := app.Group("/v1")
	api.Get("/ping-urls")
}