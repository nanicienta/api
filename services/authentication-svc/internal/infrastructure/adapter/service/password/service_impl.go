package password

import "golang.org/x/crypto/bcrypt"

// ServiceImpl implements the password service interface
type ServiceImpl struct {
}

// VerifyPassword verifies the password against the hash
func (s *ServiceImpl) VerifyPassword(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
