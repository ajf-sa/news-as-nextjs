package handlers

import (
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

func (d *DashBoard) GetListPost(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userid")
	return ctx.Render("posts", fiber.Map{"userId": userId}, "layout")

}

func (d *DashBoard) Setting(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userid")
	return ctx.Render("setting", fiber.Map{"userId": userId}, "layout")
}
