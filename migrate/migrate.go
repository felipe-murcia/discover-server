package main

import (
	"fmt"

	"discoverco.co/server/configs"
	"discoverco.co/server/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{})
	createInitialPersons()
}

func createInitialPersons() {
	// Crear registros de categorías
	categories := []models.Person{
		{Name: "Juan Milo", Address: "Calle 42", Phone: 37826464},
		{Name: "Camila Perez", Address: "Calle 22", Phone: 1111111},
		{Name: "Dana Parra", Address: "Calle 443", Phone: 2222222},
	}

	// Iterar sobre las categorías y crearlas en la base de datos
	for _, category := range categories {
		configs.DB.Create(&category)
	}

	fmt.Println("Registros iniciales de categorías creados correctamente.")
}
