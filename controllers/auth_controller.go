package controllers

import (
	"backend1/models"
	"backend1/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService *services.UserService
	jwtService  *services.JWTService
}

func NewAuthController() *AuthController {
	return &AuthController{
		userService: services.NewUserService(),
		jwtService:  services.NewJWTService(),
	}
}

// @Summary Register user baru
// @Description Mendaftarkan user baru dengan username, email, password, dan full name
// @Tags Authentication
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "Username (3-20 karakter)"
// @Param email formData string true "Email address"
// @Param password formData string true "Password (minimal 6 karakter)"
// @Param full_name formData string false "Nama lengkap"
// @Success 201 {object} models.Response "Registrasi berhasil"
// @Failure 400 {object} models.Response "Bad request"
// @Router /auth/register [post]
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

	token, _ := ctrl.jwtService.GenerateToken(user.ID, user.Username)

	c.JSON(http.StatusCreated, models.Response{
		Status:  "success",
		Message: "Registrasi berhasil",
		Data: gin.H{
			"user":  models.ToUserResponseWithHash(*user, passwordHash),
			"token": token,
		},
	})
}

// @Summary Login user
// @Description Login menggunakan username dan password
// @Tags Authentication
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} models.Response "Login berhasil"
// @Failure 400 {object} models.Response "Bad request"
// @Failure 401 {object} models.Response "Unauthorized - username atau password salah"
// @Router /auth/login [post]
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

	token, _ := ctrl.jwtService.GenerateToken(user.ID, user.Username)

	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Login berhasil",
		Data: gin.H{
			"user":  models.ToUserResponse(*user),
			"token": token,
		},
	})
}
