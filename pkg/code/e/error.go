package e

import (
	"github.com/gofiber/fiber/v2"
	c "home/pkg/code"
)

func NewError(code int, msg ...string) *fiber.Error {
	var val string
	if len(msg) > 0 && msg[0] != "" {
		val = msg[0]
	} else {
		val = c.Value(code)
	}

	return fiber.NewError(code, val)
}
