// Package entities defines the entities used by the repository adapter
package entities

import (
	"github.com/nanicienta/api/organization-svc/internal/domain/model"
	"gorm.io/gorm"
	"time"
)

// OrganizationAuthenticationMethodEntity represents the authentication method entity for an organization.
type OrganizationAuthenticationMethodEntity struct {
	ID              string               `gorm:"column:id;primaryKey"`
	OrganizationID  string               `gorm:"column:organization_id;index"`
	MethodType      model.AuthMethodType `gorm:"column:method_type"`
	DisplayName     string               `gorm:"column:display_name"`
	IsEnabled       bool                 `gorm:"column:is_enabled"`
	ClientID        string               `gorm:"column:client_id"`
	ClientSecret    string               `gorm:"column:client_secret"`
	IssuerURL       string               `gorm:"column:issuer_url"`
	RedirectURL     string               `gorm:"column:redirect_url"`
	SAMLMetadataURL string               `gorm:"column:saml_metadata_url"`
	CreatedAt       time.Time            `gorm:"column:created_at;autoCreateTime"`
	CreatedBy       string               `gorm:"column:created_by"`
	UpdatedAt       time.Time            `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt       gorm.DeletedAt       `gorm:"column:deleted_at;index"`
	IsDeleted       bool                 `gorm:"column:is_deleted"`
	DeletedBy       string               `gorm:"column:deleted_by"`
	UpdatedBy       string               `gorm:"column:updated_by"`
}
