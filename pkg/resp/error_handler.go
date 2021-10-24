package resp

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

//  ErrorHandler unified processing error
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	ret := result{
		Code:       0,
		Msg:        "",
		ServerTime: time.Now().Unix(),
		Data:       "",
	}

	if e, ok := err.(*fiber.Error); ok {
		ret.Code = e.Code
		ret.Msg = e.Message
	} else {
		ret.Code = fiber.StatusInternalServerError
		ret.Msg = err.Error()
	}

	return ctx.Status(fiber.StatusOK).JSON(ret)
}
