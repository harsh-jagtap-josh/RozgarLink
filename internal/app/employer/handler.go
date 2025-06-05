package employer

import (
	"context"
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
		employerId, id := isEmployerIdValid(ctx, w, r, errors.New(apperrors.MsgFailedToFetchEmp))
		if employerId == -1 {
			return
		}

		employer, err := employerSvc.FetchEmployerByID(ctx, employerId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoEmployerExists) {
				logger.Errorw(ctx, apperrors.ErrNoEmployerExists.Error(), zap.Error(err), zap.String("ID", id))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchEmp, apperrors.ErrNoEmployerExists.Error(), id)
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.MsgFetchFromDbErr, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchEmp, apperrors.MsgFetchFromDbErr, id)
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

		employerId, id := isEmployerIdValid(ctx, w, r, apperrors.ErrUpdateEmployer)
		if employerId == -1 {
			return
		}

		var employerData Employer

		err := json.NewDecoder(r.Body).Decode(&employerData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrUpdateEmployer.Error(), apperrors.ErrInvalidRequestBody.Error(), id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
			return
		}
		employerData.ID = employerId
		response, err := employerSvc.UpdateEmployerById(ctx, employerData)
		if err != nil {
			if errors.Is(err, apperrors.ErrInvalidUserDetails) {
				logger.Errorw(ctx, apperrors.ErrUpdateEmployer.Error(), zap.Error(err))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrUpdateEmployer.Error(), err.Error(), id)
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

		employer, err := employerSvc.RegisterEmployer(ctx, employerData)
		if err != nil {
			if errors.Is(err, apperrors.ErrInvalidUserDetails) {
				logger.Errorw(ctx, apperrors.ErrCreateEmployer.Error(), zap.Error(err))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrCreateEmployer.Error(), err.Error(), fmt.Sprintf("%d", employer.ID))
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
				return
			}

			if errors.Is(err, apperrors.ErrEmployerAlreadyExists) {
				logger.Errorw(ctx, apperrors.ErrCreateEmployer.Error(), zap.Error(err), zap.String("email: ", employerData.Email))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrCreateEmployer.Error(), err.Error(), fmt.Sprintf("%d", employer.ID))
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

		employerId, id := isEmployerIdValid(ctx, w, r, apperrors.ErrDeleteEmployer)
		if employerId == -1 {
			return
		}

		_, err := employerSvc.DeleteEmployerById(ctx, employerId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoEmployerExists) {
				logger.Errorw(ctx, apperrors.ErrNoEmployerExists.Error(), zap.Error(err), zap.String("ID", id))
				httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrDeleteEmployer.Error(), apperrors.ErrNoEmployerExists.Error(), id)
				middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.ErrDeleteEmployer.Error(), zap.Error(err))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrDeleteEmployer.Error(), err.Error(), id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func FetchJobsByEmployerId(es Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		employerId, id := isEmployerIdValid(ctx, w, r, apperrors.ErrFetchJobs)
		if employerId == -1 {
			return
		}

		jobs, err := es.FetchJobsByEmployerId(ctx, employerId)
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

func isEmployerIdValid(ctx context.Context, w http.ResponseWriter, r *http.Request, errType error) (int, string) {

	vars := mux.Vars(r)
	id := vars["employer_id"]
	employerID, err := strconv.Atoi(id)
	if err != nil {
		logger.Errorw(ctx, apperrors.MsgInvalidEmployerId, zap.Error(err), zap.String("ID", id))
		httpResponseMsg := apperrors.HttpErrorResponseMessage(errType.Error(), apperrors.MsgInvalidEmployerId, id)
		middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
		return -1, id
	}
	return employerID, id
}

func FetchAllEmployers(empS Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		employers, err := empS.FetchAllEmployers(ctx)
		if err != nil {
			logger.Errorw(ctx, apperrors.MsgFailedToFetchEmp, zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.MsgFailedToFetchEmp+", "+err.Error(), http.StatusInternalServerError)
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully fetched employers data", http.StatusOK, employers)
	}
}
