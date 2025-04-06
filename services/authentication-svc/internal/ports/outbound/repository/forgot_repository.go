// Package repository provides the interface for the forgot repository
package repository

import "github.com/nanicienta/api/authorization-svc/internal/domain/model"

// ForgotRepository Interface for the forgot repository
type ForgotRepository interface {
	// SaveForgot creates a new forgot code for the given email.
	SaveForgot(forgot model.ForgotModel) (model.ForgotModel, error)
	// FindForgotByCode retrieves the forgot for the given code.
	FindForgotByCode(code string) (model.ForgotModel, error)
	// FindForgotById retrieves the forgot for the given id
	FindForgotById(id string) (model.ForgotModel, error)
}
