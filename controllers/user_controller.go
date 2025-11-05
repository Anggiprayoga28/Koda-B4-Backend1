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

func (ctrl *UserController) GetUsers(c *gin.Context) {
	users := ctrl.userService.GetAllUsers()

	var userResponses []models.UserResponse
	for _, u := range users {
		userResponses = append(userResponses, models.ToUserResponse(u))
	}

	c.JSON(http.StatusOK, models.Response{
		Status: "success",
		Data:   userResponses,
	})
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
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
		Data:   models.ToUserResponse(*user),
	})
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	fullName := c.PostForm("full_name")

	user, passwordHash, err := ctrl.userService.CreateUser(username, email, password, fullName)
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
		Data:    models.ToUserResponseWithHash(*user, passwordHash),
	})
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
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

	if passwordHash != "" {
		c.JSON(http.StatusOK, models.Response{
			Status:  "success",
			Message: "User berhasil diupdate",
			Data:    models.ToUserResponseWithHash(*user, passwordHash),
		})
	} else {
		c.JSON(http.StatusOK, models.Response{
			Status:  "success",
			Message: "User berhasil diupdate",
			Data:    models.ToUserResponse(*user),
		})
	}
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
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
	})
}
