package sector

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

func CreateSector(sectorService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var sectorData Sector
		err := json.NewDecoder(r.Body).Decode(&sectorData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrInvalidRequestBody.Error()+err.Error(), http.StatusBadRequest)
			return
		}

		createdSector, err := sectorService.CreateNewSector(ctx, sectorData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrCreateSector.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrCreateSector.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully created new sector", http.StatusCreated, createdSector)
	}
}

func FetchSectorById(sectorService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		id := vars["sector_id"]
		sectorId, err := strconv.Atoi(id)
		if err != nil {
			logger.Errorw(ctx, apperrors.MsgInvalidSectorId, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrFetchSector.Error(), apperrors.MsgInvalidSectorId, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
			return
		}

		sector, err := sectorService.FetchSectorById(ctx, sectorId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoSectorExists) {
				logger.Errorw(ctx, apperrors.ErrNoSectorExists.Error(), zap.Error(err), zap.String("ID", id))
				middleware.HandleErrorResponse(ctx, w, apperrors.ErrFetchSector.Error()+", "+apperrors.ErrNoSectorExists.Error(), http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.MsgFetchFromDb, zap.Error(err), zap.String("ID", id))
			middleware.HandleErrorResponse(ctx, w, apperrors.MsgFailedToFetchSector+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully created new sector", http.StatusOK, sector)
	}
}

func UpdateSectorById(sectorService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		id := vars["sector_id"]
		sectorId, err := strconv.Atoi(id)
		if err != nil {
			logger.Errorw(ctx, apperrors.MsgInvalidSectorId, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrUpdateSector.Error(), apperrors.MsgInvalidSectorId, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
			return
		}

		var sectorData Sector
		err = json.NewDecoder(r.Body).Decode(&sectorData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrInvalidRequestBody.Error()+err.Error(), http.StatusBadRequest)
			return
		}

		sectorData.ID = sectorId
		updSector, err := sectorService.UpdateSectorById(ctx, sectorData)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrUpdateSector.Error(), zap.Error(err))
			http.Error(w, apperrors.ErrUpdateSector.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully update sector details", http.StatusOK, updSector)
	}
}

func DeleteSectorById(sectorService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		id := vars["sector_id"]
		sectorId, err := strconv.Atoi(id)
		if err != nil {
			logger.Errorw(ctx, apperrors.MsgInvalidSectorId, zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage(apperrors.ErrDeleteSector.Error(), apperrors.MsgInvalidSectorId, id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
			return
		}

		_, err = sectorService.DeleteSectorById(ctx, sectorId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoSectorExists) {
				logger.Errorw(ctx, apperrors.ErrNoSectorExists.Error(), zap.Error(err), zap.String("ID", id))
				middleware.HandleErrorResponse(ctx, w, apperrors.ErrDeleteSector.Error()+", "+err.Error(), http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.ErrDeleteSector.Error(), zap.Error(err), zap.String("ID", id))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrDeleteSector.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func FetchAllSectors(sectorService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sectors, err := sectorService.FetchAllSectors(ctx)
		if err != nil {
			logger.Errorw(ctx, apperrors.ErrFetchSector.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrFetchSector.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}
		middleware.HandleSuccessResponse(ctx, w, "successfully fetched all sectors", http.StatusOK, sectors)
	}
}
