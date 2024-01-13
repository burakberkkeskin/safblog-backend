package models

type RegisterModel struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordVerify string `json:"passwordVerify"`
}
