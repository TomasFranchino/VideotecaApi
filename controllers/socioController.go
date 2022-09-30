package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"videotecaapi/dtos"
	"videotecaapi/helpers"
	"videotecaapi/models"
	"videotecaapi/repositories"
	"videotecaapi/utils"
)

type SocioController struct{}

// Get godoc
// @Summary      Recupera un socio
// @Description  Obtiene algo dado un id
// @Tags         socio
// @Accept       json
// @Produce      json
// @Param        socioID   path      int  true  "Id"
// @Success      200  {object}  models.Socio
// @Failure      400  {object}  models.Error
// @Failure      404  {object}  models.Error
// @Failure      500  {object}  models.Error
// @Router       /socio/{socioID} [get]
func (controller SocioController) Get(context *gin.Context) {
	ID := context.Param("socioID")

	id, err := strconv.Atoi(ID) // se convierte un string a int
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	entityRep := new(repositories.SocioRepository)
	entity := entityRep.Get(id)

	// Paso el registro devuelto a su correspondiente DTO
	entitySocioDTO := new(dtos.SelectSocioDTO)

	if entity.ID == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {

		entitySocioDTO.Apellido = entity.Apellido
		entitySocioDTO.CorreoElectronico = entity.CorreoElectronico
		entitySocioDTO.FechaNacimiento = entity.FechaNacimiento.Format("02-01-2006")
		entitySocioDTO.Nombre = entity.Nombre
		entitySocioDTO.NombreDocumento = entity.TipoDocumento.Nombre
		entitySocioDTO.TipoDocumento = entity.TipoDocumento.Tipo

		// context.JSON(http.StatusOK, entity)

		context.JSON(http.StatusOK, entitySocioDTO)

	}
}

// Find godoc
// @Summary Busca algún socio
// @Description Obtiene un conjunto de socios aplicando filtros varios
// @Tags     socio
// @Accept   json
// @Produce  json
// @Param   size     query    int     false        "Size"
// @Param   page     query    int     false        "Page"
// @Param   sort     query    string     false        "List fields"
// @Param   filter     query    string     false        "Filter"
// @Success 200 {object} models.Socio
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/models [get]
func (controller SocioController) Find(context *gin.Context) {

	entityRep := new(repositories.SocioRepository)
	entities := entityRep.Find(context)
	context.JSON(http.StatusOK, entities)
}

// Create godoc
// @Summary Crea un socio
// @Description dado un json con los datos de un socio, éste es persistido
// @Tags     socio
// @Accept   json
// @Produce  json
// @Success 200 {object} models.Socio
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/socio [post]
func (controller SocioController) Create(context *gin.Context) {

	entityDTO := new(dtos.NuevoSocioDTO)
	entity := new(models.Socio)

	// Se convierte el Json al objeto DTO
	if err := context.BindJSON(&entityDTO); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error()))
		return
	}

	// Se realiza los chequeos de los datos de entrada del socio
	listError := checkNuevosInputs(*entityDTO)
	if len(listError) != 0 {
		context.JSON(http.StatusBadRequest, utils.Error(listError))
		return
	}

	// Se buscar el nombre del Documento en la tabla Tipo de Documento, esto con dos fines:
	// 1 * Verificar que existe el tipo de documeneto indicado en la api.
	// 2 * Se debe obtener el ID del tipo de documento para guardarlo en la tabla Socio
	entityTipoDocumento := getDocumentType(entityDTO.NombreDocumento)
	if entityTipoDocumento.ID == 0 {
		context.JSON(http.StatusBadRequest, "Documento no encontrado")
		return
	}

	// Mapeo los datos del DTO (ingresado desde la api) a la entidad Socio que se usa para persistirla.
	entity.Apellido = entityDTO.Apellido
	entity.CorreoElectronico = entityDTO.CorreoElectronico
	entity.FechaNacimiento, _ = time.Parse("01/02/2006", entityDTO.FechaNacimiento)
	entity.Nombre = entityDTO.Nombre
	entity.TipoDocumento = entityTipoDocumento // asigno el objeto tipo documento a la propiedad de la entidad Socio

	// Se crea una instancia del repositorio Socio
	rep := new(repositories.SocioRepository)

	// Persiste el nuevo Socio
	id := rep.Insert(*entity)

	// Se retorna al cliente un Http Code Status = 201 (Creación)
	context.JSON(http.StatusCreated, id)
}

// Update godoc
// @Summary Actualiza un socio
// @Description Actualiza un socio dado un id y datos nuevos
// @Tags     socio
// @Accept   json
// @Produce  json
// @Param        socioID   path      int  true  "Id"
// @Success 200 {object} models.Socio
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/socio [put]
func (controller SocioController) Update(context *gin.Context) {

	entityDTO := new(dtos.ModificarSocioDTO)
	entityDB := new(models.Socio)

	// Se convierte el Json al objeto DTO
	if err := context.BindJSON(&entityDTO); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error())) // NO MOSTRAR ERRORES DIRECTOS del SISTEMA. (NO SE HACE!!!!)
		return
	}

	// Se recupera el parámetro de la api-recurso: localhost:8080/api/v1/socio/2 (socioID = 2)
	ID := context.Param("socioID")

	// Se convierte el string en intero.
	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	// Se realiza los chequeos de los datos de entrada del socio
	listError := checkModificacionInputs(*entityDTO)
	if len(listError) != 0 {
		context.JSON(http.StatusBadRequest, utils.Error(listError))
		return
	}

	// Se buscar el nombre del Documento en la tabla Tipo de Documento, esto con dos fines:
	// 1 * Verificar que existe el tipo de documeneto indicado en la api.
	// 2 * Se debe obtener el ID del tipo de documento para guardarlo en la tabla Socio
	entityTipoDocumento := getDocumentType(entityDTO.NombreDocumento)
	if entityTipoDocumento.ID == 0 {
		context.JSON(http.StatusBadRequest, "Documento no encontrado")
		return
	}

	// Se crea la intancia del repositorio de Socios
	rep := new(repositories.SocioRepository)

	// Se busca el Socio en la base de datos para determinar que exista.
	entityDB = rep.Get(id)
	// Se verifica que el Socio exista
	if entityDB == nil {
		context.JSON(http.StatusNotFound, "")
		return
	}

	// Mapeo los datos del DTO (ingresado desde la api) a la entidad Socio que se usa para persistirla.
	entityDB.Apellido = entityDTO.Apellido
	entityDB.CorreoElectronico = entityDTO.CorreoElectronico
	entityDB.FechaNacimiento, _ = time.Parse("01/02/2006", entityDTO.FechaNacimiento)
	entityDB.Nombre = entityDTO.Nombre
	entityDB.TipoDocumentoID = int(entityTipoDocumento.ID) // asigno el id del tipo documento al id del tipo de documento de la entidad a persistir

	// Actualizo el socio en la base de datos
	rowAffected := rep.Update(id, *entityDB)

	if rowAffected == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, rowAffected)
	}
}

// Delete godoc
// @Summary Borra un socio
// @Description Borra Lógicamente (no es una baja física) un genero dado un id
// @Tags     socio
// @Accept   json
// @Produce  json
// @Param        socioID   path      int  true  "Id"
// @Success 200 {object} models.socio
// @Failure 401 {object} models.Error
// @Failure 400 {object} models.Error
// @Router /api/socio [delete]
func (controller SocioController) Delete(context *gin.Context) {
	ID := context.Param("socioID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	// Se crea la instancia del repositorio Socio
	entityRep := new(repositories.SocioRepository)
	// Se elimina lógicamente el Socio de la base de datos
	rowAffected := entityRep.Delete(id)

	if rowAffected == 0 {
		context.JSON(http.StatusNotFound, nil)
	} else {
		context.JSON(http.StatusOK, nil)
	}
}

func checkNuevosInputs(entityDTO dtos.NuevoSocioDTO) (listError string) {
	listError = ""

	// Verifico que haya indicado el Nombre del Socio
	if len(strings.TrimSpace(entityDTO.Nombre)) == 0 {
		listError = "Debe indicar el nombre del Socio.\r\n"
	}

	// Verifico que haya indicado el Apellido del Socio
	if len(strings.TrimSpace(entityDTO.Apellido)) == 0 {
		listError += "Debe indicar el Apellido del Socio.\r\n"
	}

	// Verifico que haya indicado la Fecha de Nacimiento
	if len(strings.TrimSpace(entityDTO.FechaNacimiento)) == 0 {
		listError += "Debe indicar la fecha de Nacimiento.\r\n"
	}

	// Verifico que haya indicado una Fecha Válida
	if !helpers.IsDateValue(entityDTO.FechaNacimiento) {
		listError += "Debe indicar una fecha de Nacimiento válida.\r\n"
	}

	// Verifico que haya indicado el Nombre del Documento
	if len(strings.TrimSpace(entityDTO.NombreDocumento)) == 0 {
		listError += "Debe indicar el nomnbre del documento.\r\n"
	}

	return listError
}

func checkModificacionInputs(entityDTO dtos.ModificarSocioDTO) (listError string) {
	listError = ""

	// Verifico que haya indicado el Nombre del Socio
	if len(strings.TrimSpace(entityDTO.Nombre)) == 0 {
		listError = "Debe indicar el nombre del Socio.\r\n"
	}

	// Verifico que haya indicado el Apellido del Socio
	if len(strings.TrimSpace(entityDTO.Apellido)) == 0 {
		listError += "Debe indicar el Apellido del Socio.\r\n"
	}

	// Verifico que haya indicado la Fecha de Nacimiento
	if len(strings.TrimSpace(entityDTO.FechaNacimiento)) == 0 {
		listError += "Debe indicar la fecha de Nacimiento.\r\n"
	}

	// Verifico que haya indicado una Fecha Válida
	if !helpers.IsDateValue(entityDTO.FechaNacimiento) {
		listError += "Debe indicar una fecha de Nacimiento válida.\r\n"
	}

	// Verifico que haya indicado el Nombre del Documento
	if len(strings.TrimSpace(entityDTO.NombreDocumento)) == 0 {
		listError += "Debe indicar el nomnbre del documento.\r\n"
	}

	return listError
}

func getDocumentType(tipoDocumento string) (entityTipoDocumento models.TipoDocumento) {

	// Busco el nombre del Documento en la tabla Tipo de Documento, esto con dos fines:

	repTipoDocumento := new(repositories.TipoDocumentoRepository)

	entityTipoDocumento = *repTipoDocumento.GetByDocumentName(tipoDocumento)

	return entityTipoDocumento
}
