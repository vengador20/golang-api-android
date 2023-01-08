package routes

import (
	"context"
	"fiberapi/controllers"
	"fiberapi/cron"
	"fiberapi/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(router fiber.Router) {

	router.Get("/prueba", prueba)

	router.Get("/ur", cron.Semana)

	router.Get("/find", find)

	router.Get("/users", controllers.Home)

	router.Post("/xp", controllers.AumentarXp)

	router.Get("/about", func(c *fiber.Ctx) error {
		return c.JSON("exito")
	})
	router.Get("/user/:id", controllers.GetUserId)
}

func find(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	coll := database.GetCollection(cliente, "users")

	lokup := bson.D{{"$lookup", bson.D{{"from", "nivelEducativo"},
		{"localField", "idNivelEducativo"}, {"foreignField", "_id"}, {"as", "idNivelEducativo"}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{lokup})

	if err != nil {
		panic(err)
	}

	var showsLoaded []bson.M

	if err = cursor.All(ctx, &showsLoaded); err != nil {
		panic(err)
	}

	return c.JSON(showsLoaded)
}

var cliente *mongo.Client = database.DB

type nivelEducativo struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Nombre string             `bson:"nombre,omitempty"`
}

type clasificacion struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	//Nombre           string `bson:"nombre,omitempty"`
	Grupo            primitive.ObjectID `bson:"idGrupo,omitempty"`
	IdNivelEducativo primitive.ObjectID `bson:"idNivelEducativo,omitempty"`
	IdUser           primitive.ObjectID `bson:"idUser,omitempty"`
	Xp               int                `bson:"xp,omitempty"`
}

type Grupo struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Nombre string             `bson:"nombre,omitempty"`
}

func prueba(c *fiber.Ctx) error {

	ctx, canel := context.WithTimeout(context.Background(), 10*time.Second)

	defer canel()

	// collnivel := database.GetCollection(cliente, "nivelEducativo")

	// bsonNivel := []interface{}{
	// 	nivelEducativo{Nombre: "primaria"},
	// 	nivelEducativo{Nombre: "secundaria"},
	// 	nivelEducativo{Nombre: "preparatoria"},
	// }

	// collnivel.InsertMany(ctx, bsonNivel)

	//clasificacion
	collCla := database.GetCollection(cliente, "clasificacion")

	//id primaria
	id, _ := primitive.ObjectIDFromHex("63b90053641dc70bbd95087f")
	//id user efrain
	idUser, _ := primitive.ObjectIDFromHex("63b9bc4dbf4a6d0293896cb8")
	//grupos
	idUnoA, _ := primitive.ObjectIDFromHex("63bb39a02ff9dcd3b4f5eef1")
	// idUnoB, _ := primitive.ObjectIDFromHex("63bb39a02ff9dcd3b4f5eef2")
	// idUnoC, _ := primitive.ObjectIDFromHex("63bb39a02ff9dcd3b4f5eef3")

	bsonClas := []interface{}{
		clasificacion{ID: primitive.NewObjectID(), Grupo: idUnoA, IdNivelEducativo: id, IdUser: idUser, Xp: 10},
		// clasificacion{ID: primitive.NewObjectID(), Grupo: "1b", IdNivelEducativo: id},
		// clasificacion{ID: primitive.NewObjectID(), Grupo: "1c", IdNivelEducativo: id},
	}

	collCla.InsertMany(ctx, bsonClas)

	//users
	// collUser := database.GetCollection(cliente, "users")

	// //id primaria
	// id, _ := primitive.ObjectIDFromHex("63b90053641dc70bbd95087f")

	// bsonUser := bson.D{primitive.E{Key: "nombres", Value: "efrain gustavo"}, {Key: "apellidos", Value: "baizabal"},
	// 	{Key: "idNivelEducativo", Value: id}}

	// collUser.InsertOne(ctx, bsonUser)

	// collUser := database.GetCollection(cliente, "user_clasificacion")

	// //id user efrain gustavo
	// idUserEfrain, _ := primitive.ObjectIDFromHex("63b9bc4dbf4a6d0293896cb8")
	// idClasUnoA, _ := primitive.ObjectIDFromHex("63b9ba8b07e671d46f1bb32a") //1a

	// bsonUser := user_clasificacion{ID: primitive.NewObjectID(), XpUser: 10, IdUser: idUserEfrain, IdClasificacion: idClasUnoA}

	// collUser.InsertOne(ctx, bsonUser)

	//grupo
	// collG := database.GetCollection(cliente, "grupo")

	// bsonG := []interface{}{
	// 	Grupo{ID: primitive.NewObjectID(), Nombre: "1a"},
	// 	Grupo{ID: primitive.NewObjectID(), Nombre: "1b"},
	// 	Grupo{ID: primitive.NewObjectID(), Nombre: "1c"},
	// }

	// collG.InsertMany(ctx, bsonG)

	return c.JSON("exito")
}
