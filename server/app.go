package main

import (
	"github.com/alfuhigi/news-ajf-sa/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	hd := handlers.NewHandler()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	setupRoute(app, hd)

	app.Listen(":3000")
}

func setupRoute(app *fiber.App, hd *handlers.Handler) {
	app.Get("/", hd.HomePage)
	app.Get("/:id", hd.Post)
}
