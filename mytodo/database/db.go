package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Initialize the database connection and create tasks table if it doesn't exist
func InitDB() {
	//Create or open the sqlite database file
	db, err := sql.Open("sqlite3", "mytodo.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Ensure the database file is closed when the application exists
	DB = db

	// Create the tasks table if it doesn't exist
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			description TEXT NOT NULL,
			due_date TEXT NOT NULL,
			completed BOOLEAN DEFAULT 0
		);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create tasks table: %v", err)
	}
	fmt.Println("Database initialize successfully!")
}

// close the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Database connection closed.")
	}
}
