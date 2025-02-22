package application

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

func CreateNewApplication(appService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		var applicationData Application
		err := json.NewDecoder(r.Body).Decode(&applicationData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrInvalidRequestBody.Error()+": "+err.Error(), http.StatusBadRequest)
			return
		}

		createdAppl, err := appService.CreateNewApplication(ctx, applicationData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrCreateApplication.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrCreateApplication.Error()+": "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully created new application", http.StatusCreated, createdAppl)
	}
}

func UpdateApplicationByID(appService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		applicationId, _ := isApplicationIdValid(ctx, w, r, apperrors.ErrUpdateApplication)
		if applicationId == -1 {
			return
		}

		var applicationData Application
		err := json.NewDecoder(r.Body).Decode(&applicationData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrUpdateApplication.Error()+": "+apperrors.ErrInvalidRequestBody.Error()+": "+err.Error(), http.StatusBadRequest)
			return
		}

		updatedApplication, err := appService.UpdateApplicationById(ctx, applicationData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrUpdateApplication.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrUpdateApplication.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}
		middleware.HandleSuccessResponse(ctx, w, "successfully updated application details", http.StatusOK, updatedApplication)
	}
}

func FetchApplicationByID(appService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		applicationId, id := isApplicationIdValid(ctx, w, r, apperrors.ErrFetchApplication)
		if applicationId == -1 {
			return
		}
		application, err := appService.FetchApplicationById(ctx, applicationId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoApplicationExists) {
				logger.Errorw(ctx, apperrors.ErrFetchApplication.Error(), zap.Error(err), zap.String("ID", id))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrFetchApplication.Error(), err.Error(), id)
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.MsgFetchFromDbErr, zap.Error(err), zap.String("ID", id))
			middleware.HandleErrorResponse(ctx, w, apperrors.MsgFailedToFetchApplication+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "applications details retrieved successfully", http.StatusOK, application)
	}
}

func DeleteApplicationByID(appService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		applicationId, id := isApplicationIdValid(ctx, w, r, apperrors.ErrDeleteApplication)
		if applicationId == -1 {
			return
		}

		_, err := appService.DeleteApplicationById(ctx, applicationId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoApplicationExists) {
				logger.Errorw(ctx, apperrors.ErrDeleteApplication.Error(), zap.Error(err), zap.String("ID", id))
				middleware.HandleErrorResponse(ctx, w, apperrors.ErrDeleteApplication.Error()+": "+err.Error(), http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.ErrDeleteApplication.Error(), zap.Error(err), zap.String("ID", id))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrDeleteApplication.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func FetchAllApplications(jobService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		applications, err := jobService.FetchAllApplications(ctx)
		if err != nil {
			logger.Errorw(ctx, apperrors.MsgFailedToFetchApplication, zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.MsgFailedToFetchApplication+", "+err.Error(), http.StatusInternalServerError)
		}

		middleware.HandleSuccessResponse(ctx, w, "applications retrieved successfully", http.StatusOK, applications)
	}
}

func isApplicationIdValid(ctx context.Context, w http.ResponseWriter, r *http.Request, errType error) (int, string) {
	vars := mux.Vars(r)
	id := vars["application_id"]
	applicationId, err := strconv.Atoi(id)
	if err != nil {
		logger.Errorw(ctx, apperrors.MsgInvalidApplicationId, zap.Error(err), zap.String("ID", id))
		httpResponseMsg := apperrors.HttpErrorResponseMessage(errType.Error(), apperrors.MsgInvalidApplicationId, id)
		middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
		return -1, id
	}
	return applicationId, id
}
