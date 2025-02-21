package employer

import (
	"encoding/json"
	"errors"
	"fmt"
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
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchEmp, apperrors.MsgFetchFromDb, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusInternalServerError)
			return
		}
		middleware.HandleSuccessResponse(ctx, w, "employer details retrieved successfully", http.StatusOK, employer)
	}
}

// update employer based on provided details, that also contains employer id and address id.
func UpdateEmployerById(employerSvc Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		vars := mux.Vars(r)
		id := vars["employer_id"]
		_, err := strconv.Atoi(id)
		if err != nil {
			logger.Errorw(ctx, apperrors.MsgInvalidEmployerId, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrUpdateEmployer.Error(), apperrors.MsgInvalidEmployerId, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
			return
		}

		var employerData Employer

		err = json.NewDecoder(r.Body).Decode(&employerData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrUpdateEmployer.Error(), apperrors.ErrInvalidRequestBody.Error(), id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
			return
		}

		response, err, errType := employerSvc.UpdateEmployerById(ctx, employerData)
		if err != nil {
			if errors.Is(errType, apperrors.ErrInvalidUserDetails) {
				logger.Errorw(ctx, apperrors.ErrInvalidUserDetails.Error(), zap.Error(err))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrInvalidUserDetails.Error(), err.Error(), id)
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
				return
			}

			logger.Errorw(ctx, apperrors.ErrUpdateEmployer.Error(), zap.Error(err))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrUpdateEmployer.Error(), err.Error(), id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully updated employer details", http.StatusOK, response)
	}
}

func RegisterEmployer(employerSvc Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		var employerData Employer

		err := json.NewDecoder(r.Body).Decode(&employerData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrInvalidRequestBody.Error(), err.Error(), "nil")
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
			return
		}

		employer, err, errType := employerSvc.RegisterEmployer(ctx, employerData)
		if err != nil {
			if errors.Is(errType, apperrors.ErrInvalidUserDetails) {
				logger.Errorw(ctx, apperrors.ErrInvalidUserDetails.Error(), zap.Error(err))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrInvalidUserDetails.Error(), err.Error(), fmt.Sprintf("%d", employer.ID))
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
				return
			}

			if errors.Is(errType, apperrors.ErrEmployerAlreadyExists) {
				logger.Errorw(ctx, apperrors.ErrEmployerAlreadyExists.Error(), zap.Error(err), zap.String("email: ", employerData.Email))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrCreateEmployer.Error(), apperrors.ErrEmployerAlreadyExists.Error(), fmt.Sprintf("%d", employer.ID))
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusConflict)
				return
			}

			logger.Errorw(ctx, apperrors.ErrCreateEmployer.Error(), zap.Error(err))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrCreateEmployer.Error(), err.Error(), fmt.Sprintf("%d", employer.ID))
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully created employer details", http.StatusCreated, employer)
	}
}

func DeleteEmployerByID(employerSvc Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		vars := mux.Vars(r)
		id := vars["employer_id"]
		employerID, err := strconv.Atoi(id)
		if err != nil {
			logger.Errorw(ctx, apperrors.MsgInvalidEmployerId, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrDeleteEmployer.Error(), apperrors.MsgInvalidEmployerId, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
			return
		}

		_, err = employerSvc.DeleteEmployerById(ctx, employerID)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoEmployerExists) {
				logger.Errorw(ctx, apperrors.ErrNoEmployerExists.Error(), zap.Error(err), zap.String("ID", id))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrDeleteEmployer.Error(), apperrors.ErrNoEmployerExists.Error(), id)
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.ErrDeleteEmployer.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrDeleteEmployer.Error()+","+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func FetchJobsByEmployerId(es Service) func(w http.ResponseWriter, r *http.Request) {
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

		jobs, err := es.FetchJobsByEmployerId(ctx, employerID)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoEmployerExists) {
				logger.Errorw(ctx, apperrors.ErrNoEmployerExists.Error(), zap.Error(err), zap.String("ID", id))
				middleware.HandleErrorResponse(ctx, w, apperrors.ErrNoEmployerExists.Error(), http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.ErrFetchJobs.Error(), zap.Error(err), zap.String("ID", id))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrFetchJobs.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully fetched jobs", http.StatusOK, jobs)
	}
}
