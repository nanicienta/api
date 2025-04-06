// Package domain provides the domain model for the application.
package domain

// ErrorTypeBody represents the error body for the application.
type ErrorTypeBody struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

// ErrorType represents the error for the application.
type ErrorType struct {
	Body       ErrorTypeBody
	StatusCode int
}
