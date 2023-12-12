package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"preview-conc/ngc-3/config"
	"preview-conc/ngc-3/entity"

	"github.com/julienschmidt/httprouter"
)

// GET /inventories
func InventoryGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `SELECT * FROM Inventories`

	ctx := context.Background()
	var items []entity.Item

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("Failed to connect")
	}

	for rows.Next() {
		itemOld := entity.Item{}
		err := rows.Scan(&itemOld.ID, &itemOld.Name, &itemOld.ItemCode, &itemOld.Stock, &itemOld.Description, &itemOld.Status)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, itemOld)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GET /inventories/:id
func InventoryGetID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ID := ps.ByName("id")

	query := `SELECT * FROM Inventories WHERE ID = ?`

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, query, ID)
	if err != nil {
		log.Fatal(err)
	}

	var items []entity.Item

	for rows.Next() {
		itemOld := entity.Item{}
		err := rows.Scan(&itemOld.ID, &itemOld.Name, &itemOld.ItemCode, &itemOld.Stock, &itemOld.Description, &itemOld.Status)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, itemOld)
	}

	if items == nil {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode("item with id doesn't exist")
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// POST /inventories
func InventoryPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var item entity.Item

	err = decoder.Decode(&item)
	if err != nil {
		log.Fatal(err)
	}

	query := `INSERT INTO Inventories (Name, ItemCode, Stock, Description, Status)
	VALUES (?, ?, ?, ?, ?)`

	ctx := context.Background()

	_, err = db.ExecContext(ctx, query, item.Name, item.ItemCode, item.Stock, item.Description, item.Status)
	if err != nil {
		log.Fatal(err)
	}

	var res entity.Response
	res.Message = "Success post"

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
	json.NewEncoder(w).Encode(item)
}

func InventoryPut(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ID := ps.ByName("id")

	decoder := json.NewDecoder(r.Body)
	var item entity.Item

	err = decoder.Decode(&item)
	if err != nil {
		log.Fatal(err)
	}

	query := `UPDATE Inventories SET Name=?, ItemCode=?, Stock=?, Description=?, Status=? WHERE ID = ?`

	ctx := context.Background()

	_, err = db.ExecContext(ctx, query, item.Name, item.ItemCode, item.Stock, item.Description, item.Status, ID)
	if err != nil {
		log.Fatal(err)
	}

	var res entity.Response
	res.Message = "Success update"

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
	json.NewEncoder(w).Encode(item)
}

func InventoryDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ID := ps.ByName("id")

	query := `DELETE FROM Inventories WHERE ID = ?`

	ctx := context.Background()

	_, err = db.ExecContext(ctx, query, ID)
	if err != nil {
		log.Fatal(err)
	}

	var res entity.Response
	res.Message = "Success delete"

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}
