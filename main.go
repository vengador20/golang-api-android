package main

import (
	"context"
	"fiberapi/controllers"
	query "fiberapi/database/Query"
	"fiberapi/database/mongo"
	"fiberapi/middleware"
	"fiberapi/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			ctxWithTimeout,
			fiber.New,
			//database.DbConnect,
			mongo.GetInstance,
			//sqlite.Open,
		),
		fx.Invoke(
			appMiddleware,
			NewHTTPServer,
			routes.New,
		),
	)

	app.Run()
}

func ctxWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func NewHTTPServer(lc fx.Lifecycle, app *fiber.App, conn *mongo.Connection, cancel context.CancelFunc) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

				app.Post("/login", controllers.Login)
				app.Post("/register", controllers.Register)

				app.Get("/ejemplo", func(c *fiber.Ctx) error {
					ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
					defer cancel()

					db := query.QueryUser{
						Ctx: ctx,
					}

					res, err := db.UserAll()

					if err != nil {
						return c.SendStatus(400)
					}

					return c.JSON(res)
				})

				go app.Listen(":8080")
				return nil
			},
			OnStop: func(ctx context.Context) error {
				conn.DisconnectDatabase(ctx)
				cancel()
				return nil
			},
		},
	)
}

func appMiddleware(app *fiber.App) {
	//cors
	app.Use(cors.New(cors.Config{
		//AllowOrigins: "http://localhost:5500",
		AllowCredentials: true,
	}))

	//recover middleware para evitar se que se termine el programa por panic
	app.Use(recover.New())

	//compresion
	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))

	/*
	* permite que los cachés sean más eficientes
	* y ahorren ancho de banda, ya que un servidor
	* web no necesita volver a enviar una respuesta
	* completa si el contenido no ha cambiado
	 */
	app.Use(etag.New(etag.Config{Weak: true}))

	//middleware custom valid head jwt
	app.Use("/api", middleware.JWTUser)
}

// func em() {
// 	app := fiber.New()

// 	// cron := cron.New()

// 	// go cron.AddFunc("* * * * *", func() {
// 	// 	println("ejecute minuto")
// 	// })

// 	// cron.Start()
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	defer cancel()
// 	//cors
// 	app.Use(cors.New(cors.Config{
// 		//AllowOrigins: "http://localhost:5500",
// 		AllowCredentials: true,
// 	}))

// 	//recover middleware para evitar se que se termine el programa por panic
// 	app.Use(recover.New())

// 	//compresion
// 	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))

// 	/*
// 	* permite que los cachés sean más eficientes
// 	* y ahorren ancho de banda, ya que un servidor
// 	* web no necesita volver a enviar una respuesta
// 	* completa si el contenido no ha cambiado
// 	 */
// 	app.Use(etag.New(etag.Config{Weak: true}))

// 	//middleware custom valid cookie jwt
// 	app.Use("/api", middleware.JWTUser)

// 	//conexion ala base de datos
// 	db := database.DbConnect()

// 	defer db.Disconnect(ctx)

// 	app.Route("/api", routes.SetupRouter)

// 	//app.Get("/users", controllers.Home)

// 	app.Post("/login", controllers.Login)
// 	//app.Get("/us", controllers.Home)
// 	app.Get("/hola", func(c *fiber.Ctx) error {
// 		//err := config.RandomString(25)
// 		return c.JSON("hola")
// 	})

// 	app.Get("/users", controllers.Home)

// 	//app.Get("/ur", cron.Semana)

// 	// app.Get("/jwt", func(c *fiber.Ctx) error {

// 	// 	p := new(Head)

// 	// 	if err := c.ReqHeaderParser(p); err != nil {
// 	// 		return err
// 	// 	}

// 	// 	status, err := config.VerifyJwt(p.Name)

// 	// 	if err != nil {
// 	// 		return c.SendStatus(404)
// 	// 	}

// 	// 	//si es falso significa token modificado o no authorizado
// 	// 	if !status {
// 	// 		return c.SendStatus(400)
// 	// 	}

// 	// 	//log.Println(p.Name)
// 	// 	//log.Println(c.GetReqHeaders())

// 	// 	return c.JSON(p.Name)
// 	// })

// 	app.Post("/niveleducativo", controllers.ActualizarNivelEducativo)
// 	app.Post("/register", controllers.Register)
// 	app.Post("/logout", controllers.Logout)

// 	app.Listen(":8080")
// }
