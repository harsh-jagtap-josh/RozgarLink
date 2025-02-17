package job

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/logger"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/middleware"
	"go.uber.org/zap"
)

func CreateJob(js Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		var jobData Job
		err := json.NewDecoder(r.Body).Decode(&jobData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrInvalidRequestBody.Error()+err.Error(), http.StatusBadRequest)
			return
		}

		createdJob, err := js.CreateJob(ctx, jobData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrCreateJob.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrCreateJob.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully create new job", http.StatusCreated, createdJob)
	}
}

func UpdateJobById(js Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		id := vars["job_id"]
		_, err := strconv.Atoi(id)
		if err != nil {
			logger.Errorw(ctx, apperrors.MsgInvalidEmployerId, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrUpdateJob.Error(), apperrors.MsgInvalidJobId, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
			return
		}

		var jobData Job
		err = json.NewDecoder(r.Body).Decode(&jobData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrInvalidRequestBody.Error()+err.Error(), http.StatusBadRequest)
			return
		}

		updatedJob, err := js.UpdateJobByID(ctx, jobData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrUpdateJob.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrUpdateJob.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully updated job details", http.StatusOK, updatedJob)
	}
}
