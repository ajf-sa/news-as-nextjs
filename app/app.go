package main

import (
	"fmt"
	"os"

	"github.com/alfuhigi/news-ajf-sa/db"
	"github.com/alfuhigi/news-ajf-sa/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/session/v2"
	"github.com/gofiber/template/html"
)

// Global session storage
var sessions *session.Session

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {

		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {

			panic(err)
		}

	}()

	engine := html.New("./views", ".html")
	sessions = session.New()
	config := fiber.Config{
		ServerHeader: "go",
		Views:        engine,
	}
	for i := range os.Args[1:] {
		if os.Args[1:][i] == "-prefork" {
			config.Prefork = true
		}
	}
	app := fiber.New(config)

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	setupAuth(app, client)
	setupDashboard(app, client)
	setupRouter(app, client)
	app.Get("/robots.txt", func(ctx *fiber.Ctx) error {
		return ctx.SendString(`
		User-agent: *
		Allow: /
		`)
	})
	app.Static("/cp", "webapp")

	// app.Static("/", "webapp")
	// app.Get("/*", func(ctx *fiber.Ctx) error {

	// 	ok := ctx.SendFile("./webapp/index.html")
	// 	if ok != nil {
	// 		return ctx.SendString("Ok!")
	// 	}
	// 	return ctx.SendString("OK!")

	// })

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func setupAuth(app *fiber.App, entiry *db.PrismaClient) {
	auth := handlers.NewAuth(entiry, sessions)
	acn := app.Group("auth")
	acn.Post("/login", auth.PostLogin)
	acn.Get("/login", auth.LoginForm)

	acn.Post("/register", auth.PostRegister)
	acn.Get("/register", auth.RegisterForm)

	acn.Get("/logout", auth.Logout)

}

func setupDashboard(app *fiber.App, entiry *db.PrismaClient) {
	cp := handlers.NewDashBoard(entiry)
	dsh := app.Group("cp", func(ctx *fiber.Ctx) error {
		next := string(ctx.Request().RequestURI())

		store := sessions.Get(ctx)
		userid := store.Get("user_id")
		if userid != nil {
			ctx.Locals("userid", userid)
			return ctx.Next()
		}

		return ctx.Redirect(fmt.Sprintf("/auth/login?next=%s", next))

	})

	dsh.Get("/posts", cp.GetListPost)
	dsh.Get("/setting", cp.Setting)
	dsh.Get("/users", cp.Users)
	dsh.Get("/*", cp.Dashboard)
}
func setupRouter(app *fiber.App, entiry *db.PrismaClient) {
	us := handlers.NewAuth(entiry, sessions)
	hd := handlers.NewHandler(entiry)
	grp := app.Group("api")
	grp.Get("/user/:id", us.GetOneUser)
	grp.Get("/users", us.APIRegister)
	grp.Get("/about", hd.AboutPage)
	grp.Get("/tech", hd.TechPage)
	grp.Get("/sport", hd.SportPage)
	grp.Get("/local", hd.LocalPage)
	grp.Get("/", hd.HomePage)

}
