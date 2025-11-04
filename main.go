package main

import (
	"backend1/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
