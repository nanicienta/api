// Package entity provides the entities for the authentication service (
// Connect the model with the DB)
package entity

import "time"

// RefreshTokenEntity represents the entity for the refresh token functionality.
type RefreshTokenEntity struct {
	ID        string    `gorm:"column:id;primaryKey"`
	UserID    string    `gorm:"column:user_id;index"`
	DeviceID  string    `gorm:"column:device_id;index"`
	Token     string    `gorm:"column:token"`
	CreatedBy string    `gorm:"column:created_by"`
	ExpiresAt time.Time `gorm:"column:token"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	DeletedAt time.Time `gorm:"column:deleted_at;index"`
	IsDeleted bool      `gorm:"column:is_deleted"`
	DeletedBy string    `gorm:"column:deleted_by"`
}
