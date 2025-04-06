// Package repository defines the interface for the account repositories
package repository

import (
	"github.com/nanicienta/api/account-svc/internal/domain/model"
	"github.com/nanicienta/api/pkg/domain"
)

// AccountRepository Interface for the account repository
type AccountRepository interface {
	GetAll() ([]model.Account, error)
	GetPage(pageNumber int, pageSize int) (domain.Page[model.Account], error)
	GetByID(id string) (*model.Account, error)
	GetByEmail(email string) (*model.Account, error)
	Create(account *model.Account) (*model.Account, error)
	Update(user *model.Account) (*model.Account, error)
	Delete(id string) error
}
