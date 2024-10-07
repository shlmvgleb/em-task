package utils

import (
	"errors"
)

var (
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidResponse    = errors.New("invalid response from service")
	ErrForbidden          = errors.New("forbidden")
)

var (
// TODO: error messages
)
