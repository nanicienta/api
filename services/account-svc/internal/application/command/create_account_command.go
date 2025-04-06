// Package command provides the command structure for different actions on the service.
package command

import "github.com/nanicienta/api/account-svc/internal/domain/model"

// CreateAccountCommand represents the command to create a new account.
type CreateAccountCommand struct {
	Email  string            `json:"email"`
	Name   string            `json:"name"`
	Type   model.AccountType `json:"type"`
	Source string            `json:"source"`
}
