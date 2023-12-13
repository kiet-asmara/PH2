package main

import (
	"fmt"
	"log"
	"net/http"
	"preview-conc/ngc-4/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/crimes", handler.GetCrimes)
	router.GET("/crimes/:id", handler.GetCrimesID)
	router.POST("/crimes", handler.PostCrime)
	router.PUT("/crimes/:id", handler.PutCrime)
	router.DELETE("/crimes/:id", handler.DeleteCrime)

	router.GET("/heroes", handler.GetHeroes)
	router.GET("/heroes/:id", handler.GetHeroesID)

	router.GET("/villains", handler.GetVillains)
	router.GET("/villains/:id", handler.GetVillainsID)

	fmt.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
