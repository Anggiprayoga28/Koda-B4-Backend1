package main

import (
	"backend1/config"
	_ "backend1/docs"
	"backend1/routes"
	"log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           User Management API
// @version         1.0
// @description     API sederhana untuk manajemen user dengan autentikasi menggunakan Argon2
// @termsOfService  http://swagger.io/terms/
// @host      localhost:8081
// @BasePath  /
// @schemes   http

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Server berjalan di http://localhost:%s", cfg.Port)
	log.Printf("Swagger UI tersedia di http://localhost:%s/swagger/index.html", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
