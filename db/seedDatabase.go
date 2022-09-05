package db

import (
	"fmt"
	"videotecaapi/models"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) {

	// Migraciones de las tablas a la base de datos.
	// Colocar una por cada tabla creada

	db.AutoMigrate(&models.Genero{})

	fmt.Println("Database migrations completed successfully")
}
