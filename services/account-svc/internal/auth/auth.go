// Package auth contains the logic for validating JWT tokens and handle authentication
package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nanicienta/api/pkg/domain/auth"
	"os"
	"time"
)

// ValidateToken validates the JWT token and returns the claims
func ValidateToken(signedToken string) (claims *auth.JWTClaim, err error) {
	jwtKey := os.Getenv("JWT_KEY")
	token, err := jwt.ParseWithClaims(
		signedToken,
		&auth.JWTClaim{},
		//TODO check what we are doing here
		func(_ *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*auth.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if (claims.ExpiresAt) < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
