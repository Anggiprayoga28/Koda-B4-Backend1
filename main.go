package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

var products = []Product{
	{ID: 1, Name: "Laptop", Description: "Laptop gaming high-end", Price: 15000000, Stock: 10},
	{ID: 2, Name: "Mouse", Description: "Mouse wireless", Price: 250000, Stock: 50},
	{ID: 3, Name: "Keyboard", Description: "Mechanical keyboard", Price: 850000, Stock: 30},
}

func main() {
	r := gin.Default()
	r.GET("/products", getProducts)
	r.GET("/products/:id", getProductByID)
	r.POST("/products", createProduct)
	r.PATCH("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)
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

func createProduct(c *gin.Context) {
	var p Product
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	p.ID = nextID
	nextID++
	products = append(products, p)
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": p})
}

var nextID = 4

func updateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var update Product
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	for i, p := range products {
		if p.ID == id {
			if update.Name != "" {
				products[i].Name = update.Name
			}
			if update.Description != "" {
				products[i].Description = update.Description
			}
			if update.Price > 0 {
				products[i].Price = update.Price
			}
			if update.Stock >= 0 {
				products[i].Stock = update.Stock
			}
			c.JSON(http.StatusOK, gin.H{"status": "success", "data": products[i]})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Product not found"})
}

func deleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Product deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Product not found"})
}
