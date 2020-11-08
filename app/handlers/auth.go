package handlers

import (
	"log"

	"github.com/alfuhigi/news-ajf-sa/db"
	"github.com/alfuhigi/news-ajf-sa/providers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

type Auth struct {
	*db.Entiry
	*session.Session
}

type User struct {
	Id       uint
	Username string
	Password string
}

var users = []*User{
	{Id: 1, Username: "north", Password: "ajall2ziz"},
	{Id: 2, Username: "aziz", Password: "ajall2ziz"},
	{Id: 3, Username: "bader", Password: "ajall2ziz"},
}

func NewAuth(entiry *db.Entiry, session *session.Session) *Auth {
	return &Auth{
		Entiry:  entiry,
		Session: session,
	}
}

func (a *Auth) LoginForm(ctx *fiber.Ctx) error {
	// userid, err := providers.ParseToken(ctx, "thisissecretkey")
	// if err != nil {
	// 	log.Println(err)
	// }

	// if userid > 0 {
	// 	return ctx.Redirect("/cp")

	// }
	store := a.Get(ctx)
	defer store.Save()
	userid := store.Get("user_id")
	if userid != nil {
		log.Println("Sessions: ", userid)
		return ctx.Redirect("/cp")
	}
	log.Println("Sessions not found: ", userid)
	next := ctx.Query("next")
	if next == "" || next == "/auth/login" {
		next = "/cp"
	}
	return ctx.Render("login", fiber.Map{"next": next})
}

func (a *Auth) TryLogin(ctx *fiber.Ctx) error {
	type Request struct {
		User     string `form:"username"`
		Password string `form:"password"`
		Next     string `form:"next"`
	}
	var body Request
	err := ctx.BodyParser(&body)

	if err != nil {
		log.Println(err)
	}

	next := body.Next
	store := a.Get(ctx)
	defer store.Save()
	for _, user := range users {
		if user.Username == body.User {
			if user.Password == body.Password {

				store.Set("user_id", user.Id)
				// _, err := providers.CreateToken(ctx, user.Id, "thisissecretkey")
				// if err != nil {
				// 	log.Printf("Token Error ", err)
				// }
				return ctx.Redirect(next)

			}
		}
	}
	return ctx.Render("login", fiber.Map{"error": "اسم المستخدم او كلمة المرور خاطئة", "next": next})

}

func (a *Auth) Logout(ctx *fiber.Ctx) error {
	store := *a.Get(ctx) // get/create new session
	store.Destroy()
	store.Delete("user_id")
	defer store.Save()
	providers.DeleteToken(ctx)
	log.Println("Sessions logut", store.Get("user_id"))
	log.Println("logout")

	return ctx.Redirect("/auth/login")
}
