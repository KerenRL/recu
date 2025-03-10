package infraestructure

import (
	"actividad/src/perfumes/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditPerfumeController struct {
	useCase *application.EditPerfume
}

func NewEditPerfumeController(useCase *application.EditPerfume) *EditPerfumeController {
	return &EditPerfumeController{useCase: useCase}
}

func (ep_c *EditPerfumeController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de perfume inválido"})
		return
	}

	var body struct {
		Marca  string  `json:"marca"`
		Modelo string  `json:"modelo"`
		Precio float32 `json:"precio"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer los datos"})
		return
	}

	err = ep_c.useCase.Execute(int32(id), body.Marca, body.Modelo, body.Precio)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el perfume"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perfume actualizado correctamente"})
}
