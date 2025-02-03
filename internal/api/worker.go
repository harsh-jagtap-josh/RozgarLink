package api

import (
	"encoding/json"
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
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		if response.ID == 0 {
			http.Error(w, "no worker found with ID", http.StatusNotFound)
			return
		}

		httpResp, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "error occured while marshaling response payload", http.StatusInternalServerError)
			return
		}
		w.Write(httpResp)

	}
}
