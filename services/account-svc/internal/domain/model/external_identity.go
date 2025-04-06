package model

import (
	"time"
)

// ExternalIdentity represents an external identity model in the database.
type ExternalIdentity struct {
	ID         string    `json:"id"`
	AccountID  string    `json:"accountId"`
	TenantID   string    `json:"tenantId"`
	Provider   string    `json:"provider"`   // ej: google, saml, oidc
	ExternalID string    `json:"externalId"` // ej: sub o NameID
	Email      string    `json:"email"`
	RawProfile string    `json:"rawProfile"`
	CreatedAt  time.Time `json:"createdAt"`
}
