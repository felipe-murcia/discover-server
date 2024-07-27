package seeds

import (
	"fmt"

	"discoverco.co/server/configs"
	"discoverco.co/server/models"
)

func CreateInitialFeatures() {
	// Crear registros de categorías
	features := []models.Feature{
		{Name: "Camping", Icon: "tent", State: true},
		{Name: "Mountain", Icon: "mountain", State: true},
		{Name: "Forest", Icon: "forest", State: true},
		{Name: "Beach", Icon: "beach", State: true},
		{Name: "Water", Icon: "boat", State: true},
		{Name: "Desert", Icon: "cactus", State: true},
		{Name: "Park", Icon: "park", State: true},
		{Name: "Village", Icon: "town", State: true},
		{Name: "Other", Icon: "bowl", State: true},
	}

	// Iterar sobre las categorías y crearlas en la base de datos
	for _, feature := range features {
		configs.DB.Create(&feature)
	}

	fmt.Println("Registros iniciales de features creados correctamente.")
}
