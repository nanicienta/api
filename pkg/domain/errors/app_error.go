// Package errors provides the error types and constants for the application.
package errors

import (
	"fmt"
	"net/http"
)

// AppError represents the error code.
type AppError struct {
	Message     string `json:"message"`
	Code        string `json:"code"`
	Description string `json:"description"`
	HTTPStatus  int
}

// Error returns the error message.
func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NotFoundError represents an errors when a resource is not found.
type NotFoundError struct {
	*AppError
}

// NewNotFoundError creates a new NotFoundError with the given message, description, and error code.
func NewNotFoundError(message, description string, code ErrorCode) *NotFoundError {
	return &NotFoundError{
		AppError: &AppError{
			Message:     message,
			Code:        string(code),
			Description: description,
			HTTPStatus:  http.StatusNotFound,
		},
	}
}

// UnauthorizedError represents an error when a user is not authorized to access a resource.
type UnauthorizedError struct {
	*AppError
}

// NewUnauthorizedError creates a new UnauthorizedError with the given message, description, and error code.
func NewUnauthorizedError(message, description string, code ErrorCode) *UnauthorizedError {
	return &UnauthorizedError{
		AppError: &AppError{
			Message:     message,
			Code:        string(code),
			Description: description,
			HTTPStatus:  http.StatusUnauthorized,
		},
	}
}

// BadRequestError represents an error when the request is invalid.
type BadRequestError struct {
	*AppError
}

// NewBadRequestError creates a new BadRequestError with the given message, description, and error code.
func NewBadRequestError(message, description string, code ErrorCode) *BadRequestError {
	return &BadRequestError{
		AppError: &AppError{
			Message:     message,
			Code:        string(code),
			Description: description,
			HTTPStatus:  http.StatusBadRequest,
		},
	}
}

// InternalServerError represents an error when the server encounters an unexpected condition.
type InternalServerError struct {
	*AppError
}

// NewInternalServerError creates a new InternalServerError with the given message, description, and error code.
func NewInternalServerError(message, description string, code ErrorCode) *InternalServerError {
	return &InternalServerError{
		AppError: &AppError{
			Message:     message,
			Code:        string(code),
			Description: description,
			HTTPStatus:  http.StatusInternalServerError,
		},
	}
}

// ForbiddenError represents an error when a user is not authorized to access a resource.
type ForbiddenError struct {
	*AppError
}

// NewForbiddenError creates a new ForbiddenError with the given message, description, and error code.
func NewForbiddenError(message, description string, code ErrorCode) *ForbiddenError {
	return &ForbiddenError{
		AppError: &AppError{
			Message:     message,
			Code:        string(code),
			Description: description,
			HTTPStatus:  http.StatusForbidden,
		},
	}
}

// ConflictError represents an error when a resource already exists.
type ConflictError struct {
	*AppError
}

// NewConflictError creates a new ConflictError with the given message, description, and error code.
func NewConflictError(message, description string, code ErrorCode) *ConflictError {
	return &ConflictError{
		AppError: &AppError{
			Message:     message,
			Code:        string(code),
			Description: description,
			HTTPStatus:  http.StatusConflict,
		},
	}
}
