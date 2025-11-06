package services

import (
	"backend1/models"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ProfileService struct{}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}

func (s *ProfileService) UploadPicture(userID int, file *multipart.FileHeader) (*models.Profile, error) {
	var profileIndex = -1
	for i := range models.Profiles {
		if models.Profiles[i].UserID == userID {
			profileIndex = i
			break
		}
	}

	if profileIndex == -1 {
		newProfile := models.Profile{
			ID:         models.NextProfileID,
			UserID:     userID,
			ProfilePic: "",
		}
		models.NextProfileID++
		models.Profiles = append(models.Profiles, newProfile)
		profileIndex = len(models.Profiles) - 1
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
		return nil, errors.New("format tidak didukung")
	}

	if file.Size > 5*1024*1024 {
		return nil, errors.New("file terlalu besar")
	}

	os.MkdirAll("./uploads", os.ModePerm)

	if models.Profiles[profileIndex].ProfilePic != "" {
		os.Remove(models.Profiles[profileIndex].ProfilePic)
	}

	filename := fmt.Sprintf("profile_%d_%d%s", userID, time.Now().Unix(), ext)
	filePath := filepath.Join("uploads", filename)

	src, err := file.Open()
	if err != nil {
		return nil, errors.New("gagal membuka file")
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, errors.New("gagal menyimpan file")
	}
	defer dst.Close()

	io.Copy(dst, src)

	models.Profiles[profileIndex].ProfilePic = filePath

	return &models.Profiles[profileIndex], nil
}

func (s *ProfileService) GetPicture(userID int) (*models.Profile, error) {
	for i := range models.Profiles {
		if models.Profiles[i].UserID == userID {
			if models.Profiles[i].ProfilePic == "" {
				return nil, errors.New("belum ada foto profil")
			}
			return &models.Profiles[i], nil
		}
	}
	return nil, errors.New("profil tidak ditemukan")
}
