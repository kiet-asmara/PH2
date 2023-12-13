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

func GetVillains(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `SELECT * FROM Villain`

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("Failed to connect")
	}

	var villains []entity.Villain

	for rows.Next() {
		v := entity.Villain{}
		err := rows.Scan(&v.ID, &v.Name, &v.Universe, &v.ImageUrl)
		if err != nil {
			log.Fatal(err)
		}
		villains = append(villains, v)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(villains)
}

func GetVillainsID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := ps.ByName("id")

	query := `SELECT * FROM Villain WHERE ID = ?`

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		log.Fatal(err)
	}

	var v entity.Villain

	for rows.Next() {
		err := rows.Scan(&v.ID, &v.Name, &v.Universe, &v.ImageUrl)
		if err != nil {
			log.Fatal(err)
		}
	}

	if v.ID == 0 {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("hero with id doesn't exist")
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(v)
}
