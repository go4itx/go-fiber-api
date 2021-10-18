package controller

import (
	"github.com/gofiber/fiber/v2"
	"home/pkg/server"
)

var prefix = "/v1"

// auth required user authentication
func auth(router fiber.Router) {
	userRouter(router)
}

// noAuth  no user authentication required
func noAuth(app *fiber.App) {
	app.Post(prefix+"/login", login)
}

// Init controller
func Init() (err error) {
	return server.Init(prefix, noAuth, auth)
}
