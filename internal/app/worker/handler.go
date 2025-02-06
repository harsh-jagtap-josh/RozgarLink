package worker

import (
	"database/sql"
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

		vars := mux.Vars(r)
		id := vars["worker_id"]
		workerID, err := strconv.Atoi(id)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrConvertIdToInt, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrFailedToFetchWorker, apperrors.ErrConvertIdToInt, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusInternalServerError)
			return
		}

		response, err := workerSvc.GetWorkerByID(ctx, workerID)
		if errors.Is(err, sql.ErrNoRows) {
			logger.Errorw(ctx, apperrors.ErrNoWorkerExists, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrFailedToFetchWorker, apperrors.ErrNoWorkerExists, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusInternalServerError)
			return
		}

		if err != nil {
			logger.Errorw(ctx, apperrors.ErrFetchFromDb, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrFailedToFetchWorker, apperrors.ErrFetchFromDb, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "worker details retrieved successfully", http.StatusOK, response)
	}
}
