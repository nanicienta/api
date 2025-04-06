// Package entities defines the entities used by the repository adapter
package entities

import (
	"gorm.io/gorm"
	"time"
)

// OrganizationEntity represents the organization entity in the database.
type OrganizationEntity struct {
	ID             string `gorm:"column:id;primaryKey"`
	Name           string `gorm:"column:name"`
	Subdomain      string `gorm:"column:subdomain"`
	LogoImageURL   string `gorm:"column:logo_image_url"`
	Timezone       string `gorm:"column:timezone"`
	BillingEmail   string `gorm:"column:billing_email"`
	SubscriptionID string `gorm:"column:subscription_id"`
	Plan           string `gorm:"column:plan"`
	IsActive       bool   `gorm:"column:is_active"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	CreatedBy string         `gorm:"column:created_by"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
	IsDeleted bool           `gorm:"column:is_deleted"`
	DeletedBy string         `gorm:"column:deleted_by"`
	UpdatedBy string         `gorm:"column:updated_by"`
}
