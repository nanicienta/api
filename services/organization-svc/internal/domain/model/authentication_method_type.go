// Package model provides the domain model for the authentication.
package model

// AuthMethodType represents the type of authentication method used by an external identity.
type AuthMethodType string

const (
	//AuthMethodPassword Traditional password authentication
	AuthMethodPassword AuthMethodType = "password"
	//AuthMethodGoogle Google authentication
	AuthMethodGoogle AuthMethodType = "google"
	//AuthMethodSAML SAML authentication
	AuthMethodSAML AuthMethodType = "saml"
	//AuthMethodOIDC OpenID Connect authentication
	AuthMethodOIDC AuthMethodType = "openid_connect"
)
