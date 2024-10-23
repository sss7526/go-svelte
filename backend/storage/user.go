package storage

import (
	// "go-svelte/storage"
	"log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username 	string
	Password 	string
	Role		string
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

func CreateUser(username, plainPassword string) error {
	hashedPassword, err := HashPassword(plainPassword)
	if err != nil {
		return err
	}

	err = InsertUser(username, hashedPassword, "user")
	if err != nil {
		return err
	}

	log.Printf("User created: %s", username)
	return nil 
}

func AuthenticateUser(username, plainPassword string) error {
	user, err := FindUserByUsername(username)
	if err != nil {
		return err
	}
	return VerifyPassword(user.Password, plainPassword)
}