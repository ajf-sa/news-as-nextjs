package main

import (
	"os"

	"github.com/alfuhigi/news-ajf-sa/db"
	"github.com/alfuhigi/news-ajf-sa/handlers"
	"github.com/alfuhigi/news-ajf-sa/providers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config := fiber.Config{
		CaseSensitive:            true,
		StrictRouting:            true,
		DisableHeaderNormalizing: true,
		ServerHeader:             "go",
	}
	for i := range os.Args[1:] {
		if os.Args[1:][i] == "-prefork" {
			config.Prefork = true
		}
	}
	app := fiber.New(config)
	dbConn := providers.Connect()
	entiry := db.NewEntiry(dbConn)
	hd := handlers.NewHandler(entiry)
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	setupRouter(app, hd)
	// go main()
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func setupRouter(app *fiber.App, hd *handlers.Handler) {
	app.Get("/", hd.HomePage)
	app.Get("/:id", hd.GetOnePost)
}
