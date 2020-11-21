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
	Id       int
	Username string
	Password string
	Email    string
}

func NewAuth(entiry *db.Entiry, session *session.Session) *Auth {
	return &Auth{
		Entiry:  entiry,
		Session: session,
	}
}

func (a *Auth) LoginForm(ctx *fiber.Ctx) error {

	if a.IsAuth(ctx) {
		return ctx.Redirect("/cp")
	}
	next := ctx.Query("next")
	if next == "" || next == "/auth/login" {
		next = "/cp"
	}
	return ctx.Render("login", fiber.Map{"next": next})
}

func (a *Auth) PostLogin(ctx *fiber.Ctx) error {

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
	isUser, err := a.GetOneUSer(ctx.Context(), body.User)
	if err != nil {
		log.Println(err)
	}
	if isUser.ID > 0 {
		isMatch := providers.CheckPasswordHash(body.Password, isUser.Password)

		if isMatch {
			if isUser.IsActive {
				a.Login(ctx, isUser.ID)
				return ctx.Redirect(next)
			}
		}
	}
	return ctx.Render("login", fiber.Map{"error": "اسم المستخدم او كلمة المرور خاطئة", "next": next})

}

func (a *Auth) RegisterForm(ctx *fiber.Ctx) error {
	if a.IsAuth(ctx) {
		return ctx.Redirect("/cp")
	}
	return ctx.Render("register", fiber.Map{})

}
func (a *Auth) PostRegister(ctx *fiber.Ctx) error {
	type Request struct {
		User     string `form:"username"`
		Password string `form:"password"`
		Email    string `form:"email"`
	}
	var body Request
	err := ctx.BodyParser(&body)

	if err != nil {
		log.Println(err)
	}
	isUser, err := a.GetOneUSer(ctx.Context(), body.User)
	if err != nil {
		log.Println(err)
	}
	if isUser.ID > 0 {
		return ctx.Render("register", fiber.Map{"error": "اسم المستخدم موجود مسبقا"})
	}
	password, _ := providers.HashPassword(body.Password)
	newUser, err := a.AddNewUser(ctx.Context(), db.AddNewUserParams{
		Username: body.User,
		Password: password,
		Email:    body.Email})

	if err != nil {
		log.Println(err)
	}
	if newUser.ID > 0 {
		return ctx.Redirect("/auth/login")
	}

	return ctx.Render("register", fiber.Map{"error": "اسم المستخدم او كلمة المرور خاطئة"})

}

func (a *Auth) Logout(ctx *fiber.Ctx) error {
	store := *a.Get(ctx)
	store.Destroy()
	store.Delete("user_id")
	defer store.Save()
	return ctx.Redirect("/auth/login")
}

func (a *Auth) Login(c *fiber.Ctx, id int32) {
	store := a.Get(c)
	defer store.Save()
	store.Set("user_id", id)
}

func (a *Auth) IsAuth(ctx *fiber.Ctx) bool {
	store := a.Get(ctx)
	userid := store.Get("user_id")
	if userid != nil {
		return true
	}
	return false
}
