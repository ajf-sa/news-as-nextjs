package handlers

import (
	"github.com/alfuhigi/news-ajf-sa/db"
	"github.com/gofiber/fiber/v2"
)

// Handler struct
type Handler struct {
	*db.Entiry
}

// NewHandler create Handler
func NewHandler(entiry *db.Entiry) *Handler {
	return &Handler{
		Entiry: entiry,
	}
}

// HomePage main page for news.ajf.sa
func (h *Handler) HomePage(ctx *fiber.Ctx) error {
	return ctx.Send([]byte("Hi"))
}

// GetOnePost method
func (h *Handler) GetOnePost(ctx *fiber.Ctx) error {

	return ctx.Send([]byte(ctx.Params("id")))
}
