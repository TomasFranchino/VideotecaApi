package models

import (
	"gorm.io/gorm"
)

type Genero struct {
	gorm.Model        // -> el gorm.Model implementa el ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `json:"name"`
}
