package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"videotecaapi/models"
	"videotecaapi/repositories"
	"videotecaapi/utils"
)

type TipoDocumentoController struct{}

// Get godoc
// @Summary      Recupera un tipo de documento
// @Description  Obtiene algo dado un id
// @Tags         tipo documento
// @Accept       json
// @Produce      json
// @Param        tipoDocumentoID   path      int  true  "Id"
// @Success      200  {object}  models.TipoDocumento
// @Failure      400  {object}  models.Error
// @Failure      404  {object}  models.Error
// @Failure      500  {object}  models.Error
// @Router       /genero/{generoID} [get]
func (controller TipoDocumentoController) Get(context *gin.Context) {
	ID := context.Param("tipoDocumentoID")

	id, err := strconv.Atoi(ID) // se convierte un string a int
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	entityRep := new(repositories.TipoDocumentoRepository)
	entity := entityRep.Get(id)

	if entity.ID == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, entity)
	}
}

// Find godoc
// @Summary Busca algún tipo de documento
// @Description Obtiene un conjunto de generos aplicando filtros varios
// @Tags     tipo documento
// @Accept   json
// @Produce  json
// @Param   size     query    int     false        "Size"
// @Param   page     query    int     false        "Page"
// @Param   sort     query    string     false        "List fields"
// @Param   filter     query    string     false        "Filter"
// @Success 200 {object} models.TipoDocumento
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/models [get]
func (controller TipoDocumentoController) Find(context *gin.Context) {

	entityRep := new(repositories.TipoDocumentoRepository)
	entities := entityRep.Find(context)
	context.JSON(http.StatusOK, entities)
}

// Create godoc
// @Summary Crea un tipo de documento
// @Description dado un json con los datos de un tipo de documento, éste es persistido
// @Tags     tipo documento
// @Accept   json
// @Produce  json
// @Success 200 {int} id creado
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/models [post]
func (controller TipoDocumentoController) Create(context *gin.Context) {

	entity := new(models.TipoDocumento)

	if err := context.BindJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error()))
		return
	}

	rep := new(repositories.TipoDocumentoRepository)
	id := rep.Insert(*entity)

	context.JSON(http.StatusCreated, id)
}

// Update godoc
// @Summary Actualiza un tipo de documento
// @Description Actualiza un tipo de documento dado un id y datos nuevos
// @Tags     tipo documento
// @Accept   json
// @Produce  json
// @Param        generoID   path      int  true  "Id"
// @Success 200 {int} cantidad de actualizaciones
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/models [put]
func (controller TipoDocumentoController) Update(context *gin.Context) {

	entity := new(models.TipoDocumento)

	ID := context.Param("tipoDocumentoID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	if err := context.BindJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error())) // NO MOSTRAR ERRORES DIRECTOS del SISTEMA. (NO SE HACE!!!!)
		return
	}

	rep := new(repositories.TipoDocumentoRepository)
	rowAffected := rep.Update(id, *entity)

	if rowAffected == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, rowAffected)
	}
}

// Delete godoc
// @Summary Borra un tipo de documento
// @Description Borra Lógicamente (no es una baja física) un tipo documento dado un id
// @Tags     tipo documento
// @Accept   json
// @Produce  json
// @Param        tipoDocumentoID   path      int  true  "Id"
// @Success 200 {object} null
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/models [delete]
func (controller TipoDocumentoController) Delete(context *gin.Context) {
	ID := context.Param("tipoDocumentoID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	entityRep := new(repositories.TipoDocumentoRepository)
	rowAffected := entityRep.Delete(id)

	if rowAffected == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, "")
	}
}
