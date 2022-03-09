package controller

import (
	"home/internal/im/service"
	"home/pkg/resp"
	"home/pkg/utils/jwt"
	"home/pkg/utils/validate"

	"github.com/gofiber/fiber/v2"
)

type admin struct {
}

// adminRouter current controller router
func adminRouter(r fiber.Router) {
	router := r.Group("/admin")
	{
		admin := &admin{}
		router.Get("/", admin.info)
	}
}

// login admin
func login(ctx *fiber.Ctx) (err error) {
	var p service.LoginReq
	if err = validate.StructParser(ctx, &p); err != nil {
		return
	}

	return resp.New(ctx).JSON(service.Admin.Login(p))
}

// info admin
func (*admin) info(ctx *fiber.Ctx) (err error) {
	user := ctx.Locals("user")
	return resp.New(ctx).JSON(jwt.ParseToken(user))
}
