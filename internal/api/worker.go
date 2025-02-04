package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	middleware "github.com/harsh-jagtap-josh/RozgarLink/internal/pkg"
)

func HandleWorkerByID(workerSvc worker.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		vars := mux.Vars(r)
		id := vars["worker_id"]
		workerID, err := strconv.Atoi(id)
		if err != nil {
			middleware.HandleErrorResponse(w, "error occured while converting workerID to an integer", http.StatusInternalServerError)
			return
		}

		response, err := workerSvc.GetWorkerByID(ctx, int64(workerID))
		if err != nil {
			middleware.HandleErrorResponse(w, "error occured while fetching from database", http.StatusInternalServerError)
			return
		}

		if response.ID == 0 {
			middleware.HandleErrorResponse(w, "error occured: no worker found with ID: "+id, http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(w, "worker details retrieved successfully", http.StatusOK, response)
	}
}
