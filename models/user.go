package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

var Users = []User{
	{ID: 1, Username: "admin", Email: "admin@gmail.com", Password: "admin123", FullName: "Administrator"},
	{ID: 2, Username: "anggi", Email: "anggi@gmail.com", Password: "anggi123", FullName: "Anggi"},
	{ID: 3, Username: "prayoga", Email: "prayoga@gmail.com", Password: "Prayoga123", FullName: "Prayoga"},
}

var NextID = 4
