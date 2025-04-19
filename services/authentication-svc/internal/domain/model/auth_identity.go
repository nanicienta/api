// Package model provides the domain model for the authentication service.
package model

import "time"

// AuthIdentityModel represents the model for authentication identity (Usually a user)
type AuthIdentityModel struct {
	ID               string    `json:"id"`
	Email            string    `json:"email"`
	Name             string    `json:"name"`
	ProviderID       string    `json:"providerId"`
	IsEnabled        bool      `json:"isEnabled"`
	IsVerified       bool      `json:"isVerified"`
	IsLocked         bool      `json:"isLocked"`
	OrganizationID   string    `json:"organizationId"`
	CreatedAt        time.Time `json:"createdAt"`
	CreatedBy        string    `json:"createdBy"`
	UpdatedAt        time.Time `json:"updatedAt"`
	DeletedAt        time.Time `json:"deletedAt"`
	IsDeleted        bool      `json:"isDeleted"`
	DeletedBy        string    `json:"deletedBy"`
	UpdatedBy        string    `json:"updatedBy"`
	Hash             string
	BadLoginAttempts int
}
