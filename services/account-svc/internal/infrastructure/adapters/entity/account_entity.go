// Package entity contains the database entities for the account service
package entity

import (
	"github.com/nanicienta/api/account-svc/internal/domain/model"
	"gorm.io/gorm"
	"time"
)

// AccountEntity represents the account entity in the database
type AccountEntity struct {
	ID                string            `gorm:"column:id;primaryKey"`
	Type              model.AccountType `gorm:"column:type"`   // internal vs external
	Source            string            `gorm:"column:source"` // manual, google, saml, etc.
	Email             string            `gorm:"column:email;uniqueIndex"`
	EmailVerifiedAt   *time.Time        `gorm:"column:email_verified_at"`
	Language          string            `gorm:"column:language;"`
	Phone             string            `gorm:"column:phone"`
	PhoneVerifiedAt   *time.Time        `gorm:"column:phone_verified_at"`
	Name              string            `gorm:"column:name"`
	PhotoURL          string            `gorm:"column:photo_url"`
	Timezone          string            `gorm:"column:timezone"`
	InviteConvertedAt *time.Time        `gorm:"column:invite_converted_at"`
	MuteSounds        bool              `gorm:"column:mute_sounds"`
	IsLocked          bool              `gorm:"column:is_locked"`
	IsActive          bool              `gorm:"column:is_active"`
	CreatedAt         time.Time         `gorm:"column:created_at;autoCreateTime"`
	CreatedBy         string            `gorm:"column:created_by"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt         gorm.DeletedAt    `gorm:"column:deleted_at;index"` // Indice para borrado lógico
	IsDeleted         bool              `gorm:"column:is_deleted"`
	DeletedBy         string            `gorm:"column:deleted_by"`
	UpdatedBy         string            `gorm:"column:updated_by"`
	LastLoginAt       *time.Time        `gorm:"column:last_login_at"`
}

// ToModel transform the entity to a model
func (e AccountEntity) ToModel() model.Account {
	return model.Account{
		ID:                e.ID,
		Name:              e.Name,
		Email:             e.Email,
		EmailVerifiedAt:   e.EmailVerifiedAt,
		Phone:             e.Phone,
		PhoneVerifiedAt:   e.PhoneVerifiedAt,
		PhotoURL:          e.PhotoURL,
		Timezone:          e.Timezone,
		InviteConvertedAt: e.InviteConvertedAt,
		MuteSounds:        e.MuteSounds,
		IsLocked:          e.IsLocked,
		IsActive:          e.IsActive,
		CreatedAt:         e.CreatedAt,
		CreatedBy:         e.CreatedBy,
		UpdatedAt:         e.UpdatedAt,
		DeletedAt:         e.DeletedAt,
		IsDeleted:         e.IsDeleted,
		DeletedBy:         e.DeletedBy,
		UpdatedBy:         e.UpdatedBy,
		LastLoginAt:       e.LastLoginAt,
	}
}
