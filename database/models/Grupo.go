package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Grupo struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Nombre string             `bson:"nombre,omitempty"`
}
