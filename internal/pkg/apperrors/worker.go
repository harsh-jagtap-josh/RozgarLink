package apperrors

import "fmt"

const ErrNoWorkerExists = "no worker found with id"
const ErrConvertIdToInt = "error occurred while converting id to number"
const ErrFetchFromDb = "error occurred while fetching from database"
const ErrEmailAlready = "email already exists"
const ErrFailedToFetchWorker = "failed to fetch worker"
const ErrFailedToCreateWorker = "failed to create user"
const ErrFailedToUpdateWorker = "failed to update user"
const ErrFailedToDeleteWorker = "failed to update user"

func HttpErrorResponseMessage(warning, message string, id string) string {
	return fmt.Sprintf("%s: %s, id: %v", warning, message, id)
}
