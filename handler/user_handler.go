package handler

import (
	"errors"
	"f2_miniproject/model"
	"f2_miniproject/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase usecase.IUserUseCase
}

func NewUserHandler(userUseCase usecase.IUserUseCase) UserHandler {
	return UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) Register(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "Invalid request body"},
		)
	}

	if err := h.userUseCase.Register(user); err != nil {

		var statusCode int
		var errorMessage string

		if errors.Is(err, usecase.ErrDuplicatedKey) {

			statusCode = http.StatusConflict
			errorMessage = err.Error()

		} else if errors.Is(err, usecase.ErrRegisterRequired) {

			statusCode = http.StatusBadRequest
			errorMessage = err.Error()

		} else if errors.Is(err, usecase.ErrInvalidEmailRegister) {

			statusCode = http.StatusBadRequest
			errorMessage = err.Error()

		} else if errors.Is(err, usecase.ErrInvalidPasswordRegister) {

			statusCode = http.StatusBadRequest
			errorMessage = err.Error()

		} else {

			statusCode = http.StatusInternalServerError
			errorMessage = err.Error()

		}

		return c.JSON(
			statusCode,
			map[string]string{"error": errorMessage},
		)
	}

	return c.JSON(
		http.StatusCreated,
		map[string]string{"message": "User register successful"},
	)
}

func (h *UserHandler) Login(c echo.Context) error {

	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "Invalid request body"},
		)
	}

	token, err := h.userUseCase.Login(data.Email, data.Password)

	if err != nil {

		var statusCode int
		var errorMessage string

		if errors.Is(err, usecase.ErrLoginRequired) {

			statusCode = http.StatusBadRequest
			errorMessage = err.Error()

		} else if errors.Is(err, usecase.ErrLoginNotFound) {

			statusCode = http.StatusNotFound
			errorMessage = err.Error()

		} else if errors.Is(err, usecase.ErrLoginInvalidPassword) {

			statusCode = http.StatusUnauthorized
			errorMessage = err.Error()

		} else {

			statusCode = http.StatusInternalServerError
			errorMessage = err.Error()

		}

		return c.JSON(
			statusCode,
			map[string]string{"error": errorMessage},
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"message": "Login successful",
			"token":   token,
		},
	)
}
