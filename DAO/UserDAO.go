package database

import (
	"database/sql"
	"errors"

	"self-improvement/zoom-clone-backend/config"
	"self-improvement/zoom-clone-backend/models"
)

// CreateUser inserts a new user into the database
func CreateUser(user models.User) error {
	query := `INSERT INTO users ( email, password) VALUES ( $1, $2)`
	_, err := config.DB.Exec(query, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(email string) (models.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = $1`
	row := config.DB.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return models.User{}, errors.New("user not found")
	}
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
