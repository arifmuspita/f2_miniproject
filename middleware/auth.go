package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CheckJWT(next echo.HandlerFunc) echo.HandlerFunc {

	JWTSecret := os.Getenv("JWT_SECRET")

	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "empty token"})
		}

		tokenStr := ""
		tokenArr := strings.Split(authHeader, " ")
		if len(tokenArr) == 2 && tokenArr[0] == "Bearer" {
			tokenStr = tokenArr[1]
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(JWTSecret), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "invalid token",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("email", claims["email"])
			c.Set("name", claims["name"])
			c.Set("verified", claims["verified"])
		}

		return next(c)
	}
}
