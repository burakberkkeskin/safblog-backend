package services

import (
	"errors"
	"fmt"
	"safblog-backend/database"
	"safblog-backend/models"

	"github.com/google/uuid"
)

func CreateUser(registeredUser models.RegisterModel) models.Response {
	fmt.Printf("Register running username: %s, email: %s, password: %s\n", registeredUser.Username, registeredUser.Email, registeredUser.Password)

	user := models.User{
		Username: registeredUser.Username,
		Email:    registeredUser.Email,
		Password: registeredUser.Password,
	}
	db := database.DB.Db
	err := db.Create(&user).Error
	if err != nil {
		return models.Response{Status: "error", Error: "Could not create user", Data: err.Error()}
	}

	return models.Response{Status: "success", Data: "User created."}
}

func LoginUser(loginUser models.LoginUser) (models.Response, error) {
	fmt.Printf("%s is logging in.\n", loginUser.Email)
	db := database.DB.Db

	var dbUser models.User

	db.Find(&dbUser, "email = ?", loginUser.Email)

	if dbUser.ID == uuid.Nil {
		err := "user not found"
		return models.Response{Status: "error", Error: err}, errors.New(err)
	}
	if dbUser.Password != loginUser.Password {
		err := "credentials are not valid"
		return models.Response{Status: "error", Error: err}, errors.New(err)
	}
	fmt.Println(dbUser)

	return models.Response{Status: "success", Data: "{token: abcd}"}, nil
}
