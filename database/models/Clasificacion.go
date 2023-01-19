package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Clasificacion struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	//Nombre           string `bson:"nombre,omitempty"`
	Grupo            primitive.ObjectID `bson:"idGrupo,omitempty"`
	IdNivelEducativo primitive.ObjectID `bson:"idNivelEducativo,omitempty"`
	IdUser           primitive.ObjectID `bson:"idUser,omitempty"`
	Xp               int                `bson:"xp,omitempty"`
}
