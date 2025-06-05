package app

import (
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/admin"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/application"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/auth"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/employer"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/job"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/sector"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
	"github.com/jmoiron/sqlx"
)

type Dependencies struct {
	WorkerService      worker.Service
	AuthService        auth.Service
	EmployerService    employer.Service
	JobService         job.Service
	ApplicationService application.Service
	SectorService      sector.Service
	AdminService       admin.AdminService
}

func NewServices(db *sqlx.DB) Dependencies {
	AuthRepo := repo.NewAuthRepo(db)
	WorkerRepo := repo.NewWorkerRepo(db)
	EmployerRepo := repo.NewEmployerRepo(db)
	JobRepo := repo.NewJobRepo(db)
	ApplicationRepo := repo.NewApplicationRepo(db)
	SectorRepo := repo.NewSectorRepo(db)
	AdminRepo := repo.NewAdminRepo(db)

	workerService := worker.NewService(WorkerRepo)
	authService := auth.NewService(AuthRepo)
	employerService := employer.NewService(EmployerRepo)
	jobService := job.NewService(JobRepo)
	applicationService := application.NewService(ApplicationRepo)
	sectorService := sector.NewService(SectorRepo)
	adminService := admin.NewAdminService(AdminRepo)

	return Dependencies{
		WorkerService:      workerService,
		AuthService:        authService,
		EmployerService:    employerService,
		JobService:         jobService,
		ApplicationService: applicationService,
		SectorService:      sectorService,
		AdminService:       adminService,
	}
}
