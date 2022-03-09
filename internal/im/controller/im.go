package controller

import (
	"home/internal/im/service"
	"home/pkg/resp"
	"home/pkg/utils/validate"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type im struct {
}

// imRouter current controller router
func imRouter(r fiber.Router) {
	router := r.Group("/im")
	{
		im := &im{}
		router.Get("/online", im.Online)
		router.Post("/kick", im.Kick)
		router.Post("/sendMessage", im.SendMessage)
	}
}

// Kick im
func (*im) Kick(ctx *fiber.Ctx) (err error) {
	var p service.Message
	if err = validate.StructParser(ctx, &p); err != nil {
		return
	}

	return resp.New(ctx).JSON(service.IM.Kick(p))
}

// Online im
func (*im) Online(ctx *fiber.Ctx) (err error) {
	var p service.User
	if err = validate.StructParser(ctx, &p); err != nil {
		return
	}

	return resp.New(ctx).JSON(service.IM.Online(p))
}

// SendMessage im
func (*im) SendMessage(ctx *fiber.Ctx) (err error) {
	var p service.Message
	if err = validate.StructParser(ctx, &p); err != nil {
		return
	}

	return resp.New(ctx).JSON(service.IM.SendMessage(p))
}

// ws im
func ws() fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		service.IM.WS(conn)
	})
}
