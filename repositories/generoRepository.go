package repositories

import (
	"videotecaapi/db"
	"videotecaapi/models"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

// ModelRepository ...
type GeneroRepository struct{}

func (rep GeneroRepository) Find(context *gin.Context) paginate.Page {
	db := db.DBConn
	pg := paginate.New()

	model := db.Joins("Generos").Model(&models.Genero{})

	return pg.With(model).Request(context.Request).Response(&[]models.Genero{})
}

func (rep GeneroRepository) Get(id int) *models.Genero {

	entity := new(models.Genero)

	db := db.DBConn
	db.First(&entity, id)

	return entity
}

func (rep GeneroRepository) Insert(entity models.Genero) uint {
	db := db.DBConn

	db.Create(&entity)

	return entity.ID
}

func (rep GeneroRepository) Update(ID int, entity models.Genero) int {

	entityToUpdate := new(models.Genero)

	db := db.DBConn
	db.First(&entityToUpdate, ID)

	result := db.Model(&entityToUpdate).Update("name", entity.Nombre)

	return int(result.RowsAffected)
}

func (rep GeneroRepository) Delete(ID int) int {

	entityToDelete := new(models.Genero)

	db := db.DBConn
	db.First(&entityToDelete, ID)

	result := db.Delete(&entityToDelete)

	return int(result.RowsAffected)
}
