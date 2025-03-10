package infraestructure

import (
	"actividad/src/tiendas/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTiendaController struct {
	useCase *application.CreateTienda
}

func NewCreateTiendaController(useCase *application.CreateTienda) *CreateTiendaController {
	if useCase == nil {
		log.Println("Error: useCase es nil en CreateTiendaController")
	}
	return &CreateTiendaController{useCase: useCase}
}

type RequestBody struct {
	Nombre    string `json:"nombre"`
	Ubicacion string `json:"ubicacion"`
}

func (cp_c *CreateTiendaController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	if cp_c.useCase == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Dependencia useCase no inicializada"})
		return
	}

	err := cp_c.useCase.Execute(body.Nombre, body.Ubicacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al agregar la tienda", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tienda agregada correctamente"})
}
