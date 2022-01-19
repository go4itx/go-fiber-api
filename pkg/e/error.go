package e

import (
	c "home/pkg/code"

	"github.com/gofiber/fiber/v2"
)

func NewError(code int, msg ...string) *fiber.Error {
	var val string
	if len(msg) > 0 && msg[0] != "" {
		val = msg[0]
	} else {
		val = c.Message(code)
	}

	return fiber.NewError(code, val)
}
