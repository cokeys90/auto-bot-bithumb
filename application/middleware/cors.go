package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// RegisterCorsHandler CORS 설정
func RegisterCorsHandler(_app *fiber.App) {
	cfg := cors.ConfigDefault
	cfg.AllowCredentials = true
	_app.Use(cors.New(cfg))
}
