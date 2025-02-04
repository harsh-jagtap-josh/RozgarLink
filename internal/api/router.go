package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app"
)

func NewRouter(deps app.Dependencies) *mux.Router {
	router := mux.NewRouter()

	workersRouter := router.PathPrefix("/workers").Subrouter()
	workersRouter.HandleFunc("/{worker_id}", HandleWorkerByID(deps.WorkerService)).Methods(http.MethodGet)

	return router
}
