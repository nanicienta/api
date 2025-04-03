package entities

import "time"

type ForgotEntity struct {
	ID        string    `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
	IsDeleted bool      `json:"is_deleted" db:"is_deleted"`
	DeletedBy string    `json:"deleted_by" db:"deleted_by"`
	UpdatedBy string    `json:"updated_by" db:"updated_by"`
}
