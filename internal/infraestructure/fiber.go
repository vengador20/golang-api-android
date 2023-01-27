package infraestructure

import (
	"context"
	"fiberapi/internal/auth"
	adapter "fiberapi/internal/infraestructure/Adapter"
	mongodb "fiberapi/internal/infraestructure/mongo"
	"fiberapi/internal/users"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/fx"
)

func CtxWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func Db() adapter.Adapter {
	db := mongodb.Connection{}
	return adapter.AdapterMongo{Mongo: &db}
}

func NewHTTPServer(lc fx.Lifecycle, app *fiber.App, conn adapter.Adapter, cancel context.CancelFunc) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

				conn.New()

				AuthRepository := auth.NewAuthRepository(conn)
				UsersRepository := users.NewUsersRepository(conn)

				auth.NewAuthHandler(app.Group("/auth"), AuthRepository)
				users.NewUsersHandler(app.Group("/users"), UsersRepository)

				go app.Listen(":8080")
				return nil
			},
			OnStop: func(ctx context.Context) error {
				//conn.DisconnectDatabase(ctx)
				conn.Disconnect(ctx)
				cancel()
				return nil
			},
		},
	)
}

func AppMiddleware(app *fiber.App) {
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
	app.Use("/api", auth.JWTUser)
}
