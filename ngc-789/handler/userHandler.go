package handler

import (
	"gin-ex/config"
	"gin-ex/entity"
	"gin-ex/helpers"
	"gin-ex/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// get register input
	var user entity.Store
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("error Register: %s \n", err.Error())
		utils.ErrorMessage(c, &utils.ErrBadRequest)
		return
	}

	// check if user (store) already exists
	var existingUser entity.Store

	config.DB.Where("store_email = ?", user.Store_email).First(&existingUser)
	if existingUser.Store_id != 0 {
		utils.ErrorMessage(c, &utils.ErrDataNotFound)
		return
	}

	// create hashed password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Panicln("error Register: ", err.Error())
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
		log.Printf("error Login: %s \n", err.Error())
		utils.ErrorMessage(c, &utils.ErrBadRequest)
		return
	}

	// check if user exists
	var existingUser entity.Store

	config.DB.Where("store_email = ?", user.Store_email).First(&existingUser)
	if existingUser.Store_id == 0 {
		utils.ErrorMessage(c, &utils.ErrDataNotFound)
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
		log.Panicln("error Register: ", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful. Welcome " + existingUser.Store_name + "!",
		"token":   token,
	})
}
