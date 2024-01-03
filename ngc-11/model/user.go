package model

// for query
type User struct {
	User_id        int     `json:"id" gorm:"primary_key"`
	Username       string  `json:"username" binding:"required" gorm:"unique"`
	Password       string  `json:"password" binding:"required"`
	Deposit_amount float32 `json:"deposit_amount"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}
