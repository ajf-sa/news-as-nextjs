package handlers

import (
	"log"
	"strconv"

	"github.com/alfuhigi/news-ajf-sa/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

type API struct {
	*db.PrismaClient
	*session.Session
}

func NewAPI(entiry *db.PrismaClient, session *session.Session) *API {
	return &API{
		PrismaClient: entiry,
		Session:      session,
	}
}

///####### App ########//
func (a *API) SetApp(ctx *fiber.Ctx) error {
	return nil
}

///####### User ########//

func (a *API) GetListUser(ctx *fiber.Ctx) error {
	return nil
}
func (a *API) GetOneUser(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Params("id", ""))
	user, err := a.User.FindUnique(db.User.ID.Equals(userId)).Exec(ctx.Context())
	if err != nil {
		log.Println(err)
	}
	return ctx.JSON(user)

}
func (a *API) SetOneUser(ctx *fiber.Ctx) error {
	type User struct {
		UserName string `json:"username" xml:"username" form:"username"`
		Password string `json:"password" xml:"password" form:"password"`
		Phone    string `json:"phone" xml:"phone" form:"phone"`
	}
	u := new(User)
	if err := ctx.BodyParser(u); err != nil {
		return err
	}
	log.Println(u.UserName, u.Password, u.Phone)
	return ctx.JSON(fiber.Map{"login": true, "token": 111111111111})
}

func (a *API) SetActiveOneUser(ctx *fiber.Ctx) error {
	return nil
}

///####### Post ########//

func (a *API) GetonePost(ctx *fiber.Ctx) error {
	return nil
}

func (a *API) GetListPost(ctx *fiber.Ctx) error {
	return nil
}

func (a *API) SetOnePost(ctx *fiber.Ctx) error {
	return nil
}

///####### Tag ########//

func (a *API) SetOneTag(ctx *fiber.Ctx) error {
	return nil
}
func (a *API) GetListTag(ctx *fiber.Ctx) error {
	return nil
}

func (a *API) GetOneTag(ctx *fiber.Ctx) error {
	return nil
}

///####### Image ########//

func (a *API) GetListImage(ctx *fiber.Ctx) error {
	return nil
}

func (a *API) GetOneImage(ctx *fiber.Ctx) error {
	return nil
}

func (a *API) SetOneImage(ctx *fiber.Ctx) error {
	return nil
}

func (a API) DeleteOneImage(ctx *fiber.Ctx) error {
	return nil
}
