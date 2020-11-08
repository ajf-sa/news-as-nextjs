package providers

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Token struct {
	Hash   string
	Expire int64
}
type AuthConfiguration struct {
	App_Jwt_Secret string
	Api_Jwt_Secret string
	Jwt_Expire     int
}

var AuthConfig *AuthConfiguration //nolint:gochecknoglobals

func CreateToken(ctx *fiber.Ctx, userID uint, secret string) (Token, error) {
	var t Token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	expiresIn := time.Now().Add(time.Duration(60*60) * time.Second).Unix()
	claims["exp"] = expiresIn

	tokenHash, err := token.SignedString([]byte(secret))
	if err != nil {
		return t, err
	}
	ctx.Cookie(&fiber.Cookie{
		Name:   "userid",
		Value:  tokenHash,
		Secure: false,
	})
	t.Hash = tokenHash
	t.Expire = expiresIn
	return t, nil
}

func ParseToken(ctx *fiber.Ctx, secret string) (uint, error) {
	tokenString := ctx.Cookies("userid")
	fmt.Println("Cookie:", tokenString)
	if tokenString == "" {
		return 0, errors.New("Empty auth cookie")

	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}
	err2 := claims.Valid()

	if err2 != nil {
		DeleteToken(ctx)
		return 0, err2
	}

	return uint(claims["id"].(float64)), nil
}

//DeleteToken deletes the jwt token
func DeleteToken(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name: "userid",
		// Set expiry date to the past
		Expires:  time.Now().Add(-(time.Hour * 2)),
		HTTPOnly: false,
		SameSite: "lax",
	})
}

//RefreshToken refreshes the token
func RefreshToken(c *fiber.Ctx, secret string) (Token, error) {
	var t Token
	u, err := ParseToken(c, secret)

	if err != nil {
		return t, err
	}

	return CreateToken(c, u, secret)
}
