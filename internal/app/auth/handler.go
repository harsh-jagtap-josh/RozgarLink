package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/logger"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/middleware"
	"go.uber.org/zap"
)

func RegisterWorker(workerSvc worker.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var workerData worker.Worker
		err := json.NewDecoder(r.Body).Decode(&workerData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrInvalidRequestBody.Error()+", "+err.Error(), http.StatusBadRequest)
			return
		}

		response, err := workerSvc.CreateWorker(ctx, workerData)
		if err != nil {
			if errors.Is(err, apperrors.ErrInvalidUserDetails) {
				logger.Errorw(ctx, apperrors.ErrCreateWorker.Error(), zap.Error(err))
				middleware.HandleErrorResponse(ctx, w, apperrors.ErrCreateWorker.Error()+err.Error(), http.StatusBadRequest)
				return
			}

			if errors.Is(err, apperrors.ErrWorkerAlreadyExists) {
				logger.Errorw(ctx, apperrors.ErrWorkerAlreadyExists.Error(), zap.Error(err), zap.String("email: ", workerData.Email))
				middleware.HandleErrorResponse(ctx, w, apperrors.ErrCreateWorker.Error()+": "+err.Error()+", email: "+workerData.Email, http.StatusConflict)
				return
			}

			logger.Errorw(ctx, apperrors.ErrCreateWorker.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrCreateWorker.Error()+": "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully created a new worker account", http.StatusCreated, response)
	}
}

func HandleLogin(authService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrInvalidRequestBody.Error()+", "+err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := authService.Login(ctx, req)
		if err != nil {
			if errors.Is(err, apperrors.ErrInvalidLoginCredentials) {
				logger.Errorw(ctx, apperrors.ErrFailedLogin.Error(), zap.Error(err))
				middleware.HandleErrorResponse(ctx, w, apperrors.ErrFailedLogin.Error()+": "+err.Error(), http.StatusBadRequest)
				return
			}

			logger.Errorw(ctx, apperrors.ErrFailedLogin.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		if resp.Token == "" {
			logger.Errorw(ctx, apperrors.ErrInvalidLoginCredentials.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrInvalidLoginCredentials.Error(), http.StatusBadRequest)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully logged in "+resp.User.Role, http.StatusOK, resp)
	}
}
