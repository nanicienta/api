// Package auth contains the logic for validating JWT tokens and handle authentication
package auth

import (
	"github.com/nanicienta/api/pkg/domain/auth"
)

// ValidateToken validates the JWT token and returns the claims
func ValidateToken(signedToken string) (claims *auth.JWTClaim, err error) {
	//TODO implement the client to go to the authorization service and check if the token is valid
	return &auth.JWTClaim{
		Email:        "",
		ID:           "",
		Organization: "",
		ExpiresAt:    0,
		MapClaims:    nil,
	}, nil
}
