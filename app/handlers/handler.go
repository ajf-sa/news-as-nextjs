package handlers

import "github.com/gofiber/fiber/v2"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HomePage(ctx *fiber.Ctx) error {
	return ctx.Send([]byte(""))
}

func (h *Handler) Post(ctx *fiber.Ctx) error {

	return ctx.Send([]byte(ctx.Params("id")))
}
