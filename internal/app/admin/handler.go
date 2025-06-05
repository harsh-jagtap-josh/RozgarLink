package admin

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

func RegisterAdmin(adminS AdminService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req Admin
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Errorw(ctx, apperrors.ErrInvalidRequestBody.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrInvalidRequestBody.Error()+": "+err.Error(), http.StatusBadRequest)
			return
		}

		admin, err := adminS.RegisterAdmin(ctx, req)
		if err != nil {
			if errors.Is(err, apperrors.ErrInvalidUserDetails) {
				logger.Errorw(ctx, apperrors.ErrCreateAdmin.Error(), zap.Error(err))
				middleware.HandleErrorResponse(ctx, w, apperrors.ErrCreateAdmin.Error()+": "+err.Error(), http.StatusBadRequest)
				return
			}

			if errors.Is(err, apperrors.ErrAdminExists) {
				logger.Errorw(ctx, apperrors.ErrAdminExists.Error(), zap.Error(err), zap.String("email: ", req.Email))
				middleware.HandleErrorResponse(ctx, w, apperrors.ErrCreateAdmin.Error()+": "+err.Error()+", email: "+req.Email, http.StatusConflict)
				return
			}

			logger.Errorw(ctx, apperrors.ErrCreateAdmin.Error(), zap.Error(err))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrCreateAdmin.Error()+": "+err.Error(), http.StatusInternalServerError)
			return
		}

		middleware.HandleSuccessResponse(ctx, w, "successfully created a new admin account", http.StatusCreated, admin)
	}
}

func DeleteAdmin(adminS AdminService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		id := vars["admin_id"]
		adminId, err := strconv.Atoi(id)
		if err != nil {
			logger.Errorw(ctx, "invalid admin id provided", zap.Error(err), zap.String("ID", id))
			httpResponseMsg := apperrors.HttpErrorResponseMessage("invalid admin id provided", err.Error(), id)
			middleware.HandleErrorResponse(ctx, w, httpResponseMsg, http.StatusBadRequest)
			return
		}

		err = adminS.DeleteAdmin(ctx, adminId)
		if err != nil {
			if errors.Is(err, apperrors.ErrNoAdminExists) {
				logger.Errorw(ctx, apperrors.ErrDeleteAdmin.Error(), zap.Error(err), zap.String("ID", id))
				middleware.HandleErrorResponse(ctx, w, apperrors.ErrDeleteAdmin.Error()+": "+err.Error(), http.StatusNotFound)
				return
			}

			logger.Errorw(ctx, apperrors.ErrDeleteAdmin.Error(), zap.Error(err), zap.String("ID", id))
			middleware.HandleErrorResponse(ctx, w, apperrors.ErrDeleteAdmin.Error()+", "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
