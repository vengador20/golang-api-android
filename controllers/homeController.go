package controllers

import (
	"context"
	"fiberapi/database/models"
	mong "fiberapi/database/mongo"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageStruct struct {
	Message string
	Status  uint
}

var cliente *mongo.Client = mong.GetInstance().Client

func Home(c *fiber.Ctx) error {
	//importamos el modelo user y ponemos type para que sea de tipo users
	//type Users Users

	ctx, canel := context.WithTimeout(context.Background(), 10*time.Second)

	defer canel()
	//desconectar de la base de datos
	//defer database.DisconnectDatabase(ctx,cliente)

	coll := mong.GetCollection(cliente, "users")

	cursor, err := coll.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	var results []models.UserAll

	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	//  for _, v := range results {
	// 	cursor.Decode(&results)
	// 	output, err := json.MarshalIndent(v, "", " ")
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Printf("%s\n",output)
	//  }
	return c.JSON(&results)
}

func GetUserId(c *fiber.Ctx) error {

	id := c.Params("id")
	//importamos el modelo user y ponemos type para que sea de tipo users
	var Users models.User

	ctx, canel := context.WithTimeout(context.Background(), 10*time.Second)

	defer canel()

	//desconectar de la base de datos
	//defer database.DisconnectDatabase(ctx,cliente)

	coll := mong.GetCollection(cliente, "users")

	objId, _ := primitive.ObjectIDFromHex(id)

	//decodificamos la consulta
	// *la coleccion la obtiene users ya que señala a users y devuelve el error
	err := coll.FindOne(ctx, bson.M{"_id": objId}).Decode(&Users)

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(&Users)
}

type UserXp struct {
}

func AumentarXp(c *fiber.Ctx) error {

	var user = new(models.UserAll)

	err := c.BodyParser(user)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(user)
}
