package app

import (
	"database/sql"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type Dependencies struct {
	WorkerService worker.Service
}

func NewServices(db *sql.DB) Dependencies {
	WorkerRepo := repo.NewWorkerRepo(db)
	workerService := worker.NewService(WorkerRepo)

	return Dependencies{
		WorkerService: workerService,
	}
}
