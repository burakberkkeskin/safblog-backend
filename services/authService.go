package services

import (
	"errors"
	"fmt"
	"safblog-backend/database"
	"safblog-backend/models"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(registeredUser models.RegisterModel) (models.Response, error) {
	fmt.Printf("Register running username: %s, email: %s, password: %s\n", registeredUser.Username, registeredUser.Email, registeredUser.Password)

	user := models.User{
		Username: registeredUser.Username,
		Email:    registeredUser.Email,
		Password: registeredUser.Password,
	}

	db := database.DB.Db

	var dbEmailUser models.User
	db.Find(&dbEmailUser, "email = ?", registeredUser.Email)
	if dbEmailUser.Email != "" {
		err := "email already in use"
		return models.Response{Status: "error", Error: err}, errors.New(err)
	}

	hash, err := saltAndHash(registeredUser.Password)
	if err != nil {
		error := "error while hashing the password"
		log.Errorf("%s", error)
		return models.Response{Status: "error", Error: error}, errors.New(error)
	}

	user.Password = hash

	err = db.Create(&user).Error
	if err != nil {
		error := "could not create user"
		return models.Response{Status: "error", Error: error, Data: err.Error()}, errors.New(error)
	}

	return models.Response{Status: "success", Data: "User created."}, nil
}

func saltAndHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}

func verifyPassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
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

	isPasswordValid := verifyPassword(dbUser.Password, []byte(loginUser.Password))

	if !isPasswordValid {
		err := "credentials are not valid"
		return models.Response{Status: "error", Error: err}, errors.New(err)
	}

	fmt.Println(dbUser)

	return models.Response{Status: "success", Data: "{token: abcd}"}, nil
}
