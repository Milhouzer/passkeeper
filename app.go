package main

import (
	"context"
	"database/sql"
	"log"
	app_db "passkeeper/backend/db"
	"passkeeper/backend/models"
)

type Session struct {
	user *models.User
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
	return false
	return a.session != nil
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.db = initDB()
	a.loadUser()
}

func (a *App) loadUser() {
	a.session = nil
	_, err := app_db.FindFirst(a.db)
	if err != nil {
		log.Fatal(err)
	}

	// if existingUser != nil {
	// 	a.session = &Session{
	// 		user: existingUser,
	// 	}
	// 	return
	// }

	// err = app_db.CreateUser(a.db, "exampleUser2", "securePassword")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// existingUser, err = app_db.GetUserByUsername(a.db, "exampleUser")
	// if err != nil {
	// 	return
	// }

	// a.session = &Session{
	// 	user: existingUser,
	// }
}
