package handlers

import (
	"log"

	"github.com/alfuhigi/news-ajf-sa/db"
	"github.com/gofiber/fiber/v2"
)

// DashBoard struct
type DashBoard struct {
	*db.Entiry
}

// NewDashBoard create DashBoard
func NewDashBoard(entiry *db.Entiry) *DashBoard {
	return &DashBoard{
		Entiry: entiry,
	}
}

func (d *DashBoard) LoginForm(ctx *fiber.Ctx) error {

	return ctx.Render("login", fiber.Map{}, "layout")
}

func (d *DashBoard) TryLogin(ctx *fiber.Ctx) error {
	type Request struct {
		User     string `form:"username"`
		Password string `form:"password"`
	}
	var body Request
	err := ctx.BodyParser(&body)
	if err != nil {
		log.Println(err)
	}
	if body.User == "north" {
		if body.Password == "ajall2ziz" {
			return ctx.Render("welcome", fiber.Map{"user": body}, "layout")
		}
	}

	return ctx.Render("login", fiber.Map{"error": "اسم المستخدم او كلمة المرور خاطئة"}, "layout")

}
