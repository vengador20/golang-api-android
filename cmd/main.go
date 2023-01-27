package main

import (
	"fiberapi/internal/infraestructure"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			infraestructure.CtxWithTimeout,
			fiber.New,
			infraestructure.Db,
		),
		fx.Invoke(
			infraestructure.AppMiddleware,
			infraestructure.NewHTTPServer,
		),
	)

	app.Run()
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
