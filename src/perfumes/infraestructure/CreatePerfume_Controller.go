package infraestructure

import (
	"actividad/src/perfumes/application"
	"actividad/src/perfumes/domain"
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

// Estructura para el cuerpo de la solicitud
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

	// Validar campos
	if body.Marca == "" || body.Modelo == "" || body.Precio <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: Marca, Modelo y Precio son obligatorios"})
		return
	}

	id, err := cp_c.useCase.Execute(body.Marca, body.Modelo, body.Precio)
	if err != nil {
		log.Printf("Error al ejecutar el caso de uso: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al agregar el perfume", "detalles": err.Error()})
		return
	}

	perfume := domain.Perfume{
		ID:     id,
		Marca:  body.Marca,
		Modelo: body.Modelo,
		Precio: body.Precio,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Perfume agregado correctamente",
		"id":      perfume.ID,
		"marca":   perfume.Marca,
		"modelo":  perfume.Modelo,
		"precio":  perfume.Precio,
	})
}