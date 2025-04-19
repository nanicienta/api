// Package token provides the interface for the token service.
package token

import (
	"github.com/nanicienta/api/authentication-svc/internal/domain/model"
)

// Service is the interface for the token service
type Service interface {
	Issue(id *model.AuthIdentityModel, deviceID string) (*model.TokensModel, error)
	VerifyRefresh(refresh string) ([]byte, error)
	PublicJWK() string
}
