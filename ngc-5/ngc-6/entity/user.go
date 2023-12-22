package entity

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
