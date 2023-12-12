package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"preview-conc/ngc-2/config"
)

type Hero struct {
	ID       int
	Name     string
	Universe string
	Skill    string
	ImageUrl string
}

type Villain struct {
	ID       int
	Name     string
	Universe string
	ImageUrl string
}

func HeroHandler(w http.ResponseWriter, r *http.Request) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect")
	}
	defer db.Close()

	ctx := context.Background()
	var heroes []Hero

	query := `SELECT * from heroes`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("Failed to connect")
	}

	for rows.Next() {
		h := Hero{}
		err := rows.Scan(&h.ID, &h.Name, &h.Universe, &h.Skill, &h.ImageUrl)
		if err != nil {
			log.Fatal(err)
		}
		heroes = append(heroes, h)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}

func VillainHandler(w http.ResponseWriter, r *http.Request) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect")
	}
	defer db.Close()

	ctx := context.Background()
	var villains []Villain

	query := `SELECT * from villain`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("Failed to connect")
	}

	for rows.Next() {
		v := Villain{}
		err := rows.Scan(&v.ID, &v.Name, &v.Universe, &v.ImageUrl)
		if err != nil {
			log.Fatal(err)
		}
		villains = append(villains, v)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(villains)
}

func main() {
	http.HandleFunc("/heroes", HeroHandler)
	http.HandleFunc("/villains", VillainHandler)

	fmt.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
