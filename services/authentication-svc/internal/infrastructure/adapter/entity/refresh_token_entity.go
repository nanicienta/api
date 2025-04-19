// Package entity provides the entities for the authentication service (
// Connect the model with the DB)
package entity

// RefreshTokenEntity represents the entity for the refresh token functionality.
type RefreshTokenEntity struct {
	ID        string `gorm:"column:id;primaryKey"`
	UserID    string `gorm:"column:user_id;index"`
	DeviceID  string `gorm:"column:device_id;index"`
	Token     string `gorm:"column:token"`
	CreatedAt string `gorm:"column:created_at;autoCreateTime"`
	CreatedBy string `gorm:"column:created_by"`
	DeletedAt string `gorm:"column:deleted_at;index"`
	IsDeleted bool   `gorm:"column:is_deleted"`
	DeletedBy string `gorm:"column:deleted_by"`
}
