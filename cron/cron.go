package cron

import (
	"context"
	"fiberapi/config"
	"fiberapi/database"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client = database.DB

type nivelEducativo struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Nombre string             `bson:"nombre,omitempty"`
	//IdNivelEducativo primitive.ObjectID `json:"idNivelEducativo" bson:"idNivelEducativo"`
}

type users struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	Nombres          string             `json:"nombres" bson:"nombres,omitempty"`
	Apellidos        string             `json:"apellidos" bson:"apellidos,omitempty"`
	IdNivelEducativo primitive.ObjectID `json:"idNivelEducativo" bson:"idNivelEducativo,omitempty"`
}

type clasificacion struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	//Nombre           string `bson:"nombre,omitempty"`
	Grupo            primitive.ObjectID/*[]Grupo */ `json:"idGrupo" bson:"idGrupo,omitempty"`
	IdNivelEducativo primitive.ObjectID/*[]nivelEducativo*/ `json:"idNivelEducativo" bson:"idNivelEducativo,omitempty"`
	IdUser           primitive.ObjectID/*[]users*/ `json:"idUser" bson:"idUser,omitempty"`
	Xp               int `json:"xp" bson:"xp,omitempty"`
}

type Grupo struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Nombre string             `json:"nombre" bson:"nombre,omitempty"`
}

type Header struct {
	Email string `json:"email" bson:"email,omitempty"`
	Name  string `json:"name" bson:"name,omitempty"`
}

func Clas(c *fiber.Ctx) error {
	head := new(Header)

	if err := c.ReqHeaderParser(head); err != nil {
		return err
	}

	claims := config.ExtractClaims(head.Name)

	fmt.Printf("claims: %v\n", claims)
	emailToken := fmt.Sprint(claims["userEmail"])

	fmt.Printf("em: %v\n", emailToken)

	//var wg sync.WaitGroup
	//var ch = make(chan int)
	var chUser = make(chan primitive.ObjectID)

	var clasificacionUser clasificacion

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	go func() {
		var user users
		coll := database.GetCollection(client, database.TABLE_USERS)

		bson := bson.M{"email": emailToken}

		coll.FindOne(ctx, bson).Decode(&user)

		chUser <- user.IdNivelEducativo
	}()

	coll := database.GetCollection(client, database.TABLE_CLASIFICACION)

	filter := bson.M{"idNivelEducativo": <-chUser}

	coll.FindOne(ctx, filter).Decode(&clasificacionUser)

	fmt.Printf("clasificacionUser: %v\n", clasificacionUser)

	collClas := database.GetCollection(client, database.TABLE_CLASIFICACION)

	filterUs := bson.D{{Key: "$match", Value: bson.D{{Key: "idGrupo", Value: clasificacionUser.Grupo}}}}

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

	cursor, err := collClas.Aggregate(ctx, mongo.Pipeline{filterUs, lookNivel, lookUser})

	if err != nil {
		panic(err)
	}

	var showsLoadedClas []bson.M

	if err := cursor.All(ctx, &showsLoadedClas); err != nil {
		panic(err)
	}

	return c.JSON(showsLoadedClas)

}

func Semana(c *fiber.Ctx) error {

	head := new(Header)

	if err := c.ReqHeaderParser(head); err != nil {
		return err
	}

	claims := config.ExtractClaims(head.Name)

	fmt.Printf("claims: %v\n", claims)
	emailToken := fmt.Sprint(claims["userEmail"])

	fmt.Printf("em: %v\n", emailToken)

	//var wg sync.WaitGroup
	var ch = make(chan int)
	var chUser = make(chan primitive.ObjectID)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	go func() {
		var user users
		coll := database.GetCollection(client, database.TABLE_USERS)

		bson := bson.M{"email": emailToken}

		coll.FindOne(ctx, bson).Decode(&user)

		chUser <- user.IdNivelEducativo
	}()

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

	fmt.Printf("chUser: %v\n", <-chUser)
	//filter := bson.D{{Key: "idNivelEducativo", Value: <-chUser}}

	//nivelEduc := <-chUser

	//fmt.Printf("nivelEduc: %v\n", nivelEduc.IdNivelEducativo)
	//wg.Add(1)
	go func() {
		//count := bson.D{{Key: "$count", Value: "nombres"}}
		cursor, err := collClas.EstimatedDocumentCount(ctx)

		if err != nil {
			panic(err)
		}

		//println(cursor)
		ch <- int(cursor)
		//defer wg.Done()
	}()

	cursorClas, err := collClas.Aggregate(ctx, mongo.Pipeline{lookupClas, lookNivel, lookUser})

	if err != nil {
		panic(err)
	}

	var showsLoadedClas []bson.M

	if err = cursorClas.All(ctx, &showsLoadedClas); err != nil {
		panic(err)
	}

	defer cursorClas.Close(ctx)

	//defer wg.Wait()

	return c.JSON(fiber.Map{"clasificacion": showsLoadedClas, "count": <-ch})
}

func Sem(c *fiber.Ctx) error {

	var nivelEducativo []nivelEducativo
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	coll := database.GetCollection(client, database.TABLE_NIVELEDUCATIVO)

	cursor, err := coll.Find(ctx, bson.D{})

	if err != nil {
		panic(err)
	}

	if err := cursor.All(ctx, &nivelEducativo); err != nil {
		panic(err)
	}

	for _, v := range nivelEducativo {
		cronClasificacion(v.ID)
	}

	return c.JSON("hola 222")
}

func cronClasificacion(id primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	filter := bson.D{{Key: "idNivelEducativo", Value: id}}

	coll := database.GetCollection(client, database.TABLE_USERS)

	countUsers, err := coll.CountDocuments(ctx, filter)

	if err != nil {
		panic(err)
	}

	collUsers := database.GetCollection(client, database.TABLE_USERS)

	cursorUsers, err := collUsers.Find(ctx, filter)

	if err != nil {
		panic(err)
	}

	var users, agregarUsers []users

	if err := cursorUsers.All(ctx, &users); err != nil {
		panic(err)
	}

	var (
		count       int
		grupos      []string
		grupoNombre = "abcdefghijklmnñopqrstuvwxyz"
		grupoNumero = "123456789"
	)

	nombres := strings.Split(grupoNombre, "")
	numeros := strings.Split(grupoNumero, "")

	for _, v := range numeros {
		for _, vN := range nombres {
			count++
			if count >= int(countUsers)+1 {
				//count = 0
				break
			}
			grupos = append(grupos, v+vN)
		}
	}

	count = 0

	for i := 0; i < int(countUsers); i++ {
		if i >= 25 {
			gruposAdd(agregarUsers, grupos[count])
			println(count)
			count++
			agregarUsers = nil
		}
		agregarUsers = append(agregarUsers, users[i])

	}

	if agregarUsers != nil {
		gruposAdd(agregarUsers, grupos[0])
	}
}

func gruposAdd(user []users, grupo string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collGrup := database.GetCollection(client, database.TABLE_GRUPO)

	id := primitive.NewObjectID()
	bson := Grupo{
		ID:     id,
		Nombre: grupo,
	}

	collGrup.InsertOne(ctx, bson)

	data := make([]interface{}, len(user))

	for i := range user {
		data[i] = clasificacion{
			ID:               primitive.NewObjectID(),
			Grupo:            id,
			IdUser:           user[i].ID,
			IdNivelEducativo: user[i].IdNivelEducativo,
			Xp:               0,
		}
	}

	collClas := database.GetCollection(client, database.TABLE_CLASIFICACION)

	collClas.InsertMany(ctx, data)

}
