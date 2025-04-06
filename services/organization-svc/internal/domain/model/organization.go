package model

import (
	"gorm.io/gorm"
	"time"
)

// Organization represents the organization model in the database.
type Organization struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Subdomain      string `json:"subdomain"`
	LogoImageURL   string `json:"logoImageUrl"`
	Timezone       string `json:"timezone"`
	BillingEmail   string `json:"billingEmail"`
	SubscriptionID string `json:"subscriptionId"`
	Plan           string `json:"plan"`
	IsActive       bool   `json:"isActive"`

	CreatedAt time.Time      `json:"createdAt"`
	CreatedBy string         `json:"createdBy"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	IsDeleted bool           `json:"isDeleted"`
	DeletedBy string         `json:"deletedBy"`
	UpdatedBy string         `json:"updatedBy"`
}
