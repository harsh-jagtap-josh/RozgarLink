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
		rawProductID := vars["worker_id"]
		productID, err := strconv.Atoi(rawProductID)
		if err != nil {
			middleware.HandleErrorResponse(w, "error occured while converting workerID to an integer", http.StatusInternalServerError)
			return
		}

		response, err := workerSvc.GetWorkerByID(ctx, nil, int64(productID))
		if err != nil {
			middleware.HandleErrorResponse(w, "error occured while fetching from database", http.StatusInternalServerError)
			return
		}

		if response.ID == 0 {
			middleware.HandleErrorResponse(w, "error occured: no worker found with ID", http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(w, "worker details retrieved successfully", http.StatusOK, response)
	}
}
