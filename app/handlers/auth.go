package handlers

import (
	"fmt"
	"log"

	"github.com/alfuhigi/news-ajf-sa/db"
	"github.com/alfuhigi/news-ajf-sa/providers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

type Auth struct {
	*db.PrismaClient
	*session.Session
}

type User struct {
	Id       int
	Username string
	Password string
	Phone    string
}

func NewAuth(entiry *db.PrismaClient, session *session.Session) *Auth {
	return &Auth{
		PrismaClient: entiry,
		Session:      session,
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
	isUser, err := a.User.FindFirst(
		db.User.Username.Equals(body.User),
	).Exec(ctx.Context())

	if err != nil {
		return ctx.Render("login", fiber.Map{"error": "اسم السمتخدم غير موجود"})
	}
	fmt.Println(isUser.Username)
	isMatch := providers.CheckPasswordHash(body.Password, isUser.Password)

	if isMatch {
		if isUser.IsActive {
			a.Login(ctx, isUser.ID)
			return ctx.Redirect(next)
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

func (a *Auth) APIRegister(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{"test": "Hi I'm from backend"})

}

func (a *Auth) PostRegister(ctx *fiber.Ctx) error {

	type Request struct {
		User     string `form:"username"`
		Password string `form:"password"`
		Phone    string `form:"phone"`
	}
	var body Request
	err := ctx.BodyParser(&body)

	if err != nil {
		log.Println(err)
	}
	isUser, _ := a.User.FindFirst(
		db.User.Username.Equals(body.User),
	).Exec(ctx.Context())

	if isUser != nil {
		return ctx.Render("register", fiber.Map{"error": "اسم المستخدم موجود مسبقا"})
	}

	password, _ := providers.HashPassword(body.Password)

	newUser, _ := a.User.CreateOne(
		db.User.Username.Set(body.User),
		db.User.Password.Set(password),
		db.User.Phone.Set(body.Phone),
	).Exec(ctx.Context())
	fmt.Println(newUser.Username)
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

func (a *Auth) Login(c *fiber.Ctx, id int) {
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
