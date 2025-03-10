package infraestructure

import (
	"actividad/src/tiendas/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditTiendaController struct {
	useCase *application.EditTienda
}

func NewEditTiendaController(useCase *application.EditTienda) *EditTiendaController {
	return &EditTiendaController{useCase: useCase}
}

func (ep_c *EditTiendaController) Execute(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de tienda inválido"})
		return
	}

	var body struct {
		Nombre    string `json:"nombre"`
		Ubicacion string `json:"ubicacion"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer los datos"})
		return
	}

	err = ep_c.useCase.Execute(int32(id), body.Nombre, body.Ubicacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la tienda"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tienda actualizada correctamente"})
}
