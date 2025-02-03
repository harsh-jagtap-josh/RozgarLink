package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpErrorResponse struct {
	ErrorMessage string `json:"error_message"`
	// errorStatusCode int    `json:"error_code"`
	// data            interface{}    `json:"data"`
}

type httpSuccessResponse struct {
	SuccessMessage string      `json:"success_message"`
	data           interface{} `json:"data"`
}

func HandleErrorResponse(w http.ResponseWriter, errMessage string, errStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errStatusCode)

	response := httpErrorResponse{
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

	response := httpSuccessResponse{
		SuccessMessage: successMessage,
		data:           data,
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
