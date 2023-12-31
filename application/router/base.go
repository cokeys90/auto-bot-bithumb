package router

import "github.com/gofiber/fiber/v2"

const (
	basePrefix = "/api"
)

type BaseRouter struct {
	Router       fiber.Router
	CommonRouter *CommonRouter
}

var baseRouter *BaseRouter

func RegisterBaseRouter(_app *fiber.App) {
	r := &BaseRouter{
		Router: _app,
	}

	defaultRouter := r.Router.Group(basePrefix) // URL 경로: /api

	r.CommonRouter = NewCommonRouter(defaultRouter) // URL 경로: /api/common

	baseRouter = r
}
