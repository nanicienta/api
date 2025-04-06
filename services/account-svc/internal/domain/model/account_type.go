package model

// AccountType represents the type of account (If it's generated on kimos authentication or external authentication)
type AccountType string

const (
	// AccountTypeInternal users generated on kimos authentication. (e.g. email/password)
	AccountTypeInternal AccountType = "INTERNAL"
	// AccountTypeExternal users generated on external authentication. (e.g. google, saml, oidc)
	AccountTypeExternal AccountType = "EXTERNAL"
)
