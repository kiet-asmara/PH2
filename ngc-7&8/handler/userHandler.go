package handler

import (
	"fmt"
	"gin-ex/config"
	"gin-ex/entity"
	"gin-ex/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// get register input
	var user entity.Store
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input: " + err.Error()})
		c.Abort()
		return
	}

	// check if user (store) already exists
	var existingUser entity.Store

	config.DB.Where("store_email = ?", user.Store_email).First(&existingUser)
	if existingUser.Store_id != 0 {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	// create hashed password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		c.Abort()
		return
	}
	user.Password = string(hashedPass)

	// register user
	config.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "register successful!", "user": user})
}

func Login(c *gin.Context) {
	// get user input
	var user entity.StoreLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input" + err.Error()})
		return
	}

	// check if user exists
	var existingUser entity.Store

	config.DB.Where("store_email = ?", user.Store_email).First(&existingUser)
	if existingUser.Store_id == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	// check password
	passCheck := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if passCheck != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
		return
	}

	// generate JWT
	token, err := helpers.GenerateJWT(existingUser.Store_name, existingUser.Store_type, existingUser.Store_email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		c.Abort()
		return
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"user_id":  existingUser.ID,
	// 	"username": existingUser.Username,
	// })

	// jwtToken, err := token.SignedString(jwtSecret)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful. Welcome " + existingUser.Store_name + "!",
		"token":   token,
	})
}
