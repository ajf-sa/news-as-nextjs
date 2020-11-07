package handlers

import (
	"fmt"

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
	return ctx.SendString(fmt.Sprintf("list post user %d", userId))
}

func (d *DashBoard) Setting(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userid")
	return ctx.SendString(fmt.Sprintf("list setting user %d", userId))
}
