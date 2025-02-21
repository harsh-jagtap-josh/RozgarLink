package worker

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

func FetchWorkerByID(workerSvc Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		workerId, id := isWorkerIdValid(ctx, w, r)
		if workerId == -1 {
			return
		}

		response, err := workerSvc.FetchWorkerByID(ctx, workerId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoWorkerExists) {
				logger.Errorw(ctx, apperrors.ErrNoWorkerExists.Error(), zap.Error(err), zap.String("ID", id))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchWorker, apperrors.ErrNoWorkerExists.Error(), id)
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.MsgFetchFromDb, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchWorker, apperrors.MsgFetchFromDb, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "worker details retrieved successfully", http.StatusOK, response)
	}
}

func UpdateWorkerByID(workerSvc Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var workerData Worker

		workerId, id := isWorkerIdValid(ctx, w, r)
		if workerId == -1 {
			return
		}

		err := json.NewDecoder(r.Body).Decode(&workerData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err), zap.String("ID", id))
			http.Error(w, apperrors.ErrInvalidRequestBody.Error()+err.Error(), http.StatusBadRequest)
			return
		}

		response, err, errType := workerSvc.UpdateWorkerByID(ctx, workerData)

		if err != nil {
			if errors.Is(err, apperrors.ErrInvalidUserDetails) {
				logger.Errorw(ctx, errType.Error(), zap.Error(err))
				http.Error(w, errType.Error()+", "+err.Error(), http.StatusBadRequest)
				return
			}

			logger.Errorw(ctx, apperrors.ErrUpdateWorker.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrUpdateWorker.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully updated worker details", http.StatusOK, response)
	}
}

func DeleteWorkerByID(workerSvc Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		workerId, id := isWorkerIdValid(ctx, w, r)
		if workerId == -1 {
			return
		}

		_, err := workerSvc.DeleteWorkerByID(ctx, workerId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoWorkerExists) {
				logger.Errorw(ctx, apperrors.ErrNoWorkerExists.Error(), zap.Error(err), zap.String("ID", id))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchWorker, apperrors.ErrNoWorkerExists.Error(), id)
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.ErrDeleteWorker.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrDeleteWorker.Error()+","+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func FetchApplicationsByWorkerId(workerSvc Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		workerId, id := isWorkerIdValid(ctx, w, r)
		if workerId == -1 {
			return
		}

		applications, err := workerSvc.FetchApplicationsByWorkerId(ctx, workerId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoWorkerExists) {
				logger.Errorw(ctx, apperrors.ErrNoWorkerExists.Error(), zap.Error(err), zap.String("ID", id))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrFetchApplication.Error(), apperrors.ErrNoWorkerExists.Error(), id)
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.ErrFetchApplication.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrFetchApplication.Error()+","+err.Error(), http.StatusInternalServerError)
			return
		}
		middleware.HandleSuccessResponse(ctx, w, "successfully fetched applications details", http.StatusOK, applications)
	}
}

func isWorkerIdValid(ctx context.Context, w http.ResponseWriter, r *http.Request) (int, string) {
	// retrieve id from query params
	vars := mux.Vars(r)
	id := vars["worker_id"]
	workerId, err := strconv.Atoi(id)
	if err != nil {
		logger.Errorw(ctx, apperrors.MsgInvalidWorkerId, zap.Error(err), zap.String("ID", id))
		httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchWorker, apperrors.MsgInvalidWorkerId, id)
		middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
		return -1, ""
	}

	return workerId, id
}
