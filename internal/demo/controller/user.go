package controller

import (
	"home/internal/demo/service"
	"home/pkg/resp"
	"home/pkg/utils/jwt"
	"home/pkg/utils/validate"

	"github.com/gofiber/fiber/v2"
)

type user struct {
}

// login ...
// @Title 用户登录
// @Description 用户登录
// @Tags user
// @Param body body service.LoginReq true "参数"
// @Success 200 {object} service.LoginRes "code==0请求成功，否则请求失败！"
// @router /login [post]
func login(ctx *fiber.Ctx) (err error) {
	var p service.LoginReq
	if err = validate.StructParser(ctx, &p); err != nil {
		return
	}

	return resp.New(ctx).JSON(service.User.Login(p))
}

// info ...
// @Title 用户信息
// @Description 用户信息
// @Tags user
// @Header 200 {string} string "Authorization:Bearer {token}"
// @Success 200 {object} jwt.User "code==0请求成功，否则请求失败！"
// @router /info [get]
func (*user) info(ctx *fiber.Ctx) (err error) {
	user := ctx.Locals("user")
	return resp.New(ctx).JSON(jwt.ParseToken(user))
}
