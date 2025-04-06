// Package entities defines the entities used by the repository adapter
package entities

import (
	"gorm.io/gorm"
	"time"
)

// OrganizationAuthenticationConfigurationEntity represents the organization authentication configuration entity in the database.
type OrganizationAuthenticationConfigurationEntity struct {
	ID                             string         `gorm:"column:id;primaryKey"`
	OrganizationID                 string         `gorm:"column:organization_id;index"`
	PasswordAuthenticationEnabled  bool           `gorm:"column:password_authentication_enabled"`
	TwoFactorAuthenticationEnabled bool           `gorm:"column:two_factor_authentication_enabled"`
	CreatedAt                      time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt                      time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt                      gorm.DeletedAt `gorm:"column:deleted_at;index"`
	IsDeleted                      bool           `gorm:"column:is_deleted"`
	DeletedBy                      string         `gorm:"column:deleted_by"`
	UpdatedBy                      string         `gorm:"column:updated_by"`
}
