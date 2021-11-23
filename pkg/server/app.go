package server

import (
	"home/pkg/code/e"
	"home/pkg/resp"
	"home/pkg/utils/conf"
	"home/pkg/utils/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtWare "github.com/gofiber/jwt/v3"
)

type Config struct {
	Addr       string
	EnableCors bool
	EnableCsrf bool
	Logger     bool
}

// Init server
func Init(prefix string, static map[string]string, noAuth func(*fiber.App), auth func(fiber.Router)) (err error) {
	var config Config
	if err = conf.Load("server.http", &config); err != nil {
		return
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: resp.ErrorHandler,
	})

	app.Use(recover.New())
	if config.Logger {
		app.Use(logger.New())
	}

	if config.EnableCsrf {
		app.Use(csrf.New())
	}

	if config.EnableCors {
		app.Use(cors.New(cors.Config{
			AllowCredentials: true,
		}))
	}

	for k, v := range static {
		app.Static(k, v)
	}

	// custom return results
	app.Get("/", func(ctx *fiber.Ctx) error {
		return resp.New(ctx).JSON("hello world")
	})

	noAuth(app)
	router := app.Group(prefix, jwtWare.New(jwtWare.Config{
		SigningKey: []byte(jwt.Signing),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err != nil {
				err = e.NewError(fiber.StatusUnauthorized, err.Error())
			}

			return err
		},
	}))

	auth(router)
	// Handle not founds
	app.Use(func(ctx *fiber.Ctx) error {
		return fiber.ErrNotFound
	})

	err = app.Listen(config.Addr)
	return
}