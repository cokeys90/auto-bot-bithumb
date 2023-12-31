package middleware

import (
	"database/sql"
	"errors"
	"github.com/cokeys90/auto-bot-bithumb/application/consts"
	"github.com/gofiber/fiber/v2"
)

func RegisterGlobalErrorHandler(_router fiber.Router) {
	handler := func() fiber.Handler {
		return func(c *fiber.Ctx) error {
			err := c.Next()
			if err != nil {
				return parseError(err)
			}
			return err
		}
	}()
	_router.Use(handler)
}

// parseError 에러 파싱하여 적합한 상태를 내려준다.
func parseError(err error) error {
	if errors.Is(err, consts.ErrParameterMissing) {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if errors.Is(err, consts.ErrParameterInvalid) {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if errors.Is(err, consts.ErrNotFound) {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	} else if errors.Is(err, sql.ErrNoRows) {
		return fiber.NewError(fiber.StatusNotFound, consts.ErrNotFound.Error())
	} else if errors.Is(err, consts.ErrUnauthorized) {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	} else if errors.Is(err, consts.ErrForbidden) {
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	} else {
		return err
	}
}
