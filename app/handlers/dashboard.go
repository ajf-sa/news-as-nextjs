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

func (d *DashBoard) Login(ctx *fiber.Ctx) error {
	return ctx.SendString("DashBoard Login")
}
