package infraestructure

import (
	"actividad/src/tiendas/application"
	"actividad/src/tiendas/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTiendaController struct {
	useCase *application.CreateTienda
}

func NewCreateTiendaController(useCase *application.CreateTienda) *CreateTiendaController {
	return &CreateTiendaController{useCase: useCase}
}

type RequestBody struct {
	Nombre    string `json:"nombre"`
	Ubicacion string `json:"ubicacion"`
}

func (ctc *CreateTiendaController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("Error al leer el JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	if body.Nombre == "" || body.Ubicacion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre y Ubicación son obligatorios"})
		return
	}

	id, err := ctc.useCase.Execute(body.Nombre, body.Ubicacion)
	if err != nil {
		log.Printf("Error al ejecutar el caso de uso: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al agregar la tienda", "detalles": err.Error()})
		return
	}

	tienda := domain.Tienda{
		ID:        id,
		Nombre:    body.Nombre,
		Ubicacion: body.Ubicacion,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Tienda agregada correctamente",
		"id":        tienda.ID,
		"nombre":    tienda.Nombre,
		"ubicacion": tienda.Ubicacion,
	})
}
