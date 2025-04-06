// Package model represents the domain model for the account service.
package model

import (
	"gorm.io/gorm"
	"time"
)

// Account represents the user account model.
type Account struct {
	ID                string         `json:"id"`
	Type              AccountType    `json:"type"`   // NEW: interno vs externo
	Source            string         `json:"source"` // NEW: manual, google, saml, etc.
	Email             string         `json:"email"`
	EmailVerifiedAt   *time.Time     `json:"emailVerifiedAt"`
	Language          string         `json:"language"`
	Phone             string         `json:"phone"`
	PhoneVerifiedAt   *time.Time     `json:"phoneVerifiedAt"`
	Name              string         `json:"name"`
	PhotoURL          string         `json:"photoUrl"`
	Timezone          string         `json:"timezone"`
	InviteConvertedAt *time.Time     `json:"inviteConvertedAt"`
	MuteSounds        bool           `json:"muteSounds"`
	IsLocked          bool           `json:"isLocked"`
	IsActive          bool           `json:"isActive"`
	CreatedAt         time.Time      `json:"createdAt"`
	CreatedBy         string         `json:"createdBy"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `json:"deletedAt"`
	IsDeleted         bool           `json:"isDeleted"`
	DeletedBy         string         `json:"deletedBy"`
	UpdatedBy         string         `json:"updatedBy"`
	LastLoginAt       *time.Time     `json:"lastLoginAt"`
}
