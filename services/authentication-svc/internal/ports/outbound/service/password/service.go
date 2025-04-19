// Package password provides the interface for the password service.
package password

type Service interface {
	// VerifyPassword verifies the given password against the hashed password
	VerifyPassword(hash string, password string) bool
	// HashPassword hashes the given password using bcrypt
	HashPassword(password string) (string, error)
}
