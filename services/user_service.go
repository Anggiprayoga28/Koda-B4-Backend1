package services

import (
	"backend1/models"
	"errors"
	"strings"

	"github.com/matthewhartstonge/argon2"
)

type UserService struct {
	argon2Config argon2.Config
}

func NewUserService() *UserService {
	argon2Config := argon2.DefaultConfig()

	return &UserService{
		argon2Config: argon2Config,
	}
}

func (s *UserService) IsValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func (s *UserService) IsValidUsername(username string) bool {
	return len(username) >= 3 && len(username) <= 20
}

func (s *UserService) IsValidPassword(password string) bool {
	return len(password) >= 6
}

func (s *UserService) HashPassword(password string) (string, error) {
	encoded, err := s.argon2Config.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (s *UserService) VerifyPassword(hashedPassword, password string) (bool, error) {
	return argon2.VerifyEncoded([]byte(password), []byte(hashedPassword))
}

func (s *UserService) Register(username, email, password, fullName string) (*models.User, string, error) {
	if username == "" || email == "" || password == "" {
		return nil, "", errors.New("username, email, dan password wajib diisi")
	}

	if !s.IsValidUsername(username) {
		return nil, "", errors.New("username harus 3-20 karakter")
	}

	if !s.IsValidEmail(email) {
		return nil, "", errors.New("format email tidak valid")
	}

	if !s.IsValidPassword(password) {
		return nil, "", errors.New("password minimal 6 karakter")
	}

	for _, u := range models.Users {
		if u.Username == username {
			return nil, "", errors.New("username sudah digunakan")
		}
		if u.Email == email {
			return nil, "", errors.New("email sudah terdaftar")
		}
	}

	hashedPassword, err := s.HashPassword(password)
	if err != nil {
		return nil, "", errors.New("gagal meng-hash password")
	}

	newUser := models.User{
		ID:       models.NextID,
		Username: username,
		Email:    email,
		Password: hashedPassword,
		FullName: fullName,
	}
	models.NextID++
	models.Users = append(models.Users, newUser)

	return &newUser, hashedPassword, nil
}

func (s *UserService) Login(username, password string) (*models.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username dan password wajib diisi")
	}

	for _, u := range models.Users {
		if u.Username == username {
			valid, err := s.VerifyPassword(u.Password, password)
			if err != nil {
				return nil, errors.New("error verifikasi password")
			}
			if valid {
				return &u, nil
			}
		}
	}

	return nil, errors.New("username atau password salah")
}

func (s *UserService) GetAllUsers() []models.User {
	return models.Users
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	for _, u := range models.Users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("user tidak ditemukan")
}

func (s *UserService) CreateUser(username, email, password, fullName string) (*models.User, string, error) {
	if username == "" || email == "" || password == "" {
		return nil, "", errors.New("username, email, dan password wajib diisi")
	}

	if !s.IsValidUsername(username) {
		return nil, "", errors.New("username harus 3-20 karakter")
	}

	if !s.IsValidEmail(email) {
		return nil, "", errors.New("format email tidak valid")
	}

	if !s.IsValidPassword(password) {
		return nil, "", errors.New("password minimal 6 karakter")
	}

	hashedPassword, err := s.HashPassword(password)
	if err != nil {
		return nil, "", errors.New("gagal meng-hash password")
	}

	newUser := models.User{
		ID:       models.NextID,
		Username: username,
		Email:    email,
		Password: hashedPassword,
		FullName: fullName,
	}
	models.NextID++
	models.Users = append(models.Users, newUser)

	return &newUser, hashedPassword, nil
}

func (s *UserService) UpdateUser(id int, username, email, password, fullName string) (*models.User, string, error) {
	var newHashedPassword string

	for i, u := range models.Users {
		if u.ID == id {
			if username != "" {
				if !s.IsValidUsername(username) {
					return nil, "", errors.New("username harus 3-20 karakter")
				}
				models.Users[i].Username = username
			}
			if email != "" {
				if !s.IsValidEmail(email) {
					return nil, "", errors.New("format email tidak valid")
				}
				models.Users[i].Email = email
			}
			if password != "" {
				if !s.IsValidPassword(password) {
					return nil, "", errors.New("password minimal 6 karakter")
				}
				hashedPassword, err := s.HashPassword(password)
				if err != nil {
					return nil, "", errors.New("gagal meng-hash password")
				}
				models.Users[i].Password = hashedPassword
				newHashedPassword = hashedPassword
			}
			if fullName != "" {
				models.Users[i].FullName = fullName
			}
			return &models.Users[i], newHashedPassword, nil
		}
	}
	return nil, "", errors.New("user tidak ditemukan")
}

func (s *UserService) DeleteUser(id int) error {
	for i, u := range models.Users {
		if u.ID == id {
			models.Users = append(models.Users[:i], models.Users[i+1:]...)
			return nil
		}
	}
	return errors.New("user tidak ditemukan")
}
