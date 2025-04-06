package auth

import "github.com/golang-jwt/jwt/v5"

// JWTClaim represents the JWT claims used for authentication
type JWTClaim struct {
	Email        string `json:"email"`
	ID           string `json:"id"`
	Organization string `json:"organization"`
	ExpiresAt    int64  `json:"exp"`
	jwt.MapClaims
}
