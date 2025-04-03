package outbound

import (
	"github.com/nanicienta/api/account-svc/internal/domain/model"
	"github.com/nanicienta/api/pkg/domain"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetPage(pageNumber int, pageSize int) (domain.Page[model.User], error)
	GetByID(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id string) error
}
