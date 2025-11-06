package models

import "backend1/config"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type UserResponse struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	PasswordHash string `json:"password_hash,omitempty"`
}

type ProfileResponse struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	ProfilePic string `json:"profile_pic"`
	URL        string `json:"url,omitempty"`
}

func ToUserResponse(user User) UserResponse {
	response := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
	}

	if config.AppConfig != nil && config.AppConfig.ShowPasswordHash {
		response.PasswordHash = user.Password
	}

	return response
}

func ToUserResponseWithHash(user User, passwordHash string) UserResponse {
	response := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
	}

	if config.AppConfig != nil && config.AppConfig.ShowPasswordHash {
		if passwordHash != "" {
			response.PasswordHash = passwordHash
		} else {
			response.PasswordHash = user.Password
		}
	}

	return response
}

func ToProfileResponse(profile Profile) ProfileResponse {
	response := ProfileResponse{
		ID:         profile.ID,
		UserID:     profile.UserID,
		ProfilePic: profile.ProfilePic,
	}

	if profile.ProfilePic != "" {
		response.URL = "http://localhost:8081/" + profile.ProfilePic
	}

	return response
}
