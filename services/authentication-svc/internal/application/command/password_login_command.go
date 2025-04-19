// Package command provides the command structure for different actions on the service.
package command

// PasswordLoginCommand represents the command to log in a user using their email and password.
type PasswordLoginCommand struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	OrganizationID string
	DeviceID       string
}
