package errors

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrExternalAPI  = errors.New("external API error")
)

type APIError struct {
	Code    int
	Message string
	Err     error
}

func (e *APIError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *APIError) Unwrap() error {
	return e.Err
}

func NewNotFoundError(resource string) *APIError {
	return &APIError{
		Code:    404,
		Message: fmt.Sprintf("%s not found", resource),
		Err:     ErrNotFound,
	}
}

func NewValidationError(msg string) *APIError {
	return &APIError{
		Code:    400,
		Message: msg,
		Err:     ErrInvalidInput,
	}
}

func NewExternalError(service string, err error) *APIError {
	return &APIError{
		Code:    502,
		Message: fmt.Sprintf("%s service error", service),
		Err:     fmt.Errorf("%w: %v", ErrExternalAPI, err),
	}
}
