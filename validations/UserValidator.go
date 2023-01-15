package validations

import (
	"fmt"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var (
	Validate *validator.Validate
)

type UserCredencial interface {
	User | UserRegister
}

func replaceString(str string) string {

	spl := strings.Split(str, " ")

	switch spl[0] {
	case "Password":
		return strings.Replace(str, "Password", "La contraseña", 3)
		//fmt.Println("password")

	case "Age":
		return strings.Replace(str, "Age", "La edad", 3)
		//fmt.Println("edad")

	case "Apellidos":
		return strings.Replace(str, "Apellidos", "Los Apellidos", 3)

	default:
		return str
	}
}

func ValidatorStruct[T UserCredencial](user T, trans ut.Translator) ([]string, error) {
	//var validate = validator.New()

	//var errors []*ErrorResponse
	err := Validate.Struct(user)

	var message []string

	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, v := range errs.Translate(trans) {
			//	delete(errs.Translate(trans), k)
			//	fmt.Println(k, v)
			replaceString(v)
			message = append(message, replaceString(v))
		}

		return message, fmt.Errorf("") //errs.Translate(trans),fmt.Errorf("error")
		// for _, err := range err.(validator.ValidationErrors) {
		// 	var element ErrorResponse
		// 	//err.Translate(trans)

		// 	element.FailedField = err.StructNamespace()
		// 	element.Tag = err.Tag()
		// 	element.Value = err.Param()
		// 	//err.Translate()
		// 	//fmt.Println(err.Tag())

		// 	errors = append(errors, &element)
		// }
	}
	return message, nil //validator.ValidationErrorsTranslations{},nil
	//return errors
}
