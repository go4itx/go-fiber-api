package controller

import (
	_ "home/docs/demo"
	"home/pkg/component/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

var prefix = "/v1"

// auth required user authentication
func auth(router fiber.Router) {
	{
		r := router.Group("/user")
		user := &user{}
		r.Get("/", user.info)
	}
}

// noAuth  no user authentication required
func noAuth(app *fiber.App) {
	app.Post(prefix+"/login", login)

	app.Get("/swagger/*", swagger.New())
}

// @title DEMO API
// @version 1.0
// @description  This is api document
// @host localhost:20105
// @BasePath /v1
func Init() (err error) {
	return server.Init(prefix, nil, noAuth, auth)
}
