package repository

import (
	"errors"
	"f2_miniproject/model"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Register(user *model.User) error
	Login(email, password string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (r *userRepository) Register(user *model.User) error {
	processHash, err := hashPassword(user.PasswordHash)
	if err != nil {
		log.Println("Failed to hash password ", err)
	}

	user.PasswordHash = processHash

	res := r.db.Create(&user)
	if res.Error != nil {
		log.Println(res.Error)
		return res.Error
	}

	return nil
}

func (r *userRepository) Login(email, password string) (*model.User, error) {
	var user model.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("email doesn't exist")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	return &user, nil
}
