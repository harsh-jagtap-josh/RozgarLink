package admin

import (
	"encoding/json"
	"errors"
	"net/http"

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
