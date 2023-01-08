package database

import (
	"context"
	"fiberapi/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB          *mongo.Client = DbConnect()
	DBNAME      string        = "uguia"
	TABLE_USERS string        = "users"
)

func connect(channel <-chan string) *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cliente, err := mongo.Connect(ctx, options.Client().ApplyURI(<-channel))

	if err != nil {
		panic(err)
	}
	//fmt.Println("se termino de leer el archivo .env")
	return cliente
}

func DbConnect() *mongo.Client {
	channel := make(chan string)

	go func() {
		channel <- config.ReadEnv()
	}()

	return connect(channel)
}

func GetCollection(cliente *mongo.Client, collection string) *mongo.Collection {
	coll := cliente.Database(DBNAME).Collection(collection)
	return coll
}

func DisconnectDatabase(ctx context.Context, cliente *mongo.Client) {
	if err := cliente.Disconnect(ctx); err != nil {
		panic(err)
	}
}
