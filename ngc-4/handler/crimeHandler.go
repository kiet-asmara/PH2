package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"preview-conc/ngc-4/config"
	"preview-conc/ngc-4/entity"

	"github.com/julienschmidt/httprouter"
)

func GetCrimes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `SELECT * FROM Crimes`

	ctx := context.Background()
	var crimes []entity.Crime

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("Failed to connect")
	}

	for rows.Next() {
		c := entity.Crime{}
		err := rows.Scan(&c.CrimeID, &c.HeroID, &c.VillainID, &c.Description, &c.CrimeTime)
		if err != nil {
			log.Fatal(err)
		}
		crimes = append(crimes, c)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(crimes)
}

func GetCrimesID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ID := ps.ByName("id")

	crime := SearchCrimeID(db, ID)

	if crime.CrimeID == 0 {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("crime with id doesn't exist")
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(crime)
}

func PostCrime(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var crime entity.Crime

	err = decoder.Decode(&crime)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{Message: "Invalid Body Input"})
		return
	}

	query := `INSERT INTO Crimes (heroID, VillainID, Description, CrimeTime)
	VALUES (?, ?, ?, ?)`

	ctx := context.Background()

	_, err = db.ExecContext(ctx, query, crime.HeroID, crime.VillainID, crime.Description, crime.CrimeTime)
	if err != nil {
		log.Fatal(err)
	}

	resp := map[string]any{
		"message": "success",
		"sent":    crime,
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func PutCrime(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ID := ps.ByName("id")

	// check if crime exists
	crimeOld := SearchCrimeID(db, ID)

	if crimeOld.CrimeID == 0 {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("crime with id doesn't exist")
		return
	}

	// get new values
	decoder := json.NewDecoder(r.Body)
	var crimeNew entity.Crime

	err = decoder.Decode(&crimeNew)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{Message: "Invalid Body Input"})
		return
	}

	query := `UPDATE Crimes SET HeroID=?, VillainID=?, Description=?, CrimeTime=? WHERE ID = ?`

	ctx := context.Background()

	_, err = db.ExecContext(ctx, query, crimeNew.HeroID, crimeNew.VillainID, crimeNew.Description, crimeNew.CrimeTime, ID)
	if err != nil {
		log.Fatal(err)
	}

	resp := map[string]any{
		"message": "success",
		"sent":    crimeNew,
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func DeleteCrime(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ID := ps.ByName("id")

	query := `DELETE FROM Crimes WHERE ID = ?`

	ctx := context.Background()

	_, err = db.ExecContext(ctx, query, ID)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("success")
}

func SearchCrimeID(db *sql.DB, id string) entity.Crime {
	query := `SELECT * FROM Crimes WHERE ID = ?`

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		log.Fatal(err)
	}

	var c entity.Crime

	for rows.Next() {
		err := rows.Scan(&c.CrimeID, &c.HeroID, &c.VillainID, &c.Description, &c.CrimeTime)
		if err != nil {
			log.Fatal(err)
		}
	}

	return c
}
