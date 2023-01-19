package mongo

import (
	"context"
	"fiberapi/config"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	conn *Connection
	one  sync.Once
	//DB                   *mongo.Client = DbConnect()
	DBNAME               string = "uguia"
	TABLE_USERS          string = "users"
	TABLE_NIVELEDUCATIVO string = "nivelEducativo"
	TABLE_CLASIFICACION  string = "clasificacion"
	TABLE_GRUPO          string = "grupo"
)

type Database interface {
	GetCollection(collection string) *mongo.Collection
}

type Connection struct {
	Client *mongo.Client
}

// patron singleton
func GetInstance() *Connection {
	one.Do(func() {
		conn = &Connection{
			Client: DbConnect(),
		}
	})

	return conn
}

func connect(channel <-chan string) *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cliente, err := mongo.Connect(ctx, options.Client().ApplyURI(<-channel))

	if err != nil {
		panic(err)
	}
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

func (c *Connection) GetCollection(collection string) *mongo.Collection {
	coll := c.Client.Database(DBNAME).Collection(collection)
	return coll
}

func (c *Connection) DisconnectDatabase(ctx context.Context) {
	if err := c.Client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
