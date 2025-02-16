package app

import (
	"database/sql"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/auth"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type Dependencies struct {
	WorkerService worker.Service
	AuthService   auth.Service
}

func NewServices(db *sql.DB) Dependencies {
	WorkerRepo := repo.NewWorkerRepo(db)
	workerService := worker.NewService(WorkerRepo)

	AuthRepo := repo.NewAuthRepo(db)
	authService := auth.NewService(AuthRepo)
	return Dependencies{
		WorkerService: workerService,
		AuthService:   authService,
	}
}
