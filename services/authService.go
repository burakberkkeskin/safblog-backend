package services

import (
	"errors"
	"fmt"
	"os"
	"safblog-backend/config"
	"safblog-backend/database"
	ResponseModels "safblog-backend/models/response_models"
	UserModels "safblog-backend/models/user_models"
	"strconv"
	"time"
	"unicode"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(registeredUser UserModels.RegisterModel) (ResponseModels.Response, error) {
	fmt.Printf("Register running username: %s, email: %s, password: %s\n", registeredUser.Username, registeredUser.Email, registeredUser.Password)

	// Check registeredUser.password and registeredUser.verifyPassword is same.
	if registeredUser.Password != registeredUser.PasswordVerify {
		err := "password and verify password is not same"
		return ResponseModels.Response{Message: "failed to create user", Error: err}, errors.New(err)
	}

	// Check the password meet the requirements.
	err := verifyPasswordFormat(registeredUser.Password)
	if err != nil {
		errorMessage := "password doesn't meet the requirements"
		response := ErrorResponse("failed to create user", errorMessage)
		return response, errors.New(errorMessage)
	}

	user := UserModels.User{
		Username: registeredUser.Username,
		Email:    registeredUser.Email,
		Password: registeredUser.Password,
	}

	db := database.DB.Db

	var dbEmailUser UserModels.User
	db.Find(&dbEmailUser, "email = ?", registeredUser.Email)
	if dbEmailUser.Email != "" {
		err := "email already in use"
		return ResponseModels.Response{Message: "failed to create user", Error: err}, errors.New(err)
	}

	db.Find(&dbEmailUser, "username = ?", registeredUser.Username)
	if dbEmailUser.Username != "" {
		err := "username already in use"
		return ResponseModels.Response{Message: "failed to create user", Error: err}, errors.New(err)
	}

	hash, err := saltAndHash(registeredUser.Password)
	if err != nil {
		error := "error while hashing the password"
		return ResponseModels.Response{Message: "failed to has password", Error: error}, errors.New(error)
	}
	user.Password = hash

	err = db.Create(&user).Error
	if err != nil {
		error := "could not create user"
		return ResponseModels.Response{Message: "failed to create user", Error: error}, errors.New(error)
	}

	return ResponseModels.Response{Message: "user created", Data: fiber.Map{"message": "user created."}}, nil
}

func saltAndHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}

func verifyPasswordFormat(s string) error {
	fmt.Println("Verifying password format.")
	letters := 0
	number := false
	upper := false
	special := false
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
			letters++
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			letters++
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			letters++
		}
	}
	eightOrMore := letters >= 8
	fmt.Println("Number: ", number, "Upper:", upper, "Special: ", special, "eightOrMore: ", letters)
	if number && upper && special && eightOrMore {
		return nil
	}
	return errors.New("password doesn't meet the requirements")
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

func LoginUser(loginUser UserModels.LoginUser) (ResponseModels.Response, error) {
	fmt.Printf("%s is logging in.\n", loginUser.Email)
	db := database.DB.Db

	var dbUser UserModels.User

	db.Find(&dbUser, "email = ?", loginUser.Email)

	if dbUser.ID == uuid.Nil {
		err := "user not found"
		return CreateResponse("failed to find user", nil, err), errors.New(err)
		//return models.Response{Message: "failed to find user", Error: err}, errors.New(err)
	}

	isPasswordValid := verifyPassword(dbUser.Password, []byte(loginUser.Password))

	if !isPasswordValid {
		err := "credentials are not valid"
		return ResponseModels.Response{Message: "failed to authenticate user", Error: err}, errors.New(err)
	}
	jwtHour, err := strconv.Atoi(os.Getenv("JWTHOUR"))
	if err != nil {
		err := "JWTHOUR env value is not integer"
		fmt.Println(err)
		return CreateResponse("internal server error", nil, ""), errors.New(err)
	}
	claims := jwt.MapClaims{
		"id":       dbUser.ID,
		"email":    dbUser.Email,
		"username": dbUser.Username,
		"role":     dbUser.Role,
		"exp":      time.Now().Add(time.Hour * time.Duration(jwtHour)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.

	t, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
	if err != nil {
		return ResponseModels.Response{Message: "failed to create jwt signed token", Error: err.Error()}, errors.New(err.Error())
	}

	return ResponseModels.Response{
		Message: "user login success",
		Data: fiber.Map{
			"token": t,
		},
	}, nil
}

func CreateRootUser() error {
	fmt.Println("Creating root user.")
	rootUserUsername := config.Config("ROOT_USER_USERNAME")
	rootUserEmail := config.Config("ROOT_USER_EMAIL")
	rootUserPassword := config.Config("ROOT_USER_PASSWORD")

	user := UserModels.User{
		Username: rootUserUsername,
		Email:    rootUserEmail,
		Password: rootUserPassword,
		Role:     "admin",
		IsRoot:   true,
	}

	db := database.DB.Db

	var dbRootUser UserModels.User
	db.Find(&dbRootUser, "is_root = ?", true)
	if dbRootUser.Email != "" {
		fmt.Println("Root user already exists. Skipping.")
		return nil
	}

	hash, err := saltAndHash(rootUserPassword)
	if err != nil {
		error := "error while hashing the root user password"
		return errors.New(error)
	}
	user.Password = hash

	err = db.Create(&user).Error
	if err != nil {
		error := "could not create the root user"
		return errors.New(error)
	}

	return nil
	//return services.SuccessResponse("admin user created")
}
