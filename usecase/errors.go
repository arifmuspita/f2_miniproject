package usecase

import "errors"

var (
	ErrInternalServer          = errors.New("internal server error")
	ErrRegisterRequired        = errors.New("name, email and password are required")
	ErrInvalidEmailRegister    = errors.New("email is no valid")
	ErrInvalidPasswordRegister = errors.New("min length password is 8")
	ErrDuplicatedKey           = errors.New("email already exist")

	ErrLoginRequired        = errors.New("email and password are required")
	ErrLoginNotFound        = errors.New("email doesnt exist")
	ErrLoginInvalidPassword = errors.New("wrong password")
)
