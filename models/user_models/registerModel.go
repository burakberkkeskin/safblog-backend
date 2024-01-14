package models

type RegisterModel struct {
	Username       string `json:"username" validate:"required,min=3,max=16"`
	Email          string `json:"email" validate:"required,email,min=6,max=32"`
	Password       string `json:"password" validate:"required"`
	PasswordVerify string `json:"passwordVerify" validate:"required"`
}
