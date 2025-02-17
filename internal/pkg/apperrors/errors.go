package apperrors

import (
	"errors"
	"fmt"
)

// common error messages for various different status codes

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidRequestParam = errors.New("invalid request param")
	ErrInvalidRequestBody  = errors.New("missing or invalid fields in request")
	ErrMarshalPayload      = errors.New("error occured while writing error response")
	ErrWriteHttpResposne   = errors.New("error occured while writing http response")
	ErrEncrPassword        = errors.New("error occured while hashing password")
	ErrCreateToken         = errors.New("failed to create jwt token")
	ErrFailedLogin         = errors.New("failed to login user")

	// Worker/User/Employer Errors
	ErrCreateWorker        = errors.New("failed to create worker")
	ErrUpdateWorker        = errors.New("failed to update worker data")
	ErrDeleteWorker        = errors.New("failed to delete worker data")
	ErrCreateAddress       = errors.New("error occured while creating address")
	ErrInvalidUserDetails  = errors.New("invalid user details")
	ErrNoWorkerExists      = errors.New("no worker found with id")
	ErrWorkerAlreadyExists = errors.New("worker with same email already exists")

	ErrNoEmployerExists      = errors.New("no employer found with id")
	ErrCreateEmployer        = errors.New("failed to create employer")
	ErrUpdateEmployer        = errors.New("failed to update employer data")
	ErrDeleteEmployer        = errors.New("failed to delete employer data")
	ErrEmployerAlreadyExists = errors.New("employer with same email already exists")

	// Job Errors
	ErrCreateJob   = errors.New("failed to create job")
	ErrUpdateJob   = errors.New("failed to update job data")
	ErrDeleteJob   = errors.New("failed to delete job data")
	ErrFetchJob    = errors.New("failed to fetch job data")
	ErrNoJobExists = errors.New("no job found with id")

	// Login Errors
	ErrInvalidLoginCredentials = errors.New("invalid email or password")
)

// Workers Error Messages
const MsgConvertIdToInt = "error occurred while converting id to number"
const MsgInvalidWorkerId = "invalid worker id provided"
const MsgFetchFromDb = "error occurred while fetching from database"
const MsgFailedToFetchWorker = "failed to fetch worker"
const MsgFailedToCreateWorker = "failed to create worker"

// Employer Error Messages
const MsgInvalidEmployerId = "invalid employer id provided"
const MsgFailedToFetchEmp = "failed to fetch employer"

// Job Error Messages
const MsgInvalidJobId = "invalid job id provided"
const MsgFailedToFetchJob = "failed to fetch job"

func HttpErrorResponseMessage(warning, message string, id string) string {
	return fmt.Sprintf("%s: %s, id: %v", warning, message, id)
}
