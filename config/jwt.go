package config

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// type MapClaimsUser struct {
// 	email string `json:"email"`
// }

func GenerateJwt(email string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["authorized"] = true
	claims["userEmail"] = email

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	//leemos la variable de entorno con el nombre
	privateKey := os.Getenv("PRIVATEKEY_JWT")

	tokenString, err := token.SignedString([]byte(privateKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJwt(tokenHeader string) (bool, error) {
	token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unauthorized")
		}

		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found")
		}

		//leemos la variable de entorno con el nombre
		privateKey := os.Getenv("PRIVATEKEY_JWT")

		// return []byte(privateKey),nil
		return []byte(privateKey), nil
	})

	if token == nil {
		return false, nil
	}

	//parsear resultados
	if err != nil {
		return false, nil
	}

	//validar token
	if !token.Valid {
		return false, nil
	}
	return true, nil
}

func RandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]

	//encod := base64.StdEncoding.EncodeToString([]byte("82e7aa3ff12be06fbd9f2c52bd2e0f944ca"))
}
