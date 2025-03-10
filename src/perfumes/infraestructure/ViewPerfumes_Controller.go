package infraestructure

import (
	"actividad/src/perfumes/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewPerfumesController struct {
	useCase *application.ViewPerfumes
}

func NewViewPerfumesController(useCase *application.ViewPerfumes) *ViewPerfumesController {
	return &ViewPerfumesController{useCase: useCase}
}

func (vp_c *ViewPerfumesController) Execute(c *gin.Context) {
	perfumes, err := vp_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, perfumes)
}
