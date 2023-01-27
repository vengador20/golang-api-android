package domain_users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Email    string           `json:"email" validate:"required,email" bson:"email"`
	Password primitive.Binary `json:"password" validate:"required,min=8" bson:"password"`
}

type UserRegister struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	Nombres          string             `json:"nombres" validate:"required,min=3,max=32" bson:"nombres"`
	Email            string             `json:"email" validate:"required,email" bson:"email"`
	Apellidos        string             `json:"apellidos" validate:"required" bson:"apellidos"`
	Password         string             `json:"password" validate:"required,min=8" bson:"password"`
	IdNivelEducativo primitive.ObjectID `json:"idNivelEducativo" bson:"idNivelEducativo"`
}
