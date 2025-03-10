package infraestructure

import (
	"actividad/src/tiendas/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewTiendaController struct {
	useCase *application.ViewTienda
}

func NewViewTiendaController(useCase *application.ViewTienda) *ViewTiendaController {
	return &ViewTiendaController{useCase: useCase}
}

func (vp_c *ViewTiendaController) Execute(c *gin.Context) {
	tiendas, err := vp_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tiendas)
}
