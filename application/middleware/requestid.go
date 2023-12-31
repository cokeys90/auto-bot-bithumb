package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// RegisterRequestIdHandler RequestId 미들웨어 등록
func RegisterRequestIdHandler(_app *fiber.App) {
	_app.Use(requestid.New(requestid.ConfigDefault))
}
