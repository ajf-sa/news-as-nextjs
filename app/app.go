package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alfuhigi/news-ajf-sa/db"
	"github.com/alfuhigi/news-ajf-sa/handlers"
	"github.com/alfuhigi/news-ajf-sa/providers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/session/v2"
	"github.com/gofiber/template/html"
)

// Global session storage
var sessions = session.New()

func main() {
	engine := html.New("./views", ".html")
	config := fiber.Config{
		CaseSensitive:            true,
		StrictRouting:            true,
		DisableHeaderNormalizing: true,
		ServerHeader:             "go",
		Views:                    engine,
	}
	for i := range os.Args[1:] {
		if os.Args[1:][i] == "-prefork" {
			config.Prefork = true
		}
	}
	app := fiber.New(config)
	dbConn := providers.Connect()
	entiry := db.NewEntiry(dbConn)

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	setupAuth(app, entiry)
	setupDashboard(app, entiry)
	setupRouter(app, entiry)
	app.Get("/robots.txt", func(ctx *fiber.Ctx) error {
		return ctx.SendString(`
		User-agent: *
		Allow: /
		`)
	})
	app.Static("/", "webapp")
	app.Get("/*", func(ctx *fiber.Ctx) error {

		ok := ctx.SendFile("./webapp/index.html")
		if ok != nil {
			return ctx.SendString("HI")
		}
		return ok

	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func setupAuth(app *fiber.App, entiry *db.Entiry) {
	auth := handlers.NewAuth(entiry, sessions)
	acn := app.Group("auth")
	acn.Get("/login", auth.LoginForm)
	acn.Post("/login", auth.TryLogin)

	acn.Get("/logout", auth.Logout)

}

func setupDashboard(app *fiber.App, entiry *db.Entiry) {
	cp := handlers.NewDashBoard(entiry)
	dsh := app.Group("cp", func(ctx *fiber.Ctx) error {
		next := string(ctx.Request().RequestURI())
		store := sessions.Get(ctx)
		userid := store.Get("user_id")
		if userid != nil {

			log.Println("this is protected", userid)
			ctx.Locals("userid", userid)
			return ctx.Next()
		}
		return ctx.Redirect(fmt.Sprintf("/auth/login?next=%s", next))

	})

	dsh.Get("/posts", cp.GetListPost)
	dsh.Get("/setting", cp.Setting)
	dsh.Get("/users", cp.Users)
	dsh.Get("/", cp.Dashboard)
}
func setupRouter(app *fiber.App, entiry *db.Entiry) {
	hd := handlers.NewHandler(entiry)
	grp := app.Group("api")
	grp.Get("/about", hd.AboutPage)
	grp.Get("/tech", hd.TechPage)
	grp.Get("/sport", hd.SportPage)
	grp.Get("/local", hd.LocalPage)
	grp.Get("/", hd.HomePage)
	// app.Get("/:id", hd.GetOnePost)
}
