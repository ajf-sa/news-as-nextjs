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
	"github.com/gofiber/template/html"
)

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
	auth := handlers.NewAuth(entiry)
	acn := app.Group("auth")
	acn.Get("/login", auth.LoginForm)
	acn.Post("/login", auth.TryLogin)
}

func setupDashboard(app *fiber.App, entiry *db.Entiry) {
	cp := handlers.NewDashBoard(entiry)
	dsh := app.Group("cp", func(ctx *fiber.Ctx) error {
		next := string(ctx.Request().RequestURI())
		userID, err := providers.ParseToken(ctx, "thisissecretkey")
		if err != nil {
			return ctx.Redirect(fmt.Sprintf("/auth/login?next=%s", next))

		}
		log.Println("this is protected")
		ctx.Locals("userid", userID)
		return ctx.Next()
	})

	dsh.Get("/posts", cp.GetListPost)
	dsh.Get("/setting", cp.Setting)
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
