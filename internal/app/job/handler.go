package job

import (
	"context"
	"encoding/json"
	"errors"
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
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrInvalidRequestBody.Error()+": "+err.Error(), http.StatusBadRequest)
			return
		}

		createdJob, err := js.CreateJob(ctx, jobData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrCreateJob.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrCreateJob.Error()+": "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully create new job", http.StatusCreated, createdJob)
	}
}

func UpdateJobById(js Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		jobId, _ := isJobIdValid(ctx, w, r, apperrors.ErrUpdateJob)
		if jobId == -1 {
			return
		}

		var jobData Job

		err := json.NewDecoder(r.Body).Decode(&jobData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrInvalidRequestBody.Error()+": "+err.Error(), http.StatusBadRequest)
			return
		}

		updatedJob, err := js.UpdateJobByID(ctx, jobData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrUpdateJob.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrUpdateJob.Error()+": "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully updated job details", http.StatusOK, updatedJob)
	}
}

func FetchJobByID(js Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		jobId, id := isJobIdValid(ctx, w, r, apperrors.ErrFetchJob)
		if jobId == -1 {
			return
		}

		job, err := js.FetchJobByID(ctx, jobId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoJobExists) {
				logger.Errorw(ctx, apperrors.ErrNoJobExists.Error(), zap.Error(err), zap.String("ID", id))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchJob, apperrors.ErrNoJobExists.Error(), id)
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.MsgFetchFromDbErr, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchJob, apperrors.MsgFetchFromDbErr, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "job details retrieved successfully", http.StatusOK, job)
	}
}

func DeleteJobByID(js Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		jobId, id := isJobIdValid(ctx, w, r, apperrors.ErrDeleteJob)
		if jobId == -1 {
			return
		}

		_, err := js.DeleteJobByID(ctx, jobId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoJobExists) {
				logger.Errorw(ctx, apperrors.ErrNoJobExists.Error(), zap.Error(err), zap.String("ID", id))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrDeleteJob.Error(), apperrors.ErrNoJobExists.Error(), id)
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusNotFound)
				return
			}
			logger.Errorw(ctx, apperrors.ErrDeleteJob.Error(), zap.Error(err), zap.String("ID", id))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrDeleteJob.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func FetchApplicationsByJobId(js Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		jobId, id := isJobIdValid(ctx, w, r, apperrors.ErrFetchApplication)
		if jobId == -1 {
			return
		}

		applications, err := js.FetchApplicationsByJobId(ctx, jobId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoJobExists) {
				logger.Errorw(ctx, apperrors.ErrNoJobExists.Error(), zap.Error(err), zap.String("ID", id))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrFetchApplication.Error(), apperrors.ErrNoJobExists.Error(), id)
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.ErrFetchApplication.Error(), zap.Error(err), zap.String("ID", id))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrFetchApplication.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "job applications retrieved successfully", http.StatusOK, applications)
	}
}

func FetchAllJobs(jobService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		queryParams := r.URL.Query()

		jobFilters := retrieveQueryParams(queryParams)

		jobs, err := jobService.FetchAllJobs(ctx, jobFilters)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrFetchJobs.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrFetchJobs.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "jobs retrieved successfully", http.StatusOK, jobs)
	}
}

func isJobIdValid(ctx context.Context, w http.ResponseWriter, r *http.Request, errType error) (int, string) {
	vars := mux.Vars(r)
	id := vars["job_id"]
	jobId, err := strconv.Atoi(id)
	if err != nil {
		logger.Errorw(ctx, apperrors.MsgInvalidJobId, zap.Error(err), zap.String("ID", id))
		httpResponseMsg := apperrors.HttpErrorResponseMessage(errType.Error(), apperrors.MsgInvalidJobId, id)
		middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
		return -1, id
	}
	return jobId, id

}
