package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/session/v2"
)

// Global session storage
// var sessions = session.New()

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	sessions := session.New()

	app.Get("/login", func(ctx *fiber.Ctx) error {
		store := sessions.Get(ctx)
		defer store.Save()
		store.Set("user_id", 1)
		// cookie := new(fiber.Cookie)
		// cookie.Name = "test-fiber"
		// cookie.Value = "test-test-test"
		// cookie.Expires = time.Now().Add(24 * time.Hour)
		// ctx.Cookie(cookie)
		user_id := store.Get("user_id")
		if user_id != nil {
			user := strconv.Itoa(user_id.(int))
			return ctx.SendString(user)
		}
		return ctx.SendString("nil")
	})
	app.Get("/private", func(ctx *fiber.Ctx) error {
		store := sessions.Get(ctx)
		// defer store.Save()
		// cookie := ctx.Cookies("test-fiber")
		user_id := store.Get("user_id")
		if user_id != nil {
			user := strconv.FormatInt(user_id.(int64), 10)
			return ctx.SendString(user)
		}
		return ctx.SendString("nil")

	})
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hi")
	})

	app.Listen(":3000")

}
