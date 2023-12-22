package entity

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

type Member struct {
	Id         int    `validate:"required" json:"id"`
	Email      string `validate:"required, email" json:"email"`
	Password   string `validate:"required, gte=8" json:"password"`
	Full_name  string `validate:"required, gte=6, lte=15" json:"full_name"`
	Age        int    `validate:"required, gte=17" json:"age"`
	Occupation string `validate:"required" json:"occupation"`
	Role       string `validate:"oneof=admin superadmin" json:"role"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
