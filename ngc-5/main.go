package main

import (
	"fmt"
	"log"
	"net/http"
	"ngc-5/config"
	"ngc-5/routes"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.InitDb()
	router := httprouter.New()
	routes.InitRoutes(router)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	fmt.Println("starting server on PORT 8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
