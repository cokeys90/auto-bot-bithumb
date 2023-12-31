package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/cokeys90/auto-bot-bithumb/application/log"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/url"
)

type Logger struct {
	logger *logrus.Logger
}

func RegisterLoggerHandler(_app *fiber.App) {
	l := &Logger{
		logger: log.Instance(),
	}
	_app.Use(l.handler)
}

func (l Logger) handler(ctx *fiber.Ctx) error {
	reqId := ctx.Locals("requestid")

	var data string
	method := ctx.Method()
	switch method {
	case fiber.MethodGet:
		fallthrough
	case fiber.MethodHead:
		fallthrough
	case fiber.MethodOptions:
		q, err := url.QueryUnescape(string(ctx.Request().URI().QueryString()))
		if err != nil {
			q = string(ctx.Request().URI().QueryString())
		}
		data = q
	default:
		dst := &bytes.Buffer{}
		body := ctx.Body()
		if body == nil || len(body) == 0 {
			break
		}

		switch ctx.Get(fiber.HeaderContentType, "") {
		case fiber.MIMEApplicationJSON:
			if err := json.Compact(dst, ctx.Body()); err == nil {
				data = dst.String()
			}
		default:
			data = string(body)
		}
	}

	l.logger.Infof("%s [%s] [%s, %s] %s %s", ctx.IP(), reqId, ctx.Method(), ctx.Get(fiber.HeaderContentType, "none"), ctx.Path(), data)
	return ctx.Next()
}
