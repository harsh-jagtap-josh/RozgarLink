package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/admin"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/application"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/auth"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/employer"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/job"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/sector"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
)

func NewRouter(deps Dependencies) *mux.Router {

	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))

	// Auth Routes
	router.HandleFunc("/login", auth.HandleLogin(deps.AuthService)).Methods(http.MethodPost)
	router.HandleFunc("/register/worker", auth.RegisterWorker(deps.WorkerService)).Methods(http.MethodPost)
	router.HandleFunc("/register/employer", employer.RegisterEmployer(deps.EmployerService)).Methods(http.MethodPost)

	// adminRouter.Use(middleware.RequireSuperAdminRole)
	adminRouter := router.PathPrefix("").Subrouter()
	adminRouter.HandleFunc("/register/admin", admin.RegisterAdmin(deps.AdminService)).Methods(http.MethodPost)

	// Worker Routes - protected routes
	workerRouter := router.PathPrefix("/worker").Subrouter()

	// workerRouter.Use(middleware.ValidateJWT)            // validate JWT token
	// workerRouter.Use(middleware.RequireSameUserOrAdmin) // only worker with same ID has access to or admin

	workerRouter.HandleFunc("/{worker_id}", worker.FetchWorkerByID(deps.WorkerService)).Methods(http.MethodGet)
	workerRouter.HandleFunc("/{worker_id}", worker.UpdateWorkerByID(deps.WorkerService)).Methods(http.MethodPut)
	workerRouter.HandleFunc("/{worker_id}", worker.DeleteWorkerByID(deps.WorkerService)).Methods(http.MethodDelete)
	workerRouter.HandleFunc("/{worker_id}"+"/applications", worker.FetchApplicationsByWorkerId(deps.WorkerService)).Methods(http.MethodGet)

	// Employer Routes
	employerRouter := router.PathPrefix("/employer").Subrouter()
	employerRouter.HandleFunc("/{employer_id}", employer.FetchEmployerByID(deps.EmployerService)).Methods(http.MethodGet)
	employerRouter.HandleFunc("/{employer_id}", employer.UpdateEmployerById(deps.EmployerService)).Methods(http.MethodPut)
	employerRouter.HandleFunc("/{employer_id}", employer.DeleteEmployerByID(deps.EmployerService)).Methods(http.MethodDelete)
	employerRouter.HandleFunc("/{employer_id}"+"/jobs", employer.FetchJobsByEmployerId(deps.EmployerService)).Methods(http.MethodGet)

	// Job Routes
	jobRouter := router.PathPrefix("/job").Subrouter()
	jobRouter.HandleFunc("/create", job.CreateJob(deps.JobService)).Methods(http.MethodPost)
	jobRouter.HandleFunc("/all", job.FetchAllJobs(deps.JobService)).Methods(http.MethodGet)
	jobRouter.HandleFunc("/{job_id}", job.FetchJobByID(deps.JobService)).Methods(http.MethodGet)
	jobRouter.HandleFunc("/{job_id}", job.UpdateJobById(deps.JobService)).Methods(http.MethodPut)
	jobRouter.HandleFunc("/{job_id}", job.DeleteJobByID(deps.JobService)).Methods(http.MethodDelete)
	jobRouter.HandleFunc("/{job_id}"+"/applications", job.FetchApplicationsByJobId(deps.JobService)).Methods(http.MethodGet)

	// Application Routes
	applicationRouter := router.PathPrefix("/application").Subrouter()
	applicationRouter.HandleFunc("/create", application.CreateNewApplication(deps.ApplicationService)).Methods(http.MethodPost)
	applicationRouter.HandleFunc("/{application_id}", application.FetchApplicationByID(deps.ApplicationService)).Methods(http.MethodGet)
	applicationRouter.HandleFunc("/{application_id}", application.UpdateApplicationByID(deps.ApplicationService)).Methods(http.MethodPut)
	applicationRouter.HandleFunc("/{application_id}", application.DeleteApplicationByID(deps.ApplicationService)).Methods(http.MethodDelete)

	// Sectors Routes - Only Admin has access to all these routes
	sectorRouter := router.PathPrefix("/sector").Subrouter()
	sectorRouter.HandleFunc("/create", sector.CreateSector(deps.SectorService)).Methods(http.MethodPost)
	sectorRouter.HandleFunc("/all", sector.FetchAllSectors(deps.SectorService)).Methods(http.MethodGet)
	sectorRouter.HandleFunc("/{sector_id}", sector.FetchSectorById(deps.SectorService)).Methods(http.MethodGet)
	sectorRouter.HandleFunc("/{sector_id}", sector.UpdateSectorById(deps.SectorService)).Methods(http.MethodPut)
	sectorRouter.HandleFunc("/{sector_id}", sector.DeleteSectorById(deps.SectorService)).Methods(http.MethodDelete)

	// Routes to Fetch Complete Data
	router.HandleFunc("/workers", worker.FetchAllWorkers(deps.WorkerService)).Methods(http.MethodGet)
	router.HandleFunc("/employers", employer.FetchAllEmployers(deps.EmployerService)).Methods(http.MethodGet)
	router.HandleFunc("/applications", application.FetchAllApplications(deps.ApplicationService)).Methods(http.MethodGet)
	router.HandleFunc("/jobs", job.FetchAllJobs(deps.JobService)).Methods(http.MethodGet)

	return router
}
