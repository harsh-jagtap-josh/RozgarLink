package repo

import "database/sql"

type BaseRepository struct {
	DB *sql.DB
}
