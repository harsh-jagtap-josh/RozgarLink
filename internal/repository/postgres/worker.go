package postgres

import (
	"context"
	"database/sql"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository"
)

type workerStore struct {
	BaseRepository
}

func NewWorkerRepo(db *sql.DB) repository.WorkerStorer {
	return &workerStore{
		BaseRepository: BaseRepository{db},
	}
}

func (ps *workerStore) GetWorkerByID(ctx context.Context, tx repository.Transaction, workerID int64) (repository.Worker, error) {
	query := `SELECT * FROM Workers WHERE id = $1;`
	row := ps.DB.QueryRow(query, workerID)

	var worker repository.Worker
	err := row.Scan(&worker.ID, &worker.Name, &worker.ContactNumber, &worker.Email, &worker.Gender,
		&worker.Password, &worker.Sectors, &worker.Skills, &worker.Location, &worker.IsAvailable,
		&worker.Rating, &worker.TotalJobsWorked, &worker.CreatedAt, &worker.UpdatedAt, &worker.Language)

	return worker, err
}
