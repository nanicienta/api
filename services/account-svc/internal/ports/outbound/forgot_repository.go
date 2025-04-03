package outbound

import (
	"github.com/nanicienta/api/account-svc/internal/infrastructure/adapters/entities"
)

type ForgotRepository interface {
	GetByID(id string) (*entities.UserEntity, error)
	Create(user *entities.UserEntity) (*entities.UserEntity, error)
	Update(user *entities.UserEntity) (*entities.UserEntity, error)
	Delete(id string) error
}
