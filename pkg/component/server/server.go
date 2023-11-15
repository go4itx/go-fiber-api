package server

import (
	conf "home/config"
	"home/pkg/e"
	"home/pkg/resp"
	"home/pkg/utils/jwt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtWare "github.com/gofiber/jwt/v3"
)

type Config struct {
	Addr       string
	EnableTLS  bool
	KeyFile    string
	CertFile   string
	EnableCors bool
	EnableCsrf bool
	Logger     bool
	BodyLimit  int
	DelayExit  uint8
}

type Static struct {
	Prefix      string
	Root        string
	FiberStatic []fiber.Static
}

// Init server
func Init(prefix string, static []Static, noAuth func(*fiber.App), auth func(fiber.Router)) (err error) {
	var config Config
	if err = conf.Load("server.http", &config); err != nil {
		return
	}

	if config.BodyLimit == 0 {
		config.BodyLimit = 4 * 1024 * 1024 // 4MB
	}

	app := fiber.New(fiber.Config{
		BodyLimit:    config.BodyLimit,
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

	for _, v := range static {
		app.Static(v.Prefix, v.Root, v.FiberStatic...)
	}

	// custom return results
	app.Get("/", func(ctx *fiber.Ctx) error {
		return resp.New(ctx).JSON("hello world")
	})

	// Getpid
	app.Get("/pid", func(ctx *fiber.Ctx) error {
		return resp.New(ctx).JSON(os.Getpid())
	})

	if noAuth != nil {
		noAuth(app)
	}

	if auth != nil {
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
	}

	// Handle not founds
	app.Use(func(ctx *fiber.Ctx) error {
		return fiber.ErrNotFound
	})

	go func() {
		if config.EnableTLS {
			err = app.ListenTLS(config.Addr, config.CertFile, config.KeyFile)
		} else {
			err = app.Listen(config.Addr)
		}

		if err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	<-c                                             // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	if err = app.Shutdown(); err != nil {
		log.Println("err", err)
	}

	log.Println("Running cleanup tasks...")
	// Your cleanup tasks go here
	// db.Close()、redisConn.Close() ...
	time.Sleep(time.Duration(config.DelayExit) * time.Second) //延迟退出程序，确保其它协程完成任务
	log.Println("APP was successful shutdown.")
	return
}
