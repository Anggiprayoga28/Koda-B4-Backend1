package main

import (
	"backend1/config"
	"backend1/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	log.Printf("Starting server on port %s", cfg.Port)
	log.Printf("Show Password Hash: %v", cfg.ShowPasswordHash)

	r := gin.Default()

	routes.SetupRoutes(r)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
