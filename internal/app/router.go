package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
)

func NewRouter(deps Dependencies) *mux.Router {
	router := mux.NewRouter()

	workersRouter := router.PathPrefix("/workers").Subrouter()
	workersRouter.HandleFunc("/{worker_id}", worker.FetchWorkerByID(deps.WorkerService)).Methods(http.MethodGet)

	return router
}
