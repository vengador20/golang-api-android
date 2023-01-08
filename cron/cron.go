package cron

import (
	"context"
	"fiberapi/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client = database.DB

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

	collClas := database.GetCollection(client, "user_clasificacion")

	lookupClas := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "clasificacion"},
		{Key: "localField", Value: "idClasificacion"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "idClasificacion"},
	}}}

	lookNivel := bson.D{{Key: "$lookup", Value: bson.M{
		"from":         "nivelEducativo",
		"localField":   "idClasificacion.idNivelEducativo",
		"foreignField": "_id",
		"as":           "idClasificacion.idNivelEducativo",
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

	return c.JSON(fiber.Map{"class": showsLoadedClas})
}
