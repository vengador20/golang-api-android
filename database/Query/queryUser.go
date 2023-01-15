package query

import (
	"context"
	"fiberapi/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserQuery interface {
	QueryUserFind(filter bson.D) []bson.M
}

type Query struct {
	Ctx    context.Context
	Client *mongo.Client
}

//var cliente *mongo.Client = database.DB

func (u *Query) QueryUserFind(filter bson.D) []bson.M {
	var db database.Database = &database.DBCon{Client: u.Client}

	coll := db.GetCollection(database.TABLE_USERS) //database.GetCollection(cliente, database.TABLE_USERS) //db.GetCollection(cliente, database.TABLE_USERS)

	cursor, err := coll.Find(u.Ctx, filter)

	if err != nil {
		println("erro")
		panic(err)
	}

	var user []bson.M

	cursor.All(u.Ctx, &user)

	defer cursor.Close(u.Ctx)

	return user
}
