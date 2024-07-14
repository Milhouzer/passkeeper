package main

import (
	"database/sql"
	"embed"
	"log"
	"passkeeper/backend/models"

	_ "modernc.org/sqlite"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "passkeeper",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	createPasswordTableSQL := `CREATE TABLE IF NOT EXISTS passwords (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"url" TEXT UNIQUE NOT NULL,
		"password_hash" TEXT NOT NULL
	);`

	_, err = db.Exec(createPasswordTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createFakeData(db)

	return db
}

// Create fake passwords objects to add to the db
func createFakeData(db *sql.DB) {
	fakePasswords := []models.Password{
		{Url: "https://example.com", PasswordHash: "hash1"},
		{Url: "https://test.com", PasswordHash: "hash2"},
		{Url: "https://dummy.com", PasswordHash: "hash3"},
	}

	insertPasswordSQL := `INSERT INTO passwords (url, password_hash) VALUES (?, ?)`

	for _, pwd := range fakePasswords {
		_, err := db.Exec(insertPasswordSQL, pwd.Url, pwd.PasswordHash)
		if err != nil {
			log.Printf("Error inserting data for URL %s: %v\n", pwd.Url, err)
		} else {
			log.Printf("Inserted data for URL %s\n", pwd.Url)
		}
	}
}
