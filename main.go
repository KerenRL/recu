package main

import (
	"log"
	"time"

	"actividad/src/perfumes/infraestructure"
	tiendaInfra "actividad/src/tiendas/infraestructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	perfumeRepo := infraestructure.NewMySQL()
	if perfumeRepo == nil {
		log.Fatal("No se pudo inicializar el repositorio de perfumes")
	}

	tiendaRepo := tiendaInfra.NewMySQL()
	if tiendaRepo == nil {
		log.Fatal("No se pudo inicializar el repositorio de tiendas")
	}

	perfumeRouter := infraestructure.SetupRouter(perfumeRepo)
	for _, route := range perfumeRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	tiendaRouter := tiendaInfra.SetupRouter(tiendaRepo)
	for _, route := range tiendaRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("Error al configurar proxies: %v", err)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
