package main

import (
	"github.com/alfuhigi/news-ajf-sa/db"
	"github.com/alfuhigi/news-ajf-sa/handlers"
	"github.com/alfuhigi/news-ajf-sa/providers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	dbConn := providers.Connect()
	entiry := db.NewEntiry(dbConn)
	hd := handlers.NewHandler(entiry)
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	setupRouter(app, hd)

	app.Listen(":3000")
}

func setupRouter(app *fiber.App, hd *handlers.Handler) {
	app.Get("/", hd.HomePage)
	app.Get("/:id", hd.Post)
}
