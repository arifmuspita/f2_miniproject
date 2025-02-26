package usecase

import (
	"f2_miniproject/model"
	"f2_miniproject/repository"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IUserUseCase interface {
	Register(student model.User) error
	Login(email, password string) (string, error)
}

type userUseCase struct {
	userRepo repository.IUserRepository
}

func NewUserUseCase(userRepo repository.IUserRepository) IUserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func isValidPassword(password string) bool {
	regex := `^.{8,}$`
	matched, _ := regexp.MatchString(regex, password)
	return matched
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(regex, email)
	return matched
}

func (u *userUseCase) Register(user model.User) error {
	if user.Email == "" || user.PasswordHash == "" || user.Name == "" {
		return ErrRegisterRequired
	}

	if !isValidEmail(user.Email) {
		return ErrInvalidEmailRegister
	}

	if !isValidPassword(user.PasswordHash) {
		return ErrInvalidPasswordRegister
	}

	err := u.userRepo.Register(&user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ErrDuplicatedKey
		}

		return ErrInternalServer
	}

	return nil
}

func (u *userUseCase) Login(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", ErrLoginRequired
	}

	user, err := u.userRepo.Login(email, password)
	if err != nil {
		if strings.Contains(err.Error(), "email doesn't exist") {
			return "nil", ErrLoginNotFound
		} else if strings.Contains(err.Error(), "wrong password") {
			return "nil", ErrLoginInvalidPassword
		}

		log.Println(err.Error())

		return "nil", ErrInternalServer
	}

	token, err := GenerateJWTToken(user.Email, user.Name, user.IsVerified)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateJWTToken(email string, name string, verified bool) (string, error) {
	JWTSecret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"name":        name,
		"email":       email,
		"is_verified": verified,
		"exp":         time.Now().Add(time.Minute * 60).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret))
}
