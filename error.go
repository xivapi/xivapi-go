package xivapi

import (
	"errors"
	"fmt"
)

// StatusCodeError is thrown if an unexpexted http status code is returned
type StatusCodeError struct {
	StatusCode int
	message    string
}

// NewStatusCodeError creates a new StatusCodeError
func NewStatusCodeError(code int) *StatusCodeError {
	return &StatusCodeError{
		StatusCode: code,
		message:    "Unexpected Status Code found: %v",
	}
}

func (e *StatusCodeError) Error() string {
	return fmt.Sprintf(e.message, e.StatusCode)
}

// IsStatusCodeError checks if a given error is a status code error
// A small convinience wrapper around the regular Go Type Assertion
func IsStatusCodeError(err error) bool {
	_, ok := err.(*StatusCodeError)

	return ok
}

// The different error types of xivapi
var (
	ErrUnexpectedType = errors.New("unexpected type")
	ErrNotImplemented = errors.New("call not implemented on the server")
)
