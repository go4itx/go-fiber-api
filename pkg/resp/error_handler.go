package resp

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler unified processing error
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// log.Printf("ErrorHandler: %v \n", err)
	// log.Println(string(debug.Stack()))
	return New(ctx).JSON(err)
}
