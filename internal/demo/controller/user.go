package controller

import (
	"github.com/gofiber/fiber/v2"
	"home/internal/demo/service"
	"home/pkg/resp"
	"home/pkg/utils/jwt"
	"home/pkg/utils/validate"
)

type user struct {
}

// userRouter current controller router
func userRouter(r fiber.Router) {
	router := r.Group("/user")
	{
		user := &user{}
		router.Get("/", user.info)
	}
}

// login user
func login(ctx *fiber.Ctx) (err error) {
	var p service.ParamLogin
	if err = validate.StructParser(ctx, &p); err != nil {
		return
	}

	return resp.New(ctx).JSON(service.User.Login(p))
}

// info user
func (*user) info(ctx *fiber.Ctx) (err error) {
	user := ctx.Locals("user")
	return resp.New(ctx).JSON(jwt.ParseToken(user))
}
