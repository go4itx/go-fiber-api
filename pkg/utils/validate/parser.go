package validate

import (
	"github.com/gofiber/fiber/v2"
)

// StructParser Parse and validate struct
func StructParser(ctx *fiber.Ctx, data interface{}) (err error) {
	if string(ctx.Request().Header.Method()) == "GET" {
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

// MapParser Parse and validate map
func MapParser(ctx *fiber.Ctx, rules map[string]string) (ret map[string]interface{}, err error) {
	var data map[string]interface{}
	if string(ctx.Request().Header.Method()) == "GET" {
		err = ctx.QueryParser(&data)
	} else {
		err = ctx.BodyParser(&data)
	}

	if err != nil {
		err = fiber.NewError(fiber.StatusBadRequest, err.Error())
		return
	}

	ret = make(map[string]interface{})
	for k, rule := range rules {
		v, ok := data[k]
		if !ok {
			continue
		}

		if rule != "" {
			if err = Variable(k, v, rule); err != nil {
				return
			}
		}

		ret[k] = v
	}

	return
}
