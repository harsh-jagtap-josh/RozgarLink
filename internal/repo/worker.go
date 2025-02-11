package repo

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

const fetchWorkerByIDQuery = "SELECT id, name, contact_number, email, gender, sectors, skills, location, is_available, rating, total_jobs_worked, created_at, updated_at, language FROM workers WHERE id = $1;"

type WorkerStorer interface {
	GetWorkerByID(ctx context.Context, workerID int) (Worker, error)
}

type workerStore struct {
	BaseRepository
}

func NewWorkerRepo(db *sql.DB) WorkerStorer {
	return &workerStore{
		BaseRepository: BaseRepository{db},
	}
}

func (ws *workerStore) GetWorkerByID(ctx context.Context, workerID int) (Worker, error) {

	db := sqlx.NewDb(ws.DB, "postgres")

	var worker Worker

	err := db.Get(&worker, fetchWorkerByIDQuery, workerID)
	if err != nil {
		return Worker{}, err
	}

	return worker, err
}
