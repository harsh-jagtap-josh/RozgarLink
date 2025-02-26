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
var ErrNoWorkerExists = errors.New("no worker found with id")

// Workers Error Messages
const MsgConvertIdToInt = "error occurred while converting id to number"
const MsgInvalidWorkerId = "invalid worker id provided"
const MsgFetchFromDb = "error occurred while fetching from database"
const MsgFailedToFetchWorker = "failed to fetch worker"

func HttpErrorResponseMessage(warning, message string, id string) string {
	return fmt.Sprintf("%s: %s, id: %v", warning, message, id)
}
