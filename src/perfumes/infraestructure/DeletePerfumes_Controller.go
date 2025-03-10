package infraestructure

import (
	"actividad/src/perfumes/application"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeletePerfumeController struct {
	useCase *application.DeletePerfume
}

func NewDeletePerfumeController(useCase *application.DeletePerfume) *DeletePerfumeController {
	return &DeletePerfumeController{useCase: useCase}
}

func (dp_c *DeletePerfumeController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de perfume inválido"})
		return
	}

	err = dp_c.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al eliminar el perfume: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perfume eliminado correctamente"})
}
