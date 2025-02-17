package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/jmoiron/sqlx"
)

type EmployerStorer interface {
	FetchEmployerByID(ctx context.Context, employerId int) (EmployerResponse, error)
}

type employerStore struct {
	BaseRepository
}

func NewEmployerRepo(db *sql.DB) EmployerStorer {
	return &employerStore{
		BaseRepository: BaseRepository{db},
	}
}

func (es *employerStore) FetchEmployerByID(ctx context.Context, employerId int) (EmployerResponse, error) {
	db := sqlx.NewDb(es.DB, "postgres")

	var employer EmployerResponse
	query := `SELECT * from employers inner join address on employers.location = address.id where employers.id = $1;`

	err := db.Get(&employer, query, employerId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return EmployerResponse{}, apperrors.ErrNoEmployerExists
		}
		return EmployerResponse{}, err
	}

	return employer, nil
}
