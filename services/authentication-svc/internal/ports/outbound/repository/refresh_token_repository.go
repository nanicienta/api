// Package repository provides the interface for the forgot repository
package repository

// RefreshTokenRepository provides the interface for the refresh token repository
type RefreshTokenRepository interface {
	// SaveRefreshToken saves the refresh token for the given user ID and device ID.
	SaveRefreshToken(userID, deviceID, token string) error

	// DeleteRefreshTokenByUserIDAndDeviceID deletes the refresh token for the given user ID and device ID.
	DeleteRefreshTokenByUserIDAndDeviceID(userID, deviceID string) error
	// GetRefreshTokenByDeviceIDAndToken retrieves the refresh token for the given user ID and device ID.
	GetRefreshTokenByDeviceIDAndToken(deviceID string, token string) (interface{}, interface{})
}
