package routes

import (
	"context"
	"fiberapi/controllers"

	"github.com/gofiber/fiber/v2"
)

type HandlerApi struct {
	fib *fiber.App
}

type RoutesFunction interface {
	inter(c *fiber.Ctx) error
}

func New(app *fiber.App, ctx context.Context) *HandlerApi {
	hand := HandlerApi{app}
	hand.registerRoutes()
	return &hand

}

func (h *HandlerApi) registerRoutes() {
	h.fib.Route("/api", h.SetupRouter)
}

func (h *HandlerApi) SetupRouter(router fiber.Router) {

	router.Post("/user-new-password", controllers.NewPassword)

	router.Get("/users", controllers.Home)

	router.Post("/xp", controllers.AumentarXp)

	router.Get("/about", func(c *fiber.Ctx) error {
		return c.JSON("exito")
	})
	router.Get("/user/:id", controllers.GetUserId)

}
