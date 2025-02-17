package app

import (
	"database/sql"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/auth"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/employer"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/job"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type Dependencies struct {
	WorkerService   worker.Service
	AuthService     auth.Service
	EmployerService employer.Service
	JobService      job.Service
}

func NewServices(db *sql.DB) Dependencies {
	AuthRepo := repo.NewAuthRepo(db)
	WorkerRepo := repo.NewWorkerRepo(db)
	EmployerRepo := repo.NewEmployerRepo(db)
	JobRepo := repo.NewJobRepo(db)

	workerService := worker.NewService(WorkerRepo)
	authService := auth.NewService(AuthRepo)
	employerService := employer.NewService(EmployerRepo)
	jobService := job.NewService(JobRepo)

	return Dependencies{
		WorkerService:   workerService,
		AuthService:     authService,
		EmployerService: employerService,
		JobService:      jobService,
	}
}
