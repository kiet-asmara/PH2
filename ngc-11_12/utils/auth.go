package utils

import (
	"net/http"
	"ngc-11/helpers"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")

		if authToken == "" {
			return c.JSON(http.StatusUnauthorized, "Auth token is empty")
		}

		if err := helpers.ValidateToken(authToken); err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid auth token")
		}

		return next(c)
	}
}
