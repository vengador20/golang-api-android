package middleware

import (
	"fiberapi/config"

	"github.com/gofiber/fiber/v2"
)

type Head struct {
	Name string `reqHeader:"name"`
}

func JWTUser(c *fiber.Ctx) error {
	p := new(Head)

	if err := c.ReqHeaderParser(p); err != nil {
		return err
	}

	status, err := config.VerifyJwt(p.Name)

	if err != nil {
		return c.SendStatus(404)
	}

	//si es falso significa token modificado o no authorizado
	//log.Printf("%+v\n", status)
	if !status {
		return c.SendStatus(400)
	}
	//log.Printf("%+v\n", status)

	return c.Next()
}
