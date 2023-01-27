package validations

import (
	"fiberapi/database/models"
	"fiberapi/internal/users/domain"
	"fmt"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var (
	Validate *validator.Validate
)

type UserCredencial interface {
	models.User | models.UserRegister
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

func ValidatorStructNewPassword(user domain.BodyNewPassword, trans ut.Translator) ([]string, error) {
	err := Validate.Struct(user)

	var message []string

	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, v := range errs.Translate(trans) {
			replaceString(v)
			message = append(message, replaceString(v))
		}

		return message, fmt.Errorf("")
	}
	return message, nil
}
