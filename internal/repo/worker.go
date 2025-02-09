package repo

import (
	"context"
	"database/sql"
)

const selectWorkerByIDQuery = "SELECT id, name, contact_number, email, gender, sectors, skills, location, is_available, rating, total_jobs_worked, created_at, updated_at, language FROM workers WHERE id = $1;"

// const updateWorkerByIDQuery = ""
// const deleteWorkerByIDQuery = ""

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

	row := ws.BaseRepository.DB.QueryRow(selectWorkerByIDQuery, workerID)

	var worker Worker
	err := row.Scan(&worker.ID, &worker.Name, &worker.ContactNumber, &worker.Email, &worker.Gender, &worker.Sectors, &worker.Skills, &worker.Location, &worker.IsAvailable,
		&worker.Rating, &worker.TotalJobsWorked, &worker.CreatedAt, &worker.UpdatedAt, &worker.Language)

	if err != nil {
		return Worker{}, err
	}

	return worker, err
}
