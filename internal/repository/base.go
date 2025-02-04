package repository

import "database/sql"

type BaseRepository struct {
	DB *sql.DB
}
