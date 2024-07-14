package app_db

import (
	"database/sql"
	"passkeeper/backend/models"
)

func FetchPasswords(db *sql.DB) ([]models.Password, error) {
	var passwords []models.Password

	rows, err := db.Query("SELECT id, url, password_hash FROM passwords")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var password models.Password
		err := rows.Scan(&password.ID, &password.Url, &password.PasswordHash)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, password)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return passwords, nil
}
