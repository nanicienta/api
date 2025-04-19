package gorm

import (
	"errors"
	"github.com/nanicienta/api/authentication-svc/internal/domain/model"
	"github.com/nanicienta/api/authentication-svc/internal/infrastructure/adapter/entity"
	"github.com/nanicienta/api/authentication-svc/internal/ports/outbound/repository"
	"github.com/nanicienta/api/pkg/ports/logging"
	"gorm.io/gorm"
)

type AuthIdentityRepositoryImplGorm struct {
	logger logging.Logger
	db     *gorm.DB
}

// NewAuthIdentityRepository creates a new instance of AuthIdentityRepository
func NewAuthIdentityRepository(
	logger logging.Logger,
	db *gorm.DB,
) repository.AuthIdentityRepository {
	return &AuthIdentityRepositoryImplGorm{
		logger: logger,
		db:     db,
	}
}

// GetByEmailAndOrganizationID retrieves an identity by email and organization ID
func (a *AuthIdentityRepositoryImplGorm) GetByEmailAndOrganizationID(email, organizationID string) (
	*model.AuthIdentityModel,
	error,
) {
	var authIdentityEntity entity.AuthIdentityEntity
	if err := a.db.Where(
		"email = ? AND organization_id = ?",
		email,
		organizationID,
	).First(&authIdentityEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		a.logger.Error("failed to get auth identity", "error", err)
		return nil, err
	}
	authIdentityModel := authIdentityEntity.ToModel()
	return &authIdentityModel, nil
}

// UpdateBadAttemptCount updates the bad attempt count for an identity
func (a *AuthIdentityRepositoryImplGorm) UpdateBadAttemptCount(
	authModel *model.
		AuthIdentityModel,
) error {
	tx := a.db.Where("id = ?", authModel.ID).Updates(
		map[string]interface{}{
			"bad_login_attempts": authModel.BadLoginAttempts,
			"is_locked":          authModel.IsLocked,
		},
	)
	if tx.Error != nil {
		a.logger.Error("failed to update bad attempt count", "error", tx.Error)
		return tx.Error
	}
	return nil
}
