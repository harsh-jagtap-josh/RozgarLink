package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/auth"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
)

func NewRouter(deps Dependencies) *mux.Router {

	router := mux.NewRouter()

	// Auth Routes
	router.HandleFunc("/register/worker", auth.RegisterWorker(deps.WorkerService)).Methods(http.MethodPost)
	router.HandleFunc("/login", auth.HandleLogin(deps.AuthService)).Methods(http.MethodPost)

	// Worker Routes - protected routes
	workerRouter := router.PathPrefix("/workers").Subrouter()
	// workerRouter.Use(middleware.ValidateJWT) // validate JWT token
	workerRouter.HandleFunc("/{worker_id}", worker.FetchWorkerByID(deps.WorkerService)).Methods(http.MethodGet)

	// workerRouter.Use(middleware.RequireSameUserOrAdmin) // only worker with same ID has access to or admin
	workerRouter.HandleFunc("/{worker_id}", worker.UpdateWorkeByID(deps.WorkerService)).Methods(http.MethodPut)
	workerRouter.HandleFunc("/{worker_id}", worker.DeleteWorkeByID(deps.WorkerService)).Methods(http.MethodDelete)

	return router
}
