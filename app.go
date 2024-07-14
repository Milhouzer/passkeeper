package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	app_db "passkeeper/backend/db"
	"passkeeper/backend/models"

	"golang.org/x/crypto/bcrypt"
)

type Session struct {
	user *models.User
	// timeout int
}

type App struct {
	ctx     context.Context
	db      *sql.DB
	session *Session
}

func (a *App) GetUsername() string {
	if !a.IsLogged() {
		return "LoggedOut"
	}

	return a.session.user.Username
}

func (a *App) IsLogged() bool {
	return a.session != nil
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.db = initDB()
}

func (a *App) Login(username, password string) (bool, error) {
	user, err := app_db.GetUserByUsername(a.db, username)
	if err != nil {
		return false, err
	}
	if user == nil || !verifyPassword(user.PasswordHash, password) {
		return false, errors.New("invalid username or password")
	}

	return true, nil
}

func verifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (a *App) FetchPasswords() string {
	passwords, err := app_db.FetchPasswords(a.db)
	if err != nil {
		log.Println("Error fetching passwords:", err)
		return ""
	}

	jsonPasswords, err := json.Marshal(passwords)
	if err != nil {
		log.Println("Error marshalling passwords:", err)
		return ""
	}

	log.Println(string(jsonPasswords))
	return string(jsonPasswords)
}
