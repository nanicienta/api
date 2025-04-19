// Package model contains the domain models for the authentication service.
package model

import "time"

// RefreshTokenModel represents the opaque refresh‑token issued per device
// and persisted (IsDeleted == soft‑revoked).
type RefreshTokenModel struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	DeviceID  string    `json:"deviceId"`
	Token     string    `json:"token"` // SHA‑256 hash (opaque to callers)
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy,omitempty"`

	DeletedAt time.Time `json:"deletedAt,omitempty"`
	IsDeleted bool      `json:"isDeleted"`
	DeletedBy string    `json:"deletedBy,omitempty"`
}

// IsRevoked checks if the refresh token is revoked (soft deleted).
func (r *RefreshTokenModel) IsRevoked() bool {
	return r.IsDeleted
}
