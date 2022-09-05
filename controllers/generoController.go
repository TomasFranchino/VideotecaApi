package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"videotecaapi/models"
	"videotecaapi/repositories"
	"videotecaapi/utils"
)

type GeneroController struct{}

// Get godoc
// @Summary      Recupera un genero
// @Description  Obtiene algo dado un id
// @Tags         genero
// @Accept       json
// @Produce      json
// @Param        generoID   path      int  true  "Id"
// @Success      200  {object}  models.Genero
// @Failure      400  {object}  models.Error
// @Failure      404  {object}  models.Error
// @Failure      500  {object}  models.Error
// @Router       /genero/{generoID} [get]
func (controller GeneroController) Get(context *gin.Context) {
	ID := context.Param("generoID")

	id, err := strconv.Atoi(ID) // se convierte un string a int
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	entityRep := new(repositories.GeneroRepository)
	entity := entityRep.Get(id)

	if entity.ID == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, entity)
	}
}

// Find godoc
// @Summary Busca algún genero
// @Description Obtiene un conjunto de generos aplicando filtros varios
// @Tags     genero
// @Accept   json
// @Produce  json
// @Param   size     query    int     false        "Size"
// @Param   page     query    int     false        "Page"
// @Param   sort     query    string     false        "List fields"
// @Param   filter     query    string     false        "Filter"
// @Success 200 {object} models.Genero
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/models [get]
func (controller GeneroController) Find(context *gin.Context) {

	entityRep := new(repositories.GeneroRepository)
	entities := entityRep.Find(context)
	context.JSON(http.StatusOK, entities)
}

// Create godoc
// @Summary Crea un genero
// @Description dado un json con los datos de un genero, éste es persistido
// @Tags     genero
// @Accept   json
// @Produce  json
// @Success 200 {object} models.Model
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/models [post]
func (controller GeneroController) Create(context *gin.Context) {

	entity := new(models.Genero)

	if err := context.BindJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error()))
		return
	}

	rep := new(repositories.GeneroRepository)
	id := rep.Insert(*entity)

	context.JSON(http.StatusCreated, id)
}

// Update godoc
// @Summary Actualiza un genero
// @Description Actualiza un genero dado un id y datos nuevos
// @Tags     genero
// @Accept   json
// @Produce  json
// @Param        generoID   path      int  true  "Id"
// @Success 200 {object} models.Model
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/models [put]
func (controller GeneroController) Update(context *gin.Context) {

	entity := new(models.Genero)

	ID := context.Param("generoID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	if err := context.BindJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error())) // NO MOSTRAR ERRORES DIRECTOS del SISTEMA. (NO SE HACE!!!!)
		return
	}

	rep := new(repositories.GeneroRepository)
	rowAffected := rep.Update(id, *entity)

	if rowAffected == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, rowAffected)
	}
}

// Delete godoc
// @Summary Borra un genero
// @Description Borra Lógicamente (no es una baja física) un genero dado un id
// @Tags     genero
// @Accept   json
// @Produce  json
// @Param        generoID   path      int  true  "Id"
// @Success 200 {object} models.Model
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/models [delete]
func (controller GeneroController) Delete(context *gin.Context) {
	ID := context.Param("generoID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	entityRep := new(repositories.GeneroRepository)
	rowAffected := entityRep.Delete(id)

	if rowAffected == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, "")
	}
}
