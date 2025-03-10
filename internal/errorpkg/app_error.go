package errorpkg

import (
	"errors"
	"fmt"
)

var ParseJsonError = errors.New("failed to parse json")
var ValidateFieldsError = errors.New("failed to validate fields")
var NotFoundError = errors.New("not found")

// AppError is a custom error type for consistent error handling.
type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Cause      error  `json:"-"`
}

// Error implements the error interface.
func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func New(message string, statusCode int, cause error) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
		Cause:      cause,
	}
}

func Wrap(message string, statusCode int, cause error) *AppError {
	return New(message, statusCode, cause)
}

// IsAppError checks if an error is of type AppError.
func IsAppError(err error) bool {
	var appError *AppError
	ok := errors.As(err, &appError)
	return ok
}
