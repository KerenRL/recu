package main

import (
	"actividad/src/perfumes/infraestructure"
	tiendaInfra "actividad/src/tiendas/infraestructure"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	if err := infraestructure.InitPerfume(); err != nil {
		log.Fatalf("Error al inicializar Perfume: %v", err)
	}
	if err := tiendaInfra.InitTienda(); err != nil {
		log.Fatalf("Error al inicializar Tienda: %v", err)
	}

	// Configurar rutas para perfumes
	perfumeRouter := infraestructure.SetupRouter(infraestructure.NewMySQL())
	for _, route := range perfumeRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	// Configurar rutas para tiendas
	tiendaRouter := tiendaInfra.SetupRouter(tiendaInfra.NewMySQL())
	for _, route := range tiendaRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	// Configurar proxies confiables (opcional)
	r.SetTrustedProxies([]string{"127.0.0.1"})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
