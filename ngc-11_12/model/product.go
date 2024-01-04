package model

type Product struct {
	Product_id   int     `json:"product_id" gorm:"primary_key"`
	Product_name string  `json:"product_name"`
	Stock        int     `json:"stock"`
	Price        float32 `json:"price"`
}
