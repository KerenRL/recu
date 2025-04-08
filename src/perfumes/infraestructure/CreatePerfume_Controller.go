package infraestructure

import (
	"actividad/src/perfumes/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePerfumeController struct {
	useCase *application.CreatePerfume
}

func NewCreatePerfumeController(useCase *application.CreatePerfume) *CreatePerfumeController {
	return &CreatePerfumeController{useCase: useCase}
}

type RequestBody struct {
	Marca  string  `json:"marca"`
	Modelo string  `json:"modelo"`
	Precio float32 `json:"precio"`
}

func (cp_c *CreatePerfumeController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("Error al leer el JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	log.Printf("Datos recibidos: Marca: %s, Modelo: %s, Precio: %.2f", body.Marca, body.Modelo, body.Precio)

	err := cp_c.useCase.Execute(body.Marca, body.Modelo, body.Precio)
	if err != nil {
		log.Printf("Error al ejecutar el caso de uso: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al agregar el perfume", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Perfume agregado correctamente"})
}
