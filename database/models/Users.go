package models

import (
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type Users struct {
// 	ID   primitive.ObjectID `bson:"_id"`
// 	Nombre string             `bson:"nombre"`
// 	Age  int                `bson:"age"`
// }

/**
	* patron singleton utilizar cuando se necesita un solo usuario
 **/

var (
	p   *User
	one sync.Once
)

type User struct {
	Email    string `json:"email" validate:"required,email" bson:"email"`
	Password string `json:"password" validate:"required,min=8" bson:"password"`
}

type UserRegister struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	Nombres          string             `json:"nombres" validate:"required,min=3,max=32" bson:"nombres"`
	Email            string             `json:"email" validate:"required,email" bson:"email"`
	Apellidos        string             `json:"apellidos" validate:"required" bson:"apellidos"`
	Password         string             `json:"password" validate:"required,min=8" bson:"password"`
	IdNivelEducativo primitive.ObjectID `json:"idNivelEducativo" bson:"idNivelEducativo"`
}

type UserAll struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Nombres string             `json:"nombres" validate:"required,min=3,max=32" bson:"nombres"`
	//Email            string             `json:"email" validate:"required,email" bson:"email"`
	Apellidos string `json:"apellidos" validate:"required" bson:"apellidos"`
	//Password         string             `json:"password" validate:"required,min=8" bson:"password"`
	IdNivelEducativo primitive.ObjectID `json:"idNivelEducativo" bson:"idNivelEducativo"`
}

/**
	* cuando el usuario no existe creamos un tipo de usuario y
	* si no existe regresamos un puntero asi el usuari creado
	* para eso se utiliza sync.Once del metodo Do se
	* ejecutara solo una vez (!!importante que se utiliza en esa misma variable
	* si se crea otra varible se podra utilizar normalmente como si fuera otro)
**/
func NewUser() *User {
	one.Do(func() {
		p = &User{}
	})

	return p
}
