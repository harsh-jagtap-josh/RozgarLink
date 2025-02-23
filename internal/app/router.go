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

// Query Params Constants
const (
	workerIdParam      = "/{worker_id}"
	employerIdParam    = "/{employer_id}"
	jobIdParam         = "/{job_id}"
	applicationIdParam = "/{application_id}"
	sectorIdParam      = "/{sector_id}"
)

func NewRouter(deps Dependencies) *mux.Router {

	router := mux.NewRouter()

	// Auth Routes
	router.HandleFunc("/login", auth.HandleLogin(deps.AuthService)).Methods(http.MethodPost)
	router.HandleFunc("/register/worker", auth.RegisterWorker(deps.WorkerService)).Methods(http.MethodPost)
	router.HandleFunc("/register/employer", employer.RegisterEmployer(deps.EmployerService)).Methods(http.MethodPost)

	adminRouter := router.PathPrefix("").Subrouter()
	// adminRouter.Use(middleware.RequireSuperAdminRole)
	adminRouter.HandleFunc("/register/admin", admin.RegisterAdmin(deps.AdminService)).Methods(http.MethodPost) // Only for Super Admin, create a new subrouter for this (so that only this route will have the validation of a super-admin) - may have to create a new jwt function to validate only super-admins
	

	// Worker Routes - protected routes
	workerRouter := router.PathPrefix("/worker").Subrouter()
	// workerRouter.Use(middleware.ValidateJWT) // validate JWT token
	workerRouter.HandleFunc(workerIdParam, worker.FetchWorkerByID(deps.WorkerService)).Methods(http.MethodGet)
	// workerRouter.Use(middleware.RequireSameUserOrAdmin) // only worker with same ID has access to or admin
	workerRouter.HandleFunc(workerIdParam, worker.UpdateWorkerByID(deps.WorkerService)).Methods(http.MethodPut)
	workerRouter.HandleFunc(workerIdParam, worker.DeleteWorkerByID(deps.WorkerService)).Methods(http.MethodDelete)
	workerRouter.HandleFunc(workerIdParam+"/applications", worker.FetchApplicationsByWorkerId(deps.WorkerService)).Methods(http.MethodGet)

	// Employer Routes
	employerRouter := router.PathPrefix("/employer").Subrouter()
	employerRouter.HandleFunc(employerIdParam, employer.FetchEmployerByID(deps.EmployerService)).Methods(http.MethodGet)
	employerRouter.HandleFunc(employerIdParam, employer.UpdateEmployerById(deps.EmployerService)).Methods(http.MethodPut)
	employerRouter.HandleFunc(employerIdParam, employer.DeleteEmployerByID(deps.EmployerService)).Methods(http.MethodDelete)
	employerRouter.HandleFunc(employerIdParam+"/jobs", employer.FetchJobsByEmployerId(deps.EmployerService)).Methods(http.MethodGet)

	// Job Routes
	jobRouter := router.PathPrefix("/job").Subrouter()
	jobRouter.HandleFunc("/create", job.CreateJob(deps.JobService)).Methods(http.MethodPost)
	jobRouter.HandleFunc("/all", job.FetchAllJobs(deps.JobService)).Methods(http.MethodGet)
	jobRouter.HandleFunc(jobIdParam, job.FetchJobByID(deps.JobService)).Methods(http.MethodGet)
	jobRouter.HandleFunc(jobIdParam, job.UpdateJobById(deps.JobService)).Methods(http.MethodPut)
	jobRouter.HandleFunc(jobIdParam, job.DeleteJobByID(deps.JobService)).Methods(http.MethodDelete)
	jobRouter.HandleFunc(jobIdParam+"/applications", job.FetchApplicationsByJobId(deps.JobService)).Methods(http.MethodGet)

	// Application Routes
	applicationRouter := router.PathPrefix("/application").Subrouter()
	applicationRouter.HandleFunc("/create", application.CreateNewApplication(deps.ApplicationService)).Methods(http.MethodPost)
	applicationRouter.HandleFunc(applicationIdParam, application.FetchApplicationByID(deps.ApplicationService)).Methods(http.MethodGet)
	applicationRouter.HandleFunc(applicationIdParam, application.UpdateApplicationByID(deps.ApplicationService)).Methods(http.MethodPut)
	applicationRouter.HandleFunc(applicationIdParam, application.DeleteApplicationByID(deps.ApplicationService)).Methods(http.MethodDelete)

	// Sectors Routes - Only Admin has access to all these routes
	sectorRouter := router.PathPrefix("/sector").Subrouter()
	sectorRouter.HandleFunc("/create", sector.CreateSector(deps.SectorService)).Methods(http.MethodPost)
	sectorRouter.HandleFunc("/all", sector.FetchAllSectors(deps.SectorService)).Methods(http.MethodGet)
	sectorRouter.HandleFunc(sectorIdParam, sector.FetchSectorById(deps.SectorService)).Methods(http.MethodGet)
	sectorRouter.HandleFunc(sectorIdParam, sector.UpdateSectorById(deps.SectorService)).Methods(http.MethodPut)
	sectorRouter.HandleFunc(sectorIdParam, sector.DeleteSectorById(deps.SectorService)).Methods(http.MethodDelete)

	// Routes to Fetch Complete Data - mostly only admin will have access to all this data
	router.HandleFunc("/workers", worker.FetchAllWorkers(deps.WorkerService)).Methods(http.MethodGet)
	router.HandleFunc("/employers", employer.FetchAllEmployers(deps.EmployerService)).Methods(http.MethodGet)
	router.HandleFunc("/applications", application.FetchAllApplications(deps.ApplicationService)).Methods(http.MethodGet)
	router.HandleFunc("/jobs", job.FetchAllJobs(deps.JobService)).Methods(http.MethodGet)

	return router
}
