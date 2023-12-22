package entity

type Store struct {
	Store_id    uint       `json:"store_id" gorm:"primarykey;autoIncrement"`
	Store_email string     `json:"store_email" binding:"required,email" gorm:"unique"`
	Password    string     `json:"password" binding:"required,gte=8"`
	Store_name  string     `json:"store_name" binding:"required,gte=6,lte=15"`
	Store_type  string     `json:"store_type" binding:"oneof=silver gold platinum"`
	Products    []*Product `gorm:"many2many:store_products;"`
}

type StoreLogin struct {
	Store_email string `json:"store_email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
}

type Product struct {
	Product_id  uint     `json:"product_id" gorm:"primarykey"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required,gte=50"`
	Image_url   string   `json:"image_url" binding:"required"`
	Price       int      `json:"price" binding:"required,gte=1000"`
	Stores      []*Store `gorm:"many2many:store_products;"`
}

type ProductUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image_url   string `json:"image_url"`
	Price       int    `json:"price"`
}
