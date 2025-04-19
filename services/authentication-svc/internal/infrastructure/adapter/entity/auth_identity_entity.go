// Package entity provides the entities for the authentication service (
// Connect the model with the DB)
package entity

import (
	"github.com/nanicienta/api/authentication-svc/internal/domain/model"
	"gorm.io/gorm"
	"time"
)

// AuthIdentityEntity represents the entity for the user.
type AuthIdentityEntity struct {
	ID               string `gorm:"column:id;primaryKey"`
	Email            string `gorm:"column:email;index"`
	ProviderID       string `gorm:"column:provider_id;index"`
	Name             string `gorm:"column:name"`
	Hash             string `gorm:"column:hash"`
	IsEnabled        bool   `gorm:"column:is_enabled"`
	IsVerified       bool   `gorm:"column:is_verified"`
	IsLocked         bool   `gorm:"column:is_locked"`
	BadLoginAttempts int    `gorm:"column:bad_login_attempts"`

	OrganizationID string `gorm:"column:organization_id;index"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	CreatedBy string         `gorm:"column:created_by"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
	IsDeleted bool           `gorm:"column:is_deleted"`
	DeletedBy string         `gorm:"column:deleted_by"`
	UpdatedBy string         `gorm:"column:updated_by"`
}

// ToModel converts the AuthIdentityEntity to model.AuthIdentityModel.
func (e AuthIdentityEntity) ToModel() model.AuthIdentityModel {
	return model.AuthIdentityModel{
		ID:               e.ID,
		Email:            e.Email,
		Name:             e.Name,
		Hash:             e.Hash,
		IsEnabled:        e.IsEnabled,
		IsVerified:       e.IsVerified,
		IsLocked:         e.IsLocked,
		BadLoginAttempts: e.BadLoginAttempts,
		ProviderID:       e.ProviderID,
		OrganizationID:   e.OrganizationID,

		CreatedAt: e.CreatedAt,
		CreatedBy: e.CreatedBy,
		UpdatedAt: e.UpdatedAt,
		IsDeleted: e.IsDeleted,
		DeletedBy: e.DeletedBy,
		UpdatedBy: e.UpdatedBy,
	}
}
