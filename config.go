package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type DbConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func InitDB(ctx context.Context) (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		logger.Errorw(ctx, "error occured in loading env variables while database connection", zap.Error(err))
		return nil, err
	}

	dbConfig := DbConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   os.Getenv("DB_NAME"),
	}

	psqlconn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.host, dbConfig.port, dbConfig.user, dbConfig.password, dbConfig.dbname,
	)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		logger.Errorw(ctx, "error occured initiating a database connection", zap.Error(err))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Errorw(ctx, "error occured while checking database connection status", zap.Error(err))
		db.Close()
		return nil, err
	}

	return db, nil
}
