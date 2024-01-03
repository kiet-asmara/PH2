package model

type Product struct {
	Product_id   int     `json:"product_id" gorm:"primary_key"`
	Product_name string  `json:"product_name"`
	Stock        int     `json:"stock"`
	Price        float32 `json:"price"`
}

type TransactionInput struct {
	Product_id int `json:"product_id" gorm:"primary_key"`
	Quantity   int `json:"quantity"`
}

type Transaction struct {
	User_id      int     `json:"user_id" gorm:"primary_key"`
	Product_id   int     `json:"product_id" gorm:"primary_key"`
	Quantity     int     `json:"quantity"`
	Total_amount float32 `json:"total_amount"`
}
