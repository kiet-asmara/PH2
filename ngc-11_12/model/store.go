package model

type Store struct {
	Store_id      int     `json:"store_id"`
	Store_name    string  `json:"store_name"`
	Store_address string  `json:"store_address"`
	Longitude     string  `json:"longitude"`
	Latitude      string  `json:"latitude"`
	Rating        float32 `json:"rating"`
}
