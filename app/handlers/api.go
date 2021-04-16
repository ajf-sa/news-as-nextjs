package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/alfuhigi/news-ajf-sa/db"
	"github.com/alfuhigi/news-ajf-sa/providers"
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

func (a *API) Protected(ctx *fiber.Ctx) error {
	fmt.Println("1st route!")
	userID, _ := providers.ParseToken(ctx, "secret")
	ctx.Locals("userId", strconv.Itoa(int(userID)))
	return ctx.Next()
}

///####### App ########//
func (a *API) SetApp(ctx *fiber.Ctx) error {
	return nil
}

///####### User ########//

func (a *API) GetLoginByToken(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Locals("userId").(string))
	log.Println(userId)
	user, err := a.User.FindUnique(db.User.ID.Equals(userId)).Exec(ctx.Context())
	if err != nil {
		return ctx.JSON(fiber.Map{"login": false, "error": "nologin"})
	}
	return ctx.JSON(fiber.Map{"login": true, "username": user.Username})
}
func (a *API) GetListUser(ctx *fiber.Ctx) error {
	return nil

}
func (a *API) GetLoginUser(ctx *fiber.Ctx) error {

	type User struct {
		UserName string `json:"username" xml:"username" form:"username"`
		Password string `json:"password" xml:"password" form:"password"`
	}

	u := new(User)
	if err := ctx.BodyParser(u); err != nil {
		return err
	}

	isUser, _ := a.User.FindFirst(
		db.User.Username.Equals(u.UserName),
	).Exec(ctx.Context())

	isMatch := providers.CheckPasswordHash(u.Password, isUser.Password)

	if isMatch {
		if isUser.IsActive {
			token, _ := providers.CreateToken(ctx, uint(isUser.ID), "secret")
			return ctx.JSON(fiber.Map{"login": true, "token": token.Hash})
		}
		return ctx.JSON(fiber.Map{"login": false, "error": "user is not active"})

	}
	return ctx.JSON(fiber.Map{"login": false, "error": "username error or password error"})
}

func (a *API) GetOneUser(ctx *fiber.Ctx) error {

	userId, _ := strconv.Atoi(ctx.Params("id", ""))
	user, err := a.User.FindUnique(db.User.ID.Equals(userId)).Exec(ctx.Context())
	if err != nil {
		return ctx.JSON(fiber.Map{"error": err})
	}
	return ctx.JSON(fiber.Map{"user": user})

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
	isUser, _ := a.User.FindFirst(
		db.User.Username.Equals(u.UserName),
	).Exec(ctx.Context())

	if isUser != nil {
		return ctx.JSON(fiber.Map{"error": "user is exisit"})
	}

	password, _ := providers.HashPassword(u.Password)

	newUser, _ := a.User.CreateOne(
		db.User.Username.Set(u.UserName),
		db.User.Password.Set(password),
		db.User.Phone.Set(u.Phone),
	).Exec(ctx.Context())

	return ctx.JSON(fiber.Map{"registed": true, "userid": newUser.ID})
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
