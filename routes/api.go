package routes

import (
	"context"
	"fiberapi/controllers"
	"fiberapi/cron"
	"fiberapi/database"
	query "fiberapi/database/Query"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(router fiber.Router) {

	router.Get("/prueba", prueba)

	router.Get("/cron", cron.Sem)

	router.Get("/ur", cron.Semana)

	router.Get("/clasificacion", cron.Clas)

	router.Post("/user-new-password", controllers.NewPassword)

	router.Get("/find", find)

	router.Get("/users", controllers.Home)

	router.Post("/xp", controllers.AumentarXp)

	router.Get("/about", func(c *fiber.Ctx) error {
		return c.JSON("exito")
	})
	router.Get("/user/:id", controllers.GetUserId)

	router.Get("/interface", inter)

}

func inter(c *fiber.Ctx) error {

	//var user validations.User
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//var db database.DBCon = database.DBCon{Client: cliente}
	var userQuery query.Query = query.Query{
		Client: cliente,
		Ctx:    ctx,
	}

	filter := bson.D{}

	find := userQuery.QueryUserFind(filter)

	fmt.Printf("find: %v\n", find)

	return c.JSON(find)
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

	//nivel educativo
	// collnivel := database.GetCollection(cliente, "nivelEducativo")

	// bsonNivel := []interface{}{
	// 	nivelEducativo{ID: primitive.NewObjectID(), Nombre: "primaria"},
	// 	nivelEducativo{ID: primitive.NewObjectID(), Nombre: "secundaria"},
	// 	nivelEducativo{ID: primitive.NewObjectID(), Nombre: "preparatoria"},
	// }

	// collnivel.InsertMany(ctx, bsonNivel)

	//clasificacion
	collCla := database.GetCollection(cliente, "clasificacion")

	//id primaria
	id, _ := primitive.ObjectIDFromHex("63bba5958cc38ef1a27219b1")
	//id user efrain
	idUser, _ := primitive.ObjectIDFromHex("63bba7a524035430f2333cb9")
	//grupos
	idUnoA, _ := primitive.ObjectIDFromHex("63bba5cf5e1338be88402909")
	// idUnoB, _ := primitive.ObjectIDFromHex("63bb39a02ff9dcd3b4f5eef2")
	// idUnoC, _ := primitive.ObjectIDFromHex("63bb39a02ff9dcd3b4f5eef3")

	bsonClas := []interface{}{
		clasificacion{ID: primitive.NewObjectID(), Grupo: idUnoA, IdNivelEducativo: id, IdUser: idUser, Xp: 10},
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
