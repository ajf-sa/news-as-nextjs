package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

// Global session storage
var sessions = session.New()

func main() {
	app := fiber.New()
	app.Get("/login", func(ctx *fiber.Ctx) error {
		store := sessions.Get(ctx)
		defer store.Save()
		store.Set("user_id", 1)
		// cookie := new(fiber.Cookie)
		// cookie.Name = "test-fiber"
		// cookie.Value = "test-test-test"
		// cookie.Expires = time.Now().Add(24 * time.Hour)
		// ctx.Cookie(cookie)
		user_id := store.Get("user_id").(int)
		return ctx.SendString(string(user_id))
	})
	app.Get("/private", func(ctx *fiber.Ctx) error {
		store := sessions.Get(ctx)
		defer store.Save()
		// cookie := ctx.Cookies("test-fiber")
		user_id := store.Get("user_id").(int)
		return ctx.SendString(string(user_id))

	})
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hi")
	})

	app.Listen(":3000")

}
