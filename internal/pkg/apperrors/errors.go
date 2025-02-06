package apperrors

import "errors"

// common error messages for various different status codes

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidRequestParam = errors.New("invalid request param")
	ErrInvalidRequestBody  = errors.New("invalid request body")
	ErrMarshalPayload      = errors.New("error occured while writing error response")
	ErrWriteHttpResposne   = errors.New("error occured while writing http response")
)
