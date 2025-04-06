// Package model provides the data structures for the authentication service
package model

import "time"

// ForgotModel represents the model for a password reset request
type ForgotModel struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	IsDeleted bool      `json:"isDeleted"`
	DeletedBy string    `json:"deletedBy"`
	UpdatedBy string    `json:"updatedBy"`
}
