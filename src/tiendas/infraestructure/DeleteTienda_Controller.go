package infraestructure

import (
	"actividad/src/tiendas/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteTiendaController struct {
	useCase *application.DeleteTienda
}

func NewDeleteTiendaController(useCase *application.DeleteTienda) *DeleteTiendaController {
	return &DeleteTiendaController{useCase: useCase}
}

func (dp_c *DeleteTiendaController) Execute(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de tienda inválido"})
		return
	}

	err = dp_c.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la tienda"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tienda eliminada correctamente"})
}
