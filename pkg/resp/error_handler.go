package resp

import (
	"home/pkg/code/e"

	"github.com/gofiber/fiber/v2"
)

//  ErrorHandler unified processing error
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	fe, ok := err.(*fiber.Error)
	if !ok {
		fe = e.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return New(ctx).JSON(fe)
}
