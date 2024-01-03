package handlers

import (
	"net/http"
	"ngc-11/config"
	"ngc-11/helpers"
	"ngc-11/model"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	// create hashed password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	user.Password = string(hashedPass)

	// register user
	result := config.DB.Omit("User_id").FirstOrCreate(&user, user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "register success",
		"user":    user,
	})
}

func Login(c echo.Context) error {
	// get user input
	var user model.UserLogin
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	// check if user exists
	var existingUser model.User

	config.DB.Where("username = ?", user.Username).First(&existingUser)
	if existingUser.User_id == 0 {
		return c.JSON(http.StatusUnauthorized, "Invalid username/passwords")
	}

	// check password
	passCheck := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if passCheck != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid username/passwords")
	}

	// generate JWT
	token, err := helpers.GenerateJWT(existingUser.User_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "login success",
		"token":   token,
	})
}
