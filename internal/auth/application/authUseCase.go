package application_auth

import (
	"context"
	domain_users "fiberapi/internal/auth/domain"
	mongodb "fiberapi/internal/infraestructure/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type App struct {
	domain_users.UserRepository
}

func (a *App) GetEmail(email string) domain_users.User {
	var results domain_users.User
	client := mongodb.GetInstance()
	coll := client.GetCollection("users") //mong.GetCollection(cliente, "users")
	coll.FindOne(context.Background(), bson.M{"email": email}).Decode(&results)
	return results
}

func (a *App) CreateUser(domain_users.User) {

}
