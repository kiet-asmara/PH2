package main

import (
	"ngc-11/config"
	"ngc-11/handlers"
	"ngc-11/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := config.InitDB()
	if err != nil {
		e.Logger.Fatal("failed db", err)
	}

	err = godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}
	e.HTTPErrorHandler = utils.ErrorHandler

	e.Use(utils.MiddlewareLogging)

	user := e.Group("/users")
	user.POST("/register", handlers.Register)
	user.POST("/login", handlers.Login)

	e.GET("/products", handlers.GetProducts, utils.AuthMiddleware)
	e.GET("/stores", handlers.GetStores)
	e.GET("/stores/:id", handlers.GetStoreByID)
	e.POST("/transactions", handlers.BuyProduct, utils.AuthMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}
