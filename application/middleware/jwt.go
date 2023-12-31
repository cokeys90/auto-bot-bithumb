package middleware

import (
	"errors"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const jwtContextKey = "user"

func RegisterJsonWebTokenHandler(_app *fiber.App, _password string) {
	handler := func(ctx *fiber.Ctx, err error) error {
		errMsg := err.Error()
		if errors.Is(err, jwt.ErrTokenRequiredClaimMissing) {
			return fiber.NewError(fiber.StatusUnauthorized, errMsg)
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return fiber.NewError(fiber.StatusUnauthorized, jwt.ErrTokenExpired.Error())
		}
		return fiber.NewError(fiber.StatusUnauthorized, errMsg)
	}

	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte(_password),
		},
		SuccessHandler: nil,
		ErrorHandler:   handler,
		ContextKey:     jwtContextKey,
		Filter: func(ctx *fiber.Ctx) bool {
			return true // true: JWT 인증 안함(통과)
		},
	}

	_app.Use(jwtware.New(config))
}
