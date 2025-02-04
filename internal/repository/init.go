package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DbData struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func InitDB() (*sql.DB, error) {

	godotenv.Load()

	dbData := DbData{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   os.Getenv("DB_NAME"),
	}

	psqlconn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbData.host, dbData.port, dbData.user, dbData.password, dbData.dbname,
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

	return db, nil
}
