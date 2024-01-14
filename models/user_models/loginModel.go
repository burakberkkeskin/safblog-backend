package models

type LoginUser struct {
	Email    string `json:"email" validate:"required,email,min=5,max=60"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
