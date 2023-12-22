package main

import (
	"gin-ex/config"
	"gin-ex/handler"
	"gin-ex/middleware"
	"log"
	"os"

	_ "gin-ex/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// TODO:
// - add authorize store id w/product id (only store 1 can update/delete products in store 1)
// - add many to many in post
// - delete foreign key constraint

// @title Groot CMS
// @version 1.0
// @description E-commerce REST API
// @termsOfService http://swagger.io/terms/
// @contact.name Kiet Asmara
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apiKey JWT
// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed env:", err)
	}
	DBconfig := config.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL"),
	}

	err = config.InitDB(DBconfig)
	if err != nil {
		log.Fatal("failed db", err)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", handler.Register)
		userRoutes.POST("/login", handler.Login)
	}

	productRoutes := r.Group("/products")
	productRoutes.Use(middleware.AuthMiddleware(), middleware.ErrorMiddleware())
	{
		productRoutes.POST("", handler.AddProduct)
		productRoutes.GET("", handler.GetAllProducts)
		productRoutes.GET("/:id", handler.GetProductById) // show stores w/product
		productRoutes.PUT("/:id", handler.UpdateProduct)
		productRoutes.DELETE("/:id", handler.DeleteProduct)
	}

	r.Run()
}
