package controller

import (
	"home/pkg/server"

	"github.com/gofiber/fiber/v2"
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
	return server.Init(prefix, nil, noAuth, auth)
}
