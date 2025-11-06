package controllers

import (
	"backend1/models"
	"backend1/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	service *services.ProfileService
}

func NewProfileController() *ProfileController {
	return &ProfileController{
		service: services.NewProfileService(),
	}
}

// UploadPicture - Upload gambar profile
// @Summary Upload gambar profile
// @Tags Profile
// @Accept multipart/form-data
// @Produce json
// @Param user_id path int true "User ID"
// @Param profile_pic formData file true "Gambar"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /upload/{user_id} [post]
func (c *ProfileController) UploadPicture(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("user_id"))

	file, err := ctx.FormFile("profile_pic")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "File tidak ditemukan",
		})
		return
	}

	profile, err := c.service.UploadPicture(userID, file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Upload berhasil",
		Data:    models.ToProfileResponse(*profile),
	})
}

// GetPicture - Lihat gambar profile
// @Summary Lihat gambar profile
// @Tags Profile
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /upload/{user_id} [get]
func (c *ProfileController) GetPicture(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("user_id"))

	profile, err := c.service.GetPicture(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Status: "success",
		Data:   models.ToProfileResponse(*profile),
	})
}
