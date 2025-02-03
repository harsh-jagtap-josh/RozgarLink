package app

import (
	"database/sql"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository/postgres"
)

type Dependencies struct {
	WorkerService worker.Service
}

func NewServices(db *sql.DB) Dependencies {
	WorkerRepo := postgres.NewWorkerRepo(db)
	workerService := worker.NewService(WorkerRepo)

	return Dependencies{
		WorkerService: workerService,
	}
}
