package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Price       float64 `json:"price" form:"price"`
	Stock       int     `json:"stock" form:"stock"`
}

var products = []Product{
	{ID: 1, Name: "Laptop", Description: "Laptop gaming", Price: 15000000, Stock: 10},
	{ID: 2, Name: "Mouse", Description: "Mouse wireless", Price: 250000, Stock: 50},
	{ID: 3, Name: "Keyboard", Description: "Mechanical keyboard", Price: 850000, Stock: 30},
}

func main() {
	r := gin.Default()
	r.GET("/products", getProducts)
	r.GET("/products/:id", getProductByID)

	r.Run(":8080")
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": products})
}

func getProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, p := range products {
		if p.ID == id {
			c.JSON(http.StatusOK, gin.H{"status": "success", "data": p})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Product not found"})
}
