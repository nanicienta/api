package entities

import (
	"github.com/nanicienta/api/account-svc/internal/domain/model"
	"gorm.io/gorm"
	"time"
)

type UserEntity struct {
	ID                string         `json:"id" gorm:"column:id;primaryKey"`
	Hash              string         `json:"hash" gorm:"column:hash"`
	Name              string         `json:"name" gorm:"column:name"`
	Email             string         `json:"email" gorm:"column:email"`
	EmailVerifiedAt   *time.Time     `json:"email_verified_at" gorm:"column:email_verified_at"`
	Phone             string         `json:"phone" gorm:"column:phone"`
	PhoneVerifiedAt   *time.Time     `json:"phone_verified_at" gorm:"column:phone_verified_at"`
	PhotoUrl          string         `json:"photo_url" gorm:"column:photo_url"`
	Timezone          string         `json:"timezone" gorm:"column:timezone"`
	InviteConvertedAt *time.Time     `json:"invite_converted_at" gorm:"column:invite_converted_at"`
	MuteSounds        bool           `json:"mute_sounds" gorm:"column:mute_sounds"`
	IsLocked          bool           `json:"is_locked" gorm:"column:is_locked"`
	IsActive          bool           `json:"is_active" gorm:"column:is_active"`
	CreatedAt         time.Time      `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	CreatedBy         string         `json:"created_by" gorm:"column:created_by"`
	UpdatedAt         time.Time      `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;index"`
	IsDeleted         bool           `json:"is_deleted" gorm:"column:is_deleted"`
	DeletedBy         string         `json:"deleted_by" gorm:"column:deleted_by"`
	UpdatedBy         string         `json:"updated_by" gorm:"column:updated_by"`
	LastLoginAt       *time.Time     `json:"last_login_at" gorm:"column:last_login_at"`
}

func (e UserEntity) ToModel() model.User {
	return model.User{
		ID:                e.ID,
		Hash:              e.Hash,
		Name:              e.Name,
		Email:             e.Email,
		EmailVerifiedAt:   e.EmailVerifiedAt,
		Phone:             e.Phone,
		PhoneVerifiedAt:   e.PhoneVerifiedAt,
		PhotoUrl:          e.PhotoUrl,
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
