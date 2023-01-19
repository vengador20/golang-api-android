package controllers

import (
	"context"
	"fiberapi/config"
	"fiberapi/database/models"
	mong "fiberapi/database/mongo"
	"fiberapi/validations"
	"time"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translations "github.com/go-playground/validator/v10/translations/es"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	uni *ut.UniversalTranslator
)

type Respuesta struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
	TOKEN   string   `json:"token,omitempty"`
}

func Login(c *fiber.Ctx) error {

	user := new(models.User)

	esTrans := es.New()

	enLocate := en.New()

	uni = ut.New(enLocate, enLocate, esTrans)

	trans, _ := uni.GetTranslator("es")

	validations.Validate = validator.New()

	es_translations.RegisterDefaultTranslations(validations.Validate, trans)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	errors, err := validations.ValidatorStruct(*user, trans)

	if err != nil {
		res := Respuesta{Message: "fallo", Errors: errors}
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	//ctx para que cuando el usuario interrumpe la solicitud no se quede
	//y terminanar el proceso y evitar desgaste de recursos
	ctx, canel := context.WithTimeout(context.Background(), 10*time.Second)

	defer canel()

	if err != nil {
		//res := Respuesta{Message: "fallo", Errors: err}
		return c.JSON(err)
	}

	var results models.User

	coll := mong.GetCollection(cliente, mong.TABLE_USERS)

	coll.FindOne(ctx, bson.M{"email": user.Email}).Decode(&results)

	//comparar el password del usuario introducido con el password de la base de datos
	err = bcrypt.CompareHashAndPassword([]byte(results.Password), []byte(user.Password))

	if err != nil {
		return c.JSON(Respuesta{Message: "Correo o Contraseña no coinciden"})
	}

	token, err := config.GenerateJwt(user.Email)
	if err != nil {
		return c.JSON("no generate jwt")
	}

	res := Respuesta{Message: "Exito", TOKEN: token}
	return c.Status(fiber.StatusOK).JSON(res)
}

func Register(c *fiber.Ctx) error {

	user := new(models.UserRegister)

	esTrans := es.New()

	enLocate := en.New()

	uni = ut.New(enLocate, enLocate, esTrans)

	trans, _ := uni.GetTranslator("es")

	validations.Validate = validator.New()

	es_translations.RegisterDefaultTranslations(validations.Validate, trans)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	errors, err := validations.ValidatorStruct(*user, trans)

	if err != nil {
		res := Respuesta{Message: "fallo", Errors: errors}
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	ctx, canel := context.WithTimeout(context.Background(), 10*time.Second)

	defer canel()

	//encriptacion de la contraseña
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	//bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))

	if err != nil {
		return c.JSON(Respuesta{Message: "Error al Generar la contraseña"})
	}

	coll := mong.GetCollection(cliente, mong.TABLE_USERS)

	bsonUser := models.UserRegister{
		ID:        primitive.NewObjectID(),
		Nombres:   user.Nombres,
		Apellidos: user.Apellidos,
		Email:     user.Email,
		Password:  string(password),
	}

	// bson.D{
	// 	primitive.E{Key: "nombres", Value: user.Nombres},
	// 	{Key: "apellidos", Value: user.Apellidos}, {Key: "email", Value: user.Email},
	// 	{Key: "password", Value: password},
	// }

	_, err = coll.InsertOne(ctx, bsonUser)

	if err != nil {
		return c.JSON(err)
	}

	//fmt.Println(res)
	res := Respuesta{Message: "Se ha creado el Usuario correctamente"}

	return c.Status(fiber.StatusOK).JSON(res)

}

func Logout(c *fiber.Ctx) error {
	/*
		cookie := new(fiber.Cookie)
		cookie.Name = "token"
		cookie.Value = ""
		cookie.Expires = time.Now().Add(-3 * time.Second)
		cookie.HTTPOnly = true
		cookie.Secure = true
		cookie.Path = "/"
		//cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.SameSite = "none"

		c.Cookie(cookie)*/

	return c.JSON(Respuesta{Message: "cerrar sesion"})
}

func ActualizarNivelEducativo(c *fiber.Ctx) error {
	user := new(models.UserRegister)

	err := c.BodyParser(user)
	if err != nil {
		return err
	}

	ctx, canel := context.WithTimeout(context.Background(), 10*time.Second)

	defer canel()

	coll := cliente.Database("uguia").Collection(mong.TABLE_USERS)

	filter := bson.M{"email": user.Email}

	update := bson.M{"$set": bson.M{"nivelEducativo": user.IdNivelEducativo}}

	coll.UpdateOne(ctx, filter, update)

	res := Respuesta{Message: "exito"}
	return c.JSON(res)
}
