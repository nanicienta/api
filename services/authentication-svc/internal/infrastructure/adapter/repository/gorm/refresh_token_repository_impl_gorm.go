package gorm

import (
	"github.com/nanicienta/api/authentication-svc/internal/ports/outbound/repository"
	"github.com/nanicienta/api/pkg/ports/logging"
	"gorm.io/gorm"
)

type RefreshTokenRepositoryImplGorm struct {
	db     *gorm.DB
	logger logging.Logger
}

// RefreshTokenRepositoryImplGorm is the implementation of RefreshTokenRepository
func NewRefreshTokenRepositoryImpl(
	db *gorm.DB,
	logger logging.Logger,
) repository.RefreshTokenRepository {
	return &RefreshTokenRepositoryImplGorm{
		db:     db,
		logger: logger,
	}
}

func (r *RefreshTokenRepositoryImplGorm) SaveRefreshToken(userID, deviceID, token string) error {
	//TODO implement me
	panic("implement me")
}

func (r *RefreshTokenRepositoryImplGorm) DeleteRefreshTokenByUserIDAndDeviceID(
	userID,
	deviceID string,
) error {
	//TODO implement me
	panic("implement me")
}

func (r *RefreshTokenRepositoryImplGorm) GetRefreshTokenByDeviceIDAndToken(
	deviceID string,
	token string,
) (interface{}, interface{}) {

	//TODO implement me
	panic("implement me")
}
