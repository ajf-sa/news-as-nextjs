package handlers

import (
	db "github.com/alfuhigi/news-ajf-sa/db"
	"github.com/gofiber/fiber/v2"
)

// DashBoard struct
type DashBoard struct {
	*db.PrismaClient
}

// NewDashBoard create DashBoard
func NewDashBoard(entiry *db.PrismaClient) *DashBoard {
	return &DashBoard{
		PrismaClient: entiry,
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

func (d *DashBoard) Users(ctx *fiber.Ctx) error {
	// userId := ctx.Locals("userid")
	// var limit int32
	// limit = 2
	// pagnation := ctx.Query("page", "1")
	// log.Println(pagnation)
	// pg, err := strconv.ParseInt(pagnation, 10, 32)
	// if pg == 0 {
	// 	pg = 1
	// }
	// offset := limit * (int32(pg) - 1)
	// users, err := d.ListUsers(ctx.Context(), db.ListUsersParams{Limit: limit, Offset: int32(offset)})
	// if err != nil {
	// 	log.Println(err)
	// 	return ctx.Redirect("/cp/users")
	// }

	// next := true
	// nextNum := pg + 1
	// lastPage, err := d.ListUsers(ctx.Context(), db.ListUsersParams{Limit: limit, Offset: int32(offset) + 1})
	// if err != nil {
	// 	log.Println(err)

	// }
	// if len(lastPage) == 0 {

	// 	next = false
	// }

	// prev := true
	// prevNum := pg - 1
	// fisrtPage, err := d.ListUsers(ctx.Context(), db.ListUsersParams{Limit: limit, Offset: int32(offset) - 1})
	// if err != nil {
	// 	log.Println(err)

	// }
	// if len(fisrtPage) == 0 {

	// 	prev = false
	// }

	//return ctx.Render("users", fiber.Map{"userId": userId, "users": users, "page": pg, "prev": prev, "prevNum": prevNum, "next": next, "nextNum": nextNum}, "layout")
	return ctx.SendString("TODO")
}

func (d *DashBoard) Dashboard(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userid")
	return ctx.Render("dashboard", fiber.Map{"userId": userId}, "layout")
}
