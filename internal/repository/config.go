package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type Database struct {
	DB *sql.DB
}

// ConnectDB initializes the database connection and returns a Database struct
func ConnectDB() (*Database, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1234"
		dbname   = "rozgarlink"
	)

	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	// Check if the database connection is working
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	// log.Println("Database connected!")
	return &Database{DB: db}, nil
}

// CloseDB closes the database connection
func (d *Database) CloseDB() {
	if d.DB != nil {
		d.DB.Close()
		log.Println("Database connection closed")
	}
}
