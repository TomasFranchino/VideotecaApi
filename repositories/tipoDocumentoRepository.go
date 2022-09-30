package repositories

import (
	"videotecaapi/db"
	"videotecaapi/models"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

// ModelRepository ...
type TipoDocumentoRepository struct{}

func (rep TipoDocumentoRepository) Find(context *gin.Context) paginate.Page {
	db := db.DBConn
	pg := paginate.New()

	model := db.Joins("Tipo_Documentos").Model(&models.TipoDocumento{})

	return pg.With(model).Request(context.Request).Response(&[]models.TipoDocumento{})
}

func (rep TipoDocumentoRepository) Get(id int) *models.TipoDocumento {

	entity := new(models.TipoDocumento)

	db := db.DBConn
	db.First(&entity, id)

	return entity
}

func (rep TipoDocumentoRepository) Insert(entity models.TipoDocumento) uint {
	db := db.DBConn

	db.Create(&entity)

	return entity.ID
}

func (rep TipoDocumentoRepository) Update(ID int, entity models.TipoDocumento) int {

	entityToUpdate := new(models.TipoDocumento)

	db := db.DBConn
	db.First(&entityToUpdate, ID)

	result := db.Model(&entityToUpdate).Updates(map[string]interface{}{"tipo": entity.Tipo, "nombre": entity.Nombre})

	return int(result.RowsAffected)
}

func (rep TipoDocumentoRepository) Delete(ID int) int {

	entityToDelete := new(models.Genero)

	db := db.DBConn
	db.First(&entityToDelete, ID)

	result := db.Delete(&entityToDelete)

	return int(result.RowsAffected)
}

func (rep TipoDocumentoRepository) GetByDocumentName(documentName string) *models.TipoDocumento {

	entity := new(models.TipoDocumento)

	db := db.DBConn

	// Get first matched record
	db.Where("tipo = ?", documentName).First(&entity)
	// SELECT * FROM tipodocumentos WHERE nombre = 'algo_a_buscar' ORDER BY id LIMIT 1;

	return entity
}
