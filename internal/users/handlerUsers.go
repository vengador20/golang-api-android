package users

import (
	"context"
	"fiberapi/internal/infraestructure/config"
	"fiberapi/internal/users/domain"
	"fiberapi/validations"
	"time"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translations "github.com/go-playground/validator/v10/translations/es"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type head struct {
	Name string `json:"name"`
}

type HandlerUsers struct {
	service domain.UsersRepository //adapter.Adapter
}

func NewUsersHandler(route fiber.Router, adapter domain.UsersRepository) {

	handler := &HandlerUsers{
		service: adapter,
	}

	route.Post("/new-password", handler.NewPassword)

}

func (h *HandlerUsers) NewPassword(c *fiber.Ctx) error {

	var body domain.BodyNewPassword
	var head head
	var user domain.Users

	erro := c.ReqHeaderParser(&head)

	if erro != nil {
		panic(erro)
	}

	err := c.BodyParser(&body)

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	claims := config.ExtractClaims(head.Name)

	user, err = h.service.GetEmail(ctx, claims["userEmail"].(string))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "fallo", // "errors": err.Error(),
		})
	}

	// verificar contraseña
	err = bcrypt.CompareHashAndPassword(primitive.Binary(user.Password).Data, []byte(body.OldPassword))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Contraseña invalida", //"errors": err.Error(),
		})
	}

	esTras := es.New()
	enLocate := en.New()

	uni := ut.New(enLocate, enLocate, esTras)

	trans, _ := uni.GetTranslator("es")
	validations.Validate = validator.New()

	es_translations.RegisterDefaultTranslations(validations.Validate, trans)

	errors, err := validations.ValidatorStructNewPassword(body, trans)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "fallo", "errors": errors,
		})
	}

	//generar nueva contraseña
	password, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "fallo", // "errors": err.Error(),
		})
	}

	err = h.service.NewPassword(ctx, string(password), user.Email)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "fallo", "errors": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "La contraseña se modifico"})
}
