package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

type SuccessResponse struct {
	SuccessMessage string      `json:"success_message"`
	Data           interface{} `json:"data"`
}

func HandleErrorResponse(w http.ResponseWriter, errMessage string, errStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errStatusCode)

	response := ErrorResponse{
		ErrorMessage: errMessage,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{\"message\":%s}", "error occured while marshaling response payload")))
	}
	_, err = w.Write(jsonData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{\"message\":%s}", "error occured while writing response")))
	}
}

func HandleSuccessResponse(w http.ResponseWriter, successMessage string, StatusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusCode)

	response := SuccessResponse{
		SuccessMessage: successMessage,
		Data:           data,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{\"message\":%s}", "error occured while marshaling success response payload")))
	}
	_, err = w.Write(jsonData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{\"message\":%s}", "error occured while writing success response")))
	}
}
