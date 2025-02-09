package apperrors

import "fmt"

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
