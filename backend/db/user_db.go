package app_db

import (
	"database/sql"
	"passkeeper/backend/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *sql.DB, username, password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password_hash, token) VALUES (?, ?, ?)", username, hashedPassword, "")
	return err
}

func FindFirst(db *sql.DB) (*models.User, error) {
	var user models.User

	row := db.QueryRow("SELECT id, username, password_hash, token FROM users ORDER BY id ASC LIMIT 1")

	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	var user models.User

	row := db.QueryRow("SELECT id, username, password_hash, token FROM users WHERE username = ?", username)

	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
