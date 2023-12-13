package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"preview-conc/ngc-4/config"
	"preview-conc/ngc-4/entity"

	"github.com/julienschmidt/httprouter"
)

func GetHeroes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `SELECT * FROM Heroes`

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("Failed to connect")
	}

	var heroes []entity.Hero

	for rows.Next() {
		h := entity.Hero{}
		err := rows.Scan(&h.ID, &h.Name, &h.Universe, &h.Skill, &h.ImageUrl)
		if err != nil {
			log.Fatal(err)
		}
		heroes = append(heroes, h)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}

func GetHeroesID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := ps.ByName("id")

	query := `SELECT * FROM Heroes WHERE ID = ?`

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		log.Fatal(err)
	}

	var h entity.Hero

	for rows.Next() {
		err := rows.Scan(&h.ID, &h.Name, &h.Universe, &h.Skill, &h.ImageUrl)
		if err != nil {
			log.Fatal(err)
		}
	}

	if h.ID == 0 {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("hero with id doesn't exist")
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(h)
}
