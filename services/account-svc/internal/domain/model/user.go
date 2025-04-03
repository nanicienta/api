package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                string         `json:"id"`
	Hash              string         `json:"hash"`
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	EmailVerifiedAt   *time.Time     `json:"emailVerifiedAt"`
	Phone             string         `json:"phone"`
	PhoneVerifiedAt   *time.Time     `json:"phoneVerifiedAt"`
	PhotoUrl          string         `json:"photoUrl"`
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
	AccessToken       string         `json:"accessToken"`
	RefreshToken      string         `json:"refreshToken"`
	LastLoginAt       *time.Time     `json:"lastLoginAt"`
}
