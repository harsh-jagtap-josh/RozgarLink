package apperrors

import (
	"errors"
	"fmt"
)

// common error messages for various different status codes

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidRequestParam = errors.New("invalid request param")
	ErrInvalidRequestBody  = errors.New("invalid request body")
	ErrMarshalPayload      = errors.New("error occured while writing error response")
	ErrWriteHttpResposne   = errors.New("error occured while writing http response")
)

// Worker Errors

const ErrNoWorkerExists = "no worker found with id"
const ErrConvertIdToInt = "error occurred while converting id to number"
const ErrInvalidWorkerId = "invalid worker id provided"
const ErrFetchFromDb = "error occurred while fetching from database"
const ErrEmailAlready = "email already exists"
const ErrFailedToFetchWorker = "failed to fetch worker"
const ErrFailedToCreateWorker = "failed to create worker"
const ErrFailedToUpdateWorker = "failed to update worker"
const ErrFailedToDeleteWorker = "failed to update worker"

func HttpErrorResponseMessage(warning, message string, id string) string {
	return fmt.Sprintf("%s: %s, id: %v", warning, message, id)
}
