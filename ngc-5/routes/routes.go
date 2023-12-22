package routes

import (
	"ngc-5/config"
	"ngc-5/handler"
	"ngc-5/middleware"

	"github.com/julienschmidt/httprouter"
)

func InitRoutes(r *httprouter.Router) {
	db := config.InitDb()
	handler := handler.New(db)
	middleware := middleware.NewAuth(db)

	r.POST("/register", handler.UserRegister)
	r.POST("/login", handler.UserLogin)

	r.GET("/protected", middleware.Authentication(handler.ProtectedEndpoint))
	r.GET("/public", handler.PublicEndpoint)
}
