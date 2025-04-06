package entity

import (
	"time"
)

// ExternalIdentityEntity represents the entity for external identity management.
type ExternalIdentityEntity struct {
	ID         string    `gorm:"column:id;primaryKey"`
	AccountID  string    `gorm:"column:account_id;index"`
	TenantID   string    `gorm:"column:tenant_id;index"`
	Provider   string    `gorm:"column:provider;index"`    // ej: google, saml, oidc
	ExternalID string    `gorm:"column:external_id;index"` // ej: sub o NameID
	Email      string    `gorm:"column:email"`
	RawProfile string    `gorm:"column:raw_profile;type:text"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
}
