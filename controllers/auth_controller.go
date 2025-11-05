package controllers

import (
	"backend1/models"
	"backend1/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService *services.UserService
}

func NewAuthController() *AuthController {
	return &AuthController{
		userService: services.NewUserService(),
	}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	fullName := c.PostForm("full_name")

	user, passwordHash, err := ctrl.userService.Register(username, email, password, fullName)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Status:  "success",
		Message: "Registrasi berhasil",
		Data:    models.ToUserResponseWithHash(*user, passwordHash),
	})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := ctrl.userService.Login(username, password)
	if err != nil {
		statusCode := http.StatusBadRequest
		if err.Error() == "username atau password salah" {
			statusCode = http.StatusUnauthorized
		}
		c.JSON(statusCode, models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Login berhasil",
		Data:    models.ToUserResponse(*user),
	})
}
