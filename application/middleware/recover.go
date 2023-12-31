package middleware

import (
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
)

// RegisterCrashRecoverHandler Exception 으로 인한 서버 다운 방지
func RegisterCrashRecoverHandler(_app *fiber.App) {
	_app.Use(recover2.New())
}
