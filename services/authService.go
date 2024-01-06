package services

import (
	"errors"
	"fmt"
	"os"
	"safblog-backend/database"
	"safblog-backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
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
		return models.Response{Message: "failed to create user", Error: err}, errors.New(err)
	}

	hash, err := saltAndHash(registeredUser.Password)
	if err != nil {
		error := "error while hashing the password"
		return models.Response{Message: "failed to has password", Error: error}, errors.New(error)
	}

	user.Password = hash

	err = db.Create(&user).Error
	if err != nil {
		error := "could not create user"
		return models.Response{Message: "failed to create user", Error: error}, errors.New(error)
	}

	return models.Response{Message: "user created", Data: fiber.Map{"message": "user created."}}, nil
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
		return CreateResponse("failed to find user", nil, err), errors.New(err)
		//return models.Response{Message: "failed to find user", Error: err}, errors.New(err)
	}

	isPasswordValid := verifyPassword(dbUser.Password, []byte(loginUser.Password))

	if !isPasswordValid {
		err := "credentials are not valid"
		return models.Response{Message: "failed to authenticate user", Error: err}, errors.New(err)
	}

	claims := jwt.MapClaims{
		"id":    dbUser.ID,
		"email": dbUser.Email,
		"admin": false,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.

	t, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
	if err != nil {
		return models.Response{Message: "failed to create jwt signed token", Error: err.Error()}, errors.New(err.Error())
	}

	return models.Response{
		Message: "user login success",
		Data: fiber.Map{
			"token": t,
		},
	}, nil
}
