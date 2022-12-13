package main

import (
	"github.com/VishvaNavanjana/Docker-Go-REST-API/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
    app.Get("/", handlers.ListFacts)

    app.Post("/fact", handlers.CreateFact)
}