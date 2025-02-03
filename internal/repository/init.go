package repository

import (
	"database/sql"
	"fmt"
)

func ConnectDB() (*sql.DB, error) {
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
	return db, nil
}
