package query

import (
	"context"
	"encoding/json"
	"fiberapi/config/cache"
	"fiberapi/database/models"
	"fiberapi/database/mongo"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	user = "users"
	//users []models.UserAll
)

type UserQuery interface {
	QueryUserFind(filter bson.D) []bson.M
	UserAll() models.UserAll
}

type QueryUser struct {
	Ctx context.Context
}

func (u *QueryUser) UserAll() ([]models.UserAll, error) {

	var users []models.UserAll

	cache := cache.Cache{
		Ctx: u.Ctx,
	}

	res, err := cache.GetCache(user)

	//ya existe en cache
	if err == nil {
		//fmt.Println("cache")
		json.Unmarshal([]byte(res), &users)

		return users, err
	}
	//fmt.Println("no esta en cache")

	db := mongo.GetInstance()

	coll := db.GetCollection(user)

	cursor, err := coll.Find(u.Ctx, bson.D{})

	if err != nil {
		return users, fmt.Errorf("error: %v", err)
	}

	err = cursor.All(u.Ctx, &users)

	if err != nil {
		return users, fmt.Errorf("error: %v", err)
	}

	defer cursor.Close(u.Ctx)

	cache.SaveCache(user, users)

	return users, nil
}
