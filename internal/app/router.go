package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/auth"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/employer"
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
	router.HandleFunc("/register/employer", auth.RegisterWorker(deps.WorkerService)).Methods(http.MethodPost)

	// Worker Routes - protected routes
	workerRouter := router.PathPrefix("/workers").Subrouter()
	// workerRouter.Use(middleware.ValidateJWT) // validate JWT token
	workerRouter.HandleFunc(workerIdParam, worker.FetchWorkerByID(deps.WorkerService)).Methods(http.MethodGet)
	// workerRouter.Use(middleware.RequireSameUserOrAdmin) // only worker with same ID has access to or admin
	workerRouter.HandleFunc(workerIdParam, worker.UpdateWorkeByID(deps.WorkerService)).Methods(http.MethodPut)
	workerRouter.HandleFunc(workerIdParam, worker.DeleteWorkeByID(deps.WorkerService)).Methods(http.MethodDelete)

	// Employer Routes
	employerRouter := router.PathPrefix("/employers").Subrouter()
	employerRouter.HandleFunc(employerIdParam, employer.FetchEmployerByID(deps.EmployerService)).Methods(http.MethodGet)

	return router
}
