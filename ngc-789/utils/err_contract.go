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
		Code:    http.StatusOK,
		Message: "Data Not Found",
	}

	ErrBadRequest = APIError{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}

	ErrUnauthorized = APIError{
		Code:    http.StatusUnauthorized,
		Message: "Request Unauthorized",
	}
)

// untuk swaggo

type APIErrorResponse struct {
	Error APIError
}
