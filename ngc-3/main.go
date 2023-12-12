package main

import (
	"fmt"
	"log"
	"net/http"
	"preview-conc/ngc-3/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/inventories", handler.InventoryGet)
	router.GET("/inventories/:id", handler.InventoryGetID)
	router.POST("/inventories", handler.InventoryPost)
	router.PUT("/inventories/:id", handler.InventoryPut)
	router.DELETE("/inventories/:id", handler.InventoryDelete)

	fmt.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
