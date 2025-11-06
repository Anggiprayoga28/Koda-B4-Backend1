package models

type Profile struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	ProfilePic string `json:"profile_pic"`
}

var Profiles = []Profile{
	{ID: 1, UserID: 1, ProfilePic: ""},
	{ID: 2, UserID: 2, ProfilePic: ""},
	{ID: 3, UserID: 3, ProfilePic: ""},
}

var NextProfileID = 4
