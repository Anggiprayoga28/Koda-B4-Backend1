package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

var users = []User{
	// {ID: 1, Username: "admin", Email: "admin@gmail.com", Password: "admin123", FullName: "Administrator"},
	// {ID: 2, Username: "anggi", Email: "anggi@gmail.com", Password: "anggi123", FullName: "Anggi"},
	// {ID: 3, Username: "prayoga", Email: "prayoga@gmail.com", Password: "Prayoga123", FullName: "Prayoga"},
}

var nextID = 3

func main() {
	r := gin.Default()

	r.POST("/auth/register", register)
	r.POST("/auth/login", login)

	r.GET("/users", getUsers)
	r.GET("/users/:id", getUserByID)
	r.POST("/users", createUser)
	r.PATCH("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.Run(":8080")
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func isValidUsername(username string) bool {
	return len(username) >= 3 && len(username) <= 20
}

func isValidPassword(password string) bool {
	return len(password) >= 6
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	fullName := c.PostForm("full_name")

	if username == "" || email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Username, email, dan password wajib diisi",
		})
		return
	}

	if !isValidUsername(username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Username harus 3-20 karakter",
		})
		return
	}

	if !isValidEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Format email tidak valid",
		})
		return
	}

	if !isValidPassword(password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Password minimal 6 karakter",
		})
		return
	}

	for _, u := range users {
		if u.Username == username {
			c.JSON(http.StatusConflict, gin.H{
				"status":  "error",
				"message": "Username sudah digunakan",
			})
			return
		}
		if u.Email == email {
			c.JSON(http.StatusConflict, gin.H{
				"status":  "error",
				"message": "Email sudah terdaftar",
			})
			return
		}
	}

	newUser := User{
		ID:       nextID,
		Username: username,
		Email:    email,
		Password: password,
		FullName: fullName,
	}
	nextID++
	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Registrasi berhasil",
		"data": gin.H{
			"id":        newUser.ID,
			"username":  newUser.Username,
			"email":     newUser.Email,
			"full_name": newUser.FullName,
		},
	})
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Username dan password wajib diisi",
		})
		return
	}

	for _, u := range users {
		if u.Username == username && u.Password == password {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Login berhasil",
				"data": gin.H{
					"id":        u.ID,
					"username":  u.Username,
					"email":     u.Email,
					"full_name": u.FullName,
				},
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  "error",
		"message": "Username atau password salah",
	})
}

func getUsers(c *gin.Context) {
	var safeUsers []gin.H
	for _, u := range users {
		safeUsers = append(safeUsers, gin.H{
			"id":        u.ID,
			"username":  u.Username,
			"email":     u.Email,
			"full_name": u.FullName,
		})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": safeUsers})
}

func getUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID tidak valid",
		})
		return
	}

	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"data": gin.H{
					"id":        u.ID,
					"username":  u.Username,
					"email":     u.Email,
					"full_name": u.FullName,
				},
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User tidak ditemukan"})
}

func createUser(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	fullName := c.PostForm("full_name")

	if username == "" || email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Username, email, dan password wajib diisi",
		})
		return
	}

	if !isValidUsername(username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Username harus 3-20 karakter",
		})
		return
	}

	if !isValidEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Format email tidak valid",
		})
		return
	}

	newUser := User{
		ID:       nextID,
		Username: username,
		Email:    email,
		Password: password,
		FullName: fullName,
	}
	nextID++
	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"id":        newUser.ID,
			"username":  newUser.Username,
			"email":     newUser.Email,
			"full_name": newUser.FullName,
		},
	})
}

func updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID tidak valid",
		})
		return
	}

	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	fullName := c.PostForm("full_name")

	for i, u := range users {
		if u.ID == id {
			if username != "" {
				if !isValidUsername(username) {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  "error",
						"message": "Username harus 3-20 karakter",
					})
					return
				}
				users[i].Username = username
			}
			if email != "" {
				if !isValidEmail(email) {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  "error",
						"message": "Format email tidak valid",
					})
					return
				}
				users[i].Email = email
			}
			if password != "" {
				if !isValidPassword(password) {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  "error",
						"message": "Password minimal 6 karakter",
					})
					return
				}
				users[i].Password = password
			}
			if fullName != "" {
				users[i].FullName = fullName
			}

			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"data": gin.H{
					"id":        users[i].ID,
					"username":  users[i].Username,
					"email":     users[i].Email,
					"full_name": users[i].FullName,
				},
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User tidak ditemukan"})
}

func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID tidak valid",
		})
		return
	}

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User berhasil dihapus"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User tidak ditemukan"})
}
