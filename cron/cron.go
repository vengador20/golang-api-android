package cron

import (
	"context"
	"fiberapi/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client = database.DB

type users struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Nombre string             `bson:"nombre,omitempty"`
}

type nivelEducativo struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	Nombres          string             `json:"nombres" bson:"nombres,omitempty"`
	Apellidos        string             `json:"apellidos" bson:"apellidos,omitempty"`
	IdNivelEducativo primitive.ObjectID `json:"idNivelEducativo" bson:"idNivelEducativo,omitempty"`
}

type clasificacion struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	//Nombre           string `bson:"nombre,omitempty"`
	Grupo            []Grupo          `json:"idGrupo" bson:"idGrupo,omitempty"`
	IdNivelEducativo []nivelEducativo `json:"idNivelEducativo" bson:"idNivelEducativo,omitempty"`
	IdUser           []users          `json:"idUser" bson:"idUser,omitempty"`
	Xp               int              `json:"xp" bson:"xp,omitempty"`
}

type Grupo struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Nombre string             `json:"nombre" bson:"nombre,omitempty"`
}

func Semana(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	// coll := database.GetCollection(client, database.TABLE_USERS)

	// lookup := bson.D{{"$lookup", bson.D{{"from", "nivelEducativo"},
	// 	{"localField", "idNivelEducativo"}, {"foreignField", "_id"}, {"as", "idNivelEducativo"}}}}

	// cursor, err := coll.Aggregate(ctx, lookup)

	// if err != nil {
	// 	panic(err)
	// }

	// var showsLoaded []bson.M

	// if err = cursor.All(ctx, &showsLoaded); err != nil {
	// 	panic(err)
	// }

	collClas := database.GetCollection(client, "clasificacion")

	lookupClas := bson.D{{Key: "$lookup", Value: bson.M{
		"from":         "grupo",
		"localField":   "idGrupo",
		"foreignField": "_id",
		"as":           "idGrupo",
	}}}

	lookNivel := bson.D{{Key: "$lookup", Value: bson.M{
		"from":         "nivelEducativo",
		"localField":   "idNivelEducativo",
		"foreignField": "_id",
		"as":           "idNivelEducativo",
	}}}

	lookUser := bson.D{{Key: "$lookup", Value: bson.M{
		"from":         "users",
		"localField":   "idUser",
		"foreignField": "_id",
		"as":           "idUser",
	}}}

	cursorClas, err := collClas.Aggregate(ctx, mongo.Pipeline{lookupClas, lookNivel, lookUser})

	if err != nil {
		panic(err)
	}

	var showsLoadedClas []bson.M

	if err = cursorClas.All(ctx, &showsLoadedClas); err != nil {
		panic(err)
	}

	defer cursorClas.Close(ctx)

	return c.JSON(fiber.Map{"class": showsLoadedClas})
}
