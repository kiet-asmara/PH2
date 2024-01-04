package model

type TransactionInput struct {
	Product_id int `json:"product_id" gorm:"primary_key"`
	Quantity   int `json:"quantity"`
}

type Transaction struct {
	User_id      int     `json:"user_id" gorm:"primary_key"`
	Product_id   int     `json:"product_id" gorm:"primary_key"`
	Quantity     int     `json:"quantity"`
	Total_amount float32 `json:"total_amount"`
	Store_id     int     `json:"store_id" gorm:"foreignKey:Store_id"`
}
