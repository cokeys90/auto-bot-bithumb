package router

import "github.com/gofiber/fiber/v2"

const CommonPrefix = "/common"

type CommonRouter struct {
	router fiber.Router
}

func NewCommonRouter(_router fiber.Router) *CommonRouter {
	r := &CommonRouter{
		router: _router.Group(CommonPrefix),
	}

	r.initHandler()

	return r
}

func (r *CommonRouter) initHandler() {
	r.router.Get("/ping", r.ping) // /api/common/ping
}

func (r *CommonRouter) ping(_ctx *fiber.Ctx) error {
	return _ctx.JSON("pong")
}
