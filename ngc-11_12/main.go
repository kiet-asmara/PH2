package main

import (
	"ngc-11/config"
	"ngc-11/handlers"
	"ngc-11/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title Avengers Ecommerce API
// @version 1.0
// @description Buy products from stores.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	err = config.InitDB()
	if err != nil {
		e.Logger.Fatal("failed db", err)
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
