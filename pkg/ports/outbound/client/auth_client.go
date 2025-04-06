// Package client provides the client interfaces for outbound communication with external services.
package client

import "github.com/nanicienta/api/pkg/domain/auth"

// AuthClient defines the contract for validating tokens with the authentication service.
type AuthClient interface {
	// ValidateToken validates the provided JWT token and returns the claims if valid.
	ValidateToken(token string) (auth.JWTClaim, error)
}
