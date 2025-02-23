package repo

import "github.com/jmoiron/sqlx"

type BaseRepository struct {
	DB *sqlx.DB
}
