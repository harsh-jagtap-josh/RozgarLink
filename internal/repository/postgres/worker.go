package postgres

import (
	"context"
	"database/sql"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/dto"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository"
)

type workerStore struct {
	repository.BaseRepository
}

func NewWorkerRepo(db *sql.DB) repository.WorkerStorer {
	return &workerStore{
		BaseRepository: repository.BaseRepository{db},
	}
}

func (ws *workerStore) GetWorkerByID(ctx context.Context, workerID int64) (dto.Worker, error) {
	query := `SELECT * FROM Workers WHERE id = $1;`
	row := ws.DB.QueryRow(query, workerID)

	var worker dto.Worker
	err := row.Scan(&worker.ID, &worker.Name, &worker.ContactNumber, &worker.Email, &worker.Gender,
		&worker.Password, &worker.Sectors, &worker.Skills, &worker.Location, &worker.IsAvailable,
		&worker.Rating, &worker.TotalJobsWorked, &worker.CreatedAt, &worker.UpdatedAt, &worker.Language)

	return worker, err
}
