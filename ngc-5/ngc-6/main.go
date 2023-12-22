package main

import (
	"example/config"
	"example/routes"
	"fmt"
	"log"
	"net/http"

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

	fmt.Println("starting server on PORT 8080 uhuy...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
