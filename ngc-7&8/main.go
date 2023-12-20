package main

import (
	"gin-ex/config"
	"gin-ex/handler"
	"gin-ex/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", handler.Register)
		userRoutes.POST("/login", handler.Login)
	}

	productRoutes := r.Group("/products")
	productRoutes.Use(middleware.AuthMiddleware())
	{
		productRoutes.POST("", handler.AddProduct)
		productRoutes.GET("", handler.GetAllProducts)
		productRoutes.GET("/:id", handler.GetProductById) // show stores w/product
		productRoutes.PUT("/:id", handler.UpdateProduct)
		productRoutes.DELETE("/:id", handler.DeleteProduct)
	}

	r.Run()
}
