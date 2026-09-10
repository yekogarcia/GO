package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	nombre := "Santiago"
	apellido := "García"
	nombreCompleto := nombre + apellido

	println("Hola, mi nombre es", nombreCompleto)

	error := godotenv.Load()
	if error != nil {
		println("Error al cargar las variables de entorno:", error.Error())
		return
	}

	appName := os.Getenv("APP_NAME")
	appEnv := os.Getenv("APP_ENV")
	appPort := os.Getenv("APP_PORT")

	println("El nombre de la aplicación es:", appName)
	println("El entorno de la aplicación es:", appEnv)
	println("El puerto de la aplicación es:", appPort)

}
