package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/auth"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/employer"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/job"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
)

// Query Params
const workerIdParam = "/{worker_id}"
const employerIdParam = "/{employer_id}"

func NewRouter(deps Dependencies) *mux.Router {

	router := mux.NewRouter()

	// Auth Routes
	router.HandleFunc("/login", auth.HandleLogin(deps.AuthService)).Methods(http.MethodPost)
	router.HandleFunc("/register/worker", auth.RegisterWorker(deps.WorkerService)).Methods(http.MethodPost)
	router.HandleFunc("/register/employer", employer.RegisterEmployer(deps.EmployerService)).Methods(http.MethodPost)

	// Worker Routes - protected routes
	workerRouter := router.PathPrefix("/worker").Subrouter()
	// workerRouter.Use(middleware.ValidateJWT) // validate JWT token
	workerRouter.HandleFunc(workerIdParam, worker.FetchWorkerByID(deps.WorkerService)).Methods(http.MethodGet)
	// workerRouter.Use(middleware.RequireSameUserOrAdmin) // only worker with same ID has access to or admin
	workerRouter.HandleFunc(workerIdParam, worker.UpdateWorkerByID(deps.WorkerService)).Methods(http.MethodPut)
	workerRouter.HandleFunc(workerIdParam, worker.DeleteWorkerByID(deps.WorkerService)).Methods(http.MethodDelete)

	// Employer Routes
	employerRouter := router.PathPrefix("/employer").Subrouter()
	employerRouter.HandleFunc(employerIdParam, employer.FetchEmployerByID(deps.EmployerService)).Methods(http.MethodGet)
	employerRouter.HandleFunc(employerIdParam, employer.UpdateEmployerById(deps.EmployerService)).Methods(http.MethodPut)
	employerRouter.HandleFunc(employerIdParam, employer.DeleteEmployerByID(deps.EmployerService)).Methods(http.MethodDelete)

	// Job Routes
	jobRouter := router.PathPrefix("/job").Subrouter()
	jobRouter.HandleFunc("/create", job.CreateJob(deps.JobService)).Methods(http.MethodPost)
	jobRouter.HandleFunc("/{job_id}", job.UpdateJobById(deps.JobService)).Methods(http.MethodPut)
	return router
}
