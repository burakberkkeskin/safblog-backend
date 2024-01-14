package models

type RegisterModel struct {
	Username       string `json:"username" validate:"required,min=3,max=16"`
	Email          string `json:"email" validate:"required,email,min=5,max=60"`
	Password       string `json:"password" validate:"required,min=8,max=32"`
	PasswordVerify string `json:"passwordVerify" validate:"required,min=8,max=32"`
}
