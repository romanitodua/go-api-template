package main

import (
	"github.com/gofiber/fiber/v3"
	"go-api-template/handlers"
)

func initRoutes(app *fiber.App, h *handlers.Handler) {
	// register routes here
	app.Get("/status", h.Status)

}
