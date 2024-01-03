package utils

import (
	"net/http"
)

type APIError struct {
	Code    int
	Message string
}

var (
	ErrInternalServer = APIError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}

	ErrDataNotFound = APIError{
		Code:    http.StatusNotFound,
		Message: "Data Not Found",
	}

	ErrInvalidInput = APIError{
		Code:    http.StatusBadRequest,
		Message: "Invalid input",
	}

	ErrBadRequest = APIError{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}

	ErrUserExists = APIError{
		Code:    http.StatusBadRequest,
		Message: "Student already exists",
	}

	ErrEnrollExists = APIError{
		Code:    http.StatusBadRequest,
		Message: "Already enrolled to course",
	}

	ErrUnauthorized = APIError{
		Code:    http.StatusUnauthorized,
		Message: "Authorization failed",
	}
)
