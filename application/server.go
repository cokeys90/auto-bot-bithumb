package application

import (
	"fmt"
	"github.com/cokeys90/auto-bot-bithumb/application/config"
	logger "github.com/cokeys90/auto-bot-bithumb/application/log"
	"github.com/cokeys90/auto-bot-bithumb/application/middleware"
	"github.com/cokeys90/auto-bot-bithumb/application/router"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"log"
)

type Server struct {
	app *fiber.App
}

func NewApplication() *Server {
	return &Server{
		app: fiber.New(),
	}
}

func (s *Server) Listen() {
	cfg := config.GetConfig()
	fmt.Print(cfg)

	_initLogger(cfg)
	_initMiddleware(s.app, cfg)
	_initRouter(s.app)

	port := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Panic(s.app.Listen(port))
}

func _initLogger(_cfg config.ApplicationInfo) {
	logConfig := logger.Config{
		ZonedId:  "Asia/Seoul",
		Level:    logrus.InfoLevel,
		FilePath: "",
	}
	logger.Setup(logConfig)
}

// _initMiddleware 미들웨어 초기화
func _initMiddleware(_app *fiber.App, _cfg config.ApplicationInfo) {
	middleware.RegisterCrashRecoverHandler(_app)
	middleware.RegisterCorsHandler(_app)
	middleware.RegisterRequestIdHandler(_app)
	middleware.RegisterGlobalErrorHandler(_app)
	middleware.RegisterJsonWebTokenHandler(_app, _cfg.Server.JwtSecret)
}

// _initRouter 라우터 초기화
func _initRouter(_app *fiber.App) {
	router.RegisterBaseRouter(_app)
}
