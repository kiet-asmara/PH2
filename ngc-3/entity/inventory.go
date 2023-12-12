package entity

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ItemCode    string `json:"itemcode"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
