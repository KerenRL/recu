package infraestructure

import (
	"actividad/src/tiendas/application"
	"actividad/src/tiendas/domain"
	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.ITienda) *gin.Engine {
	r := gin.Default()

	createTienda := application.NewCreateTienda(repo)
	createTiendaController := NewCreateTiendaController(createTienda)

	viewTienda := application.NewViewTienda(repo)
	viewTiendaController := NewViewTiendaController(viewTienda)

	editTiendaUseCase := application.NewEditTienda(repo)
	editTiendaController := NewEditTiendaController(editTiendaUseCase)

	deleteTiendaUseCase := application.NewDeleteTienda(repo)
	deleteTiendaController := NewDeleteTiendaController(deleteTiendaUseCase)

	r.POST("/tiendas", createTiendaController.Execute)
	r.GET("/tiendas", viewTiendaController.Execute)
	r.PUT("/tiendas/:id", editTiendaController.Execute)
	r.DELETE("/tiendas/:id", deleteTiendaController.Execute)

	return r
}
