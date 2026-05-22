package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/shaurya21343/go-url-shortner/internal/config"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect(cfg *config.Config) {
	fmt.Println(cfg)
	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Verify the database file is accessible
	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("connected to db")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS urls(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT NOT NULL
	)`)
	if err != nil {
		db.Close() // Prevent connection leak
		log.Fatalf("Error creating table: %v", err)
	}

	DB = db
}

func CreateShortUrl(url string) (int64, error) {
	stmt, err := DB.Prepare("INSERT INTO urls (url) VALUES (?)")
	if err != nil {
		return 0, fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(url)
	if err != nil {
		return 0, fmt.Errorf("error executing statement: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID: %w", err)
	}
	return id, nil
}
