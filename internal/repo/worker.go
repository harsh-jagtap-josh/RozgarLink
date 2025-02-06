package repo

import (
	"context"
	"database/sql"
	// "github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
	// repo "github.com/harsh-jagtap-josh/RozgarLink/internal/repository"
)

const SelectWorkerByIDQuery = "SELECT id, name, contact_number, email, gender, sectors, skills, location, is_available, rating, total_jobs_worked, created_at, updated_at, language FROM workers WHERE id = $1;"
const UpdateWorkerByIDQuery = ""
const DeleteWorkerByIDQuery = ""

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

	row := ws.BaseRepository.DB.QueryRow(SelectWorkerByIDQuery, workerID)

	var worker Worker
	err := row.Scan(&worker.ID, &worker.Name, &worker.ContactNumber, &worker.Email, &worker.Gender, &worker.Sectors, &worker.Skills, &worker.Location, &worker.IsAvailable,
		&worker.Rating, &worker.TotalJobsWorked, &worker.CreatedAt, &worker.UpdatedAt, &worker.Language)

	if err != nil {
		return Worker{}, err
	}

	return worker, err
}
