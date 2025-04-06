// Package entity provides the entities for the authentication service (
// Connect the model with the DB)
package entity

import (
	"gorm.io/gorm"
	"time"
)

// ForgotEntity represents the entity for the forgot password functionality.
type ForgotEntity struct {
	ID        string         `gorm:"column:id;primaryKey"`
	Code      string         `gorm:"column:code"`
	UserID    string         `gorm:"column:user_id;index"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	CreatedBy string         `gorm:"column:created_by"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
	IsDeleted bool           `gorm:"column:is_deleted"`
	DeletedBy string         `gorm:"column:deleted_by"`
	UpdatedBy string         `gorm:"column:updated_by"`
}
