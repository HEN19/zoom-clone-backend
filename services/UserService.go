package services

import (
	"errors"
	database "self-improvement/zoom-clone-backend/DAO"
	"self-improvement/zoom-clone-backend/models"
)

func RegisterUser(data map[string]string) error {
	if data["email"] == "" || data["password"] == "" {
		return errors.New("email and password are required")
	}

	// Add user to the database
	return database.CreateUser(models.User{
		Email:    data["email"],
		Password: data["password"], // Hash this in production
	})
}

func LoginUser(data map[string]string) (string, error) {
	if data["email"] == "" || data["password"] == "" {
		return "", errors.New("email and password are required")
	}

	user, err := database.GetUserByEmail(data["email"])
	if err != nil || user.Password != data["password"] { // In production, compare hashed passwords
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token := "fake-jwt-token" // Replace with actual token generation
	return token, nil
}
