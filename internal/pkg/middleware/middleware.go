package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/logger"
	"go.uber.org/zap"
)

type ErrorResponse struct {
	ErrorMessage string `json:"message"`
}

type SuccessResponse struct {
	SuccessMessage string      `json:"message"`
	Data           interface{} `json:"data"`
}

func HandleErrorResponse(ctx context.Context, w http.ResponseWriter, errMessage string, errStatusCode int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errStatusCode)
	response := ErrorResponse{
		ErrorMessage: errMessage,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		logger.Errorw(ctx, apperrors.ErrMarshalPayload.Error(), zap.Error(err))
		w.Write(HttpErrorResponseMessages(w, apperrors.ErrMarshalPayload.Error(), http.StatusInternalServerError))
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		logger.Errorw(ctx, "error occured while writing http error response", zap.Error(err))
		w.Write(HttpErrorResponseMessages(w, apperrors.ErrWriteHttpResposne.Error(), http.StatusInternalServerError))
		return
	}
}

func HandleSuccessResponse(ctx context.Context, w http.ResponseWriter, successMessage string, StatusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusCode)

	response := SuccessResponse{
		SuccessMessage: successMessage,
		Data:           data,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		logger.Errorw(ctx, "error occured while marshaling response payload", zap.Error(err))
		w.Write(HttpErrorResponseMessages(w, apperrors.ErrMarshalPayload.Error(), http.StatusInternalServerError))
	}

	_, err = w.Write(jsonData)
	if err != nil {
		logger.Errorw(ctx, apperrors.ErrMarshalPayload.Error(), zap.Error(err))
		w.Write(HttpErrorResponseMessages(w, apperrors.ErrWriteHttpResposne.Error(), http.StatusInternalServerError))
	}
}

func HttpErrorResponseMessages(w http.ResponseWriter, message string, statusCode int) []byte {
	w.WriteHeader(statusCode)
	msg := fmt.Sprintf("message: %s", message)
	return []byte(msg)
}
