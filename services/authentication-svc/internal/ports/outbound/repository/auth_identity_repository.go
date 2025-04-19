// Package repository provides the interface for the forgot repository
package repository

import "github.com/nanicienta/api/authentication-svc/internal/domain/model"

// AuthIdentityRepository is the interface for the authentication identity repository
type AuthIdentityRepository interface {
	// GetByEmailAndOrganizationID retrieves an identity by email and organization ID
	GetByEmailAndOrganizationID(email, organizationID string) (*model.AuthIdentityModel, error)
	// UpdateBadAttemptCount updates the bad attempt count for an identity
	UpdateBadAttemptCount(authModel *model.AuthIdentityModel) error
}
