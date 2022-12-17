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

// Kick ...
// @Title 踢人
// @Description 踢人下线
// @Tags im
// @Param body body service.Message true "参数"
// @Success 200 {object} resp.Ret "code==0请求成功，否则请求失败！"
// @router /kick [post]
func (*im) Kick(ctx *fiber.Ctx) (err error) {
	var p service.Message
	if err = validate.StructParser(ctx, &p); err != nil {
		return
	}

	return resp.New(ctx).JSON(service.IM.Kick(p))
}

// Online ...
// @Title 在线
// @Description 在线用户
// @Tags im
// @Param body body service.User true "参数"
// @Success 200 {object} []service.User "code==0请求成功，否则请求失败！"
// @router /online [get]
func (*im) Online(ctx *fiber.Ctx) (err error) {
	var p service.User
	if err = validate.StructParser(ctx, &p); err != nil {
		return
	}

	return resp.New(ctx).JSON(service.IM.Online(p))
}

// SendMessage ...
// @Title 发送
// @Description 发送消息
// @Tags im
// @Param body body service.Message true "参数"
// @Success 200 {object} resp.Ret "code==0请求成功，否则请求失败！"
// @router /sendMessage [post]
func (*im) SendMessage(ctx *fiber.Ctx) (err error) {
	var p service.Message
	if err = validate.StructParser(ctx, &p); err != nil {
		return
	}

	return resp.New(ctx).JSON(service.IM.SendMessage(p))
}

// ws ...
func ws() fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		service.IM.WS(conn)
	})
}
