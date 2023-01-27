package adapter

import (
	"context"
	mongodb "fiberapi/internal/infraestructure/mongo"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type AdapterMongo struct {
	Mongo *mongodb.Connection
}

func (m AdapterMongo) New() {
	m.Mongo.GetInstance()
}

func (m AdapterMongo) Disconnect(ctx context.Context) {

}

func (m AdapterMongo) FindId(ctx context.Context, id string) (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}
func (m AdapterMongo) FindEmail(ctx context.Context, email string) (map[string]interface{}, error) {

	var res map[string]interface{}

	client := m.Mongo.GetInstance()

	m.Mongo.Client = client.Client

	coll := m.Mongo.GetCollection("users")

	coll.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&res)

	if len(res) == 0 {
		return res, fmt.Errorf("resultado vacio")
	}

	return res, nil
}

func (m AdapterMongo) UpdateOne(ctx context.Context, filter, update map[string]interface{}) error {
	client := m.Mongo.GetInstance()

	m.Mongo.Client = client.Client

	coll := m.Mongo.GetCollection("users")

	keysFilter := make([]string, 0, len(filter))
	for k, v := range filter {
		keysFilter = append(keysFilter, k)
		keysFilter = append(keysFilter, v.(string))
	}

	keys := make([]string, 0, len(update))
	for k, v := range update {
		keys = append(keys, k)
		keys = append(keys, v.(string))
	}

	_, err := coll.UpdateOne(ctx, bson.M{keysFilter[0]: keysFilter[1]}, bson.M{"$set": bson.M{keys[0]: keys[1]}})

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
