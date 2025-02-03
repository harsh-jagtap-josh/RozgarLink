package api

import (
	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app"
)

func NewRouter(deps app.Dependencies) *mux.Router {
	router := mux.NewRouter()

	WorkersRouter := router.PathPrefix("/workers").Subrouter()
	WorkersRouter.HandleFunc("/{worker_id}", HandleWorkerByID(deps.WorkerService)).Methods("GET")

	return router
}
