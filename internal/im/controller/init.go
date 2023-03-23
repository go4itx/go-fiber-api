package controller

import (
	_ "home/docs/im"
	"home/pkg/component/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

var prefix = "/v1"

// auth required user authentication
func auth(router fiber.Router) {
	{
		r := router.Group("/admin")
		admin := &admin{}
		r.Get("/", admin.info)
	}

	{
		r := router.Group("/im")
		im := &im{}
		r.Get("/online", im.Online)
		r.Post("/kick", im.Kick)
		r.Post("/sendMessage", im.SendMessage)
	}
}

// noAuth  no user authentication required
func noAuth(app *fiber.App) {
	app.Get(prefix+"/im", ws()) //websocket
	app.Post(prefix+"/im/login", login)

	app.Get("/swagger/*", swagger.New())
}

// Init ...
// @title IM API
// @version 1.0
// @description  This is api document
// @host localhost:20105
// @BasePath /v1
func Init() (err error) {
	return server.Init(prefix, nil, noAuth, auth)
}
