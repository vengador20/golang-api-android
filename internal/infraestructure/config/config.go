package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func ReadEnv() string {
	//cargar el archivo .env
	if err := godotenv.Load(filepath.Join("../", ".env")); err != nil {
		log.Println("No .env file found")
	}

	//leemos la variable de entorno con el nombre
	uri := os.Getenv("MONGODB_URI")

	//si la varible de entorno es vacia mandamos mensaje
	// if uri == "" {
	// 	log.Fatal("url mongodb vacio")
	// }
	return uri
}
