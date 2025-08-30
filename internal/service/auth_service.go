package service

import (
	"errors"
	"time"

	"todo-api/internal/middleware"
	"todo-api/internal/model"
	"todo-api/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo  *repository.UserRepository
	JWTSecret []byte
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepo:  userRepo,
		JWTSecret: middleware.JWTSecretKey,
	}
}

func (s *AuthService) Register(username, password string) error {
	if username == "" {
		return errors.New("username is required")
	}
	if password == "" {
		return errors.New("password is required")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user, err := s.UserRepo.FindByUsername(username)
	if err == nil && user != nil {
		if user.Username == username {
			return errors.New("username already exists")
		}
	}

	newUser := model.User{Username: username, Password: string(hashed)}
	return s.UserRepo.Create(newUser)
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(s.JWTSecret)
}
