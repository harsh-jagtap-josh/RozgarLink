package employer

import (
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

func FetchEmployerByID(employerSvc Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		vars := mux.Vars(r)
		id := vars["employer_id"]
		employerID, err := strconv.Atoi(id)
		if err != nil {
			logger.Errorw(ctx, apperrors.MsgInvalidEmployerId, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchEmp, apperrors.MsgInvalidEmployerId, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
			return
		}

		employer, err := employerSvc.FetchEmployerByID(ctx, employerID)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoEmployerExists) {
				logger.Errorw(ctx, apperrors.ErrNoEmployerExists.Error(), zap.Error(err), zap.String("ID", id))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchEmp, apperrors.ErrNoEmployerExists.Error(), id)
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.MsgFetchFromDb, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchWorker, apperrors.MsgFetchFromDb, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusInternalServerError)
			return
		}
		middleware.HandleSuccessResponse(ctx, w, "employer details retrieved successfully", http.StatusOK, employer)
	}
}

func UpdateEmployerById(employerSvc Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		vars := mux.Vars(r)
		id := vars["employer_id"]
		_, err := strconv.Atoi(id)
		if err != nil {
			logger.Errorw(ctx, apperrors.MsgInvalidEmployerId, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchEmp, apperrors.MsgInvalidEmployerId, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
			return
		}

		var employerData Employer

		err = json.NewDecoder(r.Body).Decode(&employerData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrInvalidRequestBody.Error()+err.Error(), http.StatusBadRequest)
			return
		}

		response, err, errType := employerSvc.UpdateEmployerById(ctx, employerData)
		if err != nil {
			if errors.Is(errType, apperrors.ErrInvalidUserDetails) {
				logger.Errorw(ctx, apperrors.ErrInvalidUserDetails.Error(), zap.Error(err))
				http.Error(w, apperrors.ErrInvalidUserDetails.Error()+", "+err.Error(), http.StatusBadRequest)
				return
			}
			logger.Errorw(ctx, apperrors.ErrUpdateEmployer.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrUpdateEmployer.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully updated employer details", http.StatusOK, response)
	}
}
