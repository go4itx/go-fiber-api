package validate

import (
	"github.com/gofiber/fiber/v2"
)

// StructParser Parse and validate struct
func StructParser(ctx *fiber.Ctx, data interface{}, query ...bool) (err error) {
	if (len(query) > 0 && query[0]) || ctx.Request().Header.IsGet() {
		err = ctx.QueryParser(data)
	} else {
		err = ctx.BodyParser(data)
	}

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = Struct(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return
}
