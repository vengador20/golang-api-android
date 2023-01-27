package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	//ID        primitive.ObjectID `json:"id" bson:"_id"`
	Nombres   string `json:"nombres" validate:"required,min=3,max=32" bson:"nombres"`
	Email     string `json:"email" validate:"required,email" bson:"email"`
	Apellidos string `json:"apellidos" validate:"required" bson:"apellidos"`
	//Xp        int    `json:"xp,omitempty" bson:"xp"`
	Password primitive.Binary `json:"password" validate:"required,min=8" bson:"password"`
	//NivelEducativo string `json:"nivelEducativo,omempty" bson:"nivelEducativo"`
}

type BodyNewPassword struct {
	OldPassword string `json:"oldpassword" validate:"required,min=8"`
	NewPassword string `json:"newpassword" validate:"required,min=8"`
}
