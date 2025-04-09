package infraestructure

import (
	"actividad/src/perfumes/application"
	"actividad/src/perfumes/domain"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.IPerfume) *gin.Engine {
	r := gin.Default()

	// Crear casos de uso y controladores
	createPerfume := application.NewCreatePerfume(repo)
	createPerfumeController := NewCreatePerfumeController(createPerfume)

	viewPerfumes := application.NewViewPerfumes(repo)
	viewPerfumesController := NewViewPerfumesController(viewPerfumes)

	editPerfume := application.NewEditPerfume(repo)
	editPerfumeController := NewEditPerfumeController(editPerfume)

	deletePerfume := application.NewDeletePerfume(repo)
	deletePerfumeController := NewDeletePerfumeController(deletePerfume)

	// Rutas
	r.POST("/perfumes", createPerfumeController.Execute)
	r.GET("/perfumes", viewPerfumesController.Execute)
	r.PUT("/perfumes/:id", editPerfumeController.Execute)
	r.DELETE("/perfumes/:id", deletePerfumeController.Execute)

	return r
}