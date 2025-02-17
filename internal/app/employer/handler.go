package employer

import (
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

		// response, err := workerSvc.FetchWorkerByID(ctx, workerID)
		// if err != nil {
		// 	if errors.Is(err, apperrors.ErrNoWorkerExists) {
		// 		logger.Errorw(ctx, apperrors.ErrNoWorkerExists.Error(), zap.Error(err), zap.String("ID", id))
		// 		httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchWorker, apperrors.ErrNoWorkerExists.Error(), id)
		// 		middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusNotFound)
		// 		return
		// 	}

		// 	logger.Errorw(ctx, apperrors.MsgFetchFromDb, zap.Error(err), zap.String("ID", id))
		// 	httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.MsgFailedToFetchWorker, apperrors.MsgFetchFromDb, id)
		// 	middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusInternalServerError)
		// 	return
		// }

		// middleware.HandleSuccessResponse(ctx, w, "worker details retrieved successfully", http.StatusOK, response)
	}
}
