package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/login", func(ctx *fiber.Ctx) error {
		cookie := new(fiber.Cookie)
		cookie.Name = "test-fiber"
		cookie.Value = "test-test-test"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		ctx.Cookie(cookie)
		return ctx.SendString(cookie.Value)
	})
	app.Get("/private", func(ctx *fiber.Ctx) error {
		cookie := ctx.Cookies("test-fiber")
		return ctx.SendString(cookie)
	})
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hi")
	})

	app.Listen(":3000")

}
