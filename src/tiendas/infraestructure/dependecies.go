package infraestructure

import (
	"log"
)

func Iniciar() {
	ps := NewMySQL()
	router := SetupRouter(ps)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
