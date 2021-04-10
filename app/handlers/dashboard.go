package handlers

import (
	"log"
	"strconv"

	"github.com/alfuhigi/news-ajf-sa/db"
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
	userId := ctx.Locals("userid")
	var limit int
	limit = 2
	pagnation := ctx.Query("page", "1")
	pg, err := strconv.ParseInt(pagnation, 10, 32)
	if pg == 0 {
		pg = 1
	}
	offset := limit * (int(pg) - 1)
	users, err := d.User.FindMany().Take(limit).Skip(offset).Exec(ctx.Context())
	if err != nil {
		log.Println(err)
		return ctx.Redirect("/cp/users")
	}

	next := true
	nextNum := pg + 1
	lastPage, err := d.User.FindMany().Take(limit).Skip(offset + limit).Exec(ctx.Context())
	if err != nil {
		log.Println(err)

	}
	if len(lastPage) == 0 {

		next = false
	}

	prev := true
	prevNum := pg - 1
	firstPage, err := d.User.FindMany().Take(limit).Skip(offset - limit).Exec(ctx.Context())
	if len(firstPage) == 0 || len(users) == 0 {

		prev = false
	}

	return ctx.Render("users", fiber.Map{"userId": userId, "users": users, "page": pg, "prev": prev, "prevNum": prevNum, "next": next, "nextNum": nextNum}, "layout")

}

func (d *DashBoard) Dashboard(ctx *fiber.Ctx) error {
	// userId := ctx.Locals("userid")
	ok := ctx.SendFile("./webapp/index.html")
	if ok != nil {
		return ctx.SendString("Ok!")
	}
	return ok
	// return ctx.Render("dashboard", fiber.Map{"userId": userId}, "layout")
}
