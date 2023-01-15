package controllers

import (
	"context"
	"fiberapi/config"
	"fiberapi/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type body struct {
	Old string `json:"old"`
	New string `json:"new"`
}

type head struct {
	Name string `json:"name"`
}

type users struct {
	//ID        primitive.ObjectID `json:"id" bson:"_id"`
	Nombres   string `json:"nombres" validate:"required,min=3,max=32" bson:"nombres"`
	Email     string `json:"email" validate:"required,email" bson:"email"`
	Apellidos string `json:"apellidos" validate:"required" bson:"apellidos"`
	//Xp        int    `json:"xp,omitempty" bson:"xp"`
	Password string `json:"password" validate:"required,min=8" bson:"password"`
	//NivelEducativo string `json:"nivelEducativo,omempty" bson:"nivelEducativo"`
}

func NewPassword(c *fiber.Ctx) error {

	var body body
	var head head
	var user users

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

	collUser := database.GetCollection(cliente, database.TABLE_USERS)

	filterUser := bson.M{"email": claims["userEmail"]}

	collUser.FindOne(ctx, filterUser).Decode(&user)

	//verificar contraseña
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Old))

	if err != nil {
		return c.JSON(Respuesta{Message: "Contraseña invalida"})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(body.New), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	coll := database.GetCollection(cliente, database.TABLE_USERS)

	filter := bson.M{"email": claims["userEmail"]}

	update := bson.M{"password": string(password)}

	coll.UpdateOne(ctx, filter, update)

	return c.JSON(Respuesta{Message: "Se la contraseña se modifico"})
}
