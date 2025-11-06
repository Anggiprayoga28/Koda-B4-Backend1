package controllers

import (
	"backend1/models"
	"backend1/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// GetUsers godoc
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} models.Response
// @Failure 401 {object} models.Response
// @Router /users [get]
func (ctrl *UserController) GetUsers(c *gin.Context) {
	userID, _ := c.Get("user_id")
	username, _ := c.Get("username")

	users := ctrl.userService.GetAllUsers()

	var userResponses []models.UserResponse
	for _, u := range users {
		userResponses = append(userResponses, models.ToUserResponse(u))
	}

	c.JSON(http.StatusOK, models.Response{
		Status: "success",
		Data: gin.H{
			"authorized_user": gin.H{
				"user_id":  userID,
				"username": username,
			},
			"users": userResponses,
		},
	})
}

// GetUserByID godoc
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /users/{id} [get]
func (ctrl *UserController) GetUserByID(c *gin.Context) {
	userID, _ := c.Get("user_id")
	username, _ := c.Get("username")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "ID tidak valid",
		})
		return
	}

	user, err := ctrl.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status: "success",
		Data: gin.H{
			"authorized_user": gin.H{
				"user_id":  userID,
				"username": username,
			},
			"user": models.ToUserResponse(*user),
		},
	})
}

// CreateUser godoc
// @Tags Users
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "Username"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param full_name formData string false "Full Name"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Router /users [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	userID, _ := c.Get("user_id")
	username, _ := c.Get("username")

	usernameForm := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	fullName := c.PostForm("full_name")

	user, passwordHash, err := ctrl.userService.CreateUser(usernameForm, email, password, fullName)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Status:  "success",
		Message: "User berhasil dibuat",
		Data: gin.H{
			"created_by": gin.H{
				"user_id":  userID,
				"username": username,
			},
			"user": models.ToUserResponseWithHash(*user, passwordHash),
		},
	})
}

// UpdateUser godoc
// @Tags Users
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "User ID"
// @Param username formData string false "Username"
// @Param email formData string false "Email"
// @Param password formData string false "Password"
// @Param full_name formData string false "Full Name"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /users/{id} [patch]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	userID, _ := c.Get("user_id")
	usernameToken, _ := c.Get("username")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "ID tidak valid",
		})
		return
	}

	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	fullName := c.PostForm("full_name")

	user, passwordHash, err := ctrl.userService.UpdateUser(id, username, email, password, fullName)
	if err != nil {
		statusCode := http.StatusBadRequest
		if err.Error() == "user tidak ditemukan" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	responseData := gin.H{
		"updated_by": gin.H{
			"user_id":  userID,
			"username": usernameToken,
		},
	}

	if passwordHash != "" {
		responseData["user"] = models.ToUserResponseWithHash(*user, passwordHash)
	} else {
		responseData["user"] = models.ToUserResponse(*user)
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "User berhasil diupdate",
		Data:    responseData,
	})
}

// DeleteUser godoc
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	userID, _ := c.Get("user_id")
	username, _ := c.Get("username")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "ID tidak valid",
		})
		return
	}

	err = ctrl.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "User berhasil dihapus",
		Data: gin.H{
			"deleted_by": gin.H{
				"user_id":  userID,
				"username": username,
			},
			"deleted_user_id": id,
		},
	})
}
