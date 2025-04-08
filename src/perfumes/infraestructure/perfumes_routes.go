package infraestructure

import (
	"actividad/src/perfumes/application"
	"actividad/src/perfumes/domain"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.IPerfume) *gin.Engine {
	r := gin.Default()

	encryptService := NewEncryptService()

	createPerfume := application.NewCreatePerfume(repo, encryptService)
	createPerfumeController := NewCreatePerfumeController(createPerfume)

	r.POST("/perfumes", createPerfumeController.Execute)

	return r
}
