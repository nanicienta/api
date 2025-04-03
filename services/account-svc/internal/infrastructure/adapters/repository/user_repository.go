package repository

import (
	"github.com/nanicienta/api/account-svc/internal/domain/model"
	"github.com/nanicienta/api/account-svc/internal/infrastructure/adapters/entities"
	"github.com/nanicienta/api/account-svc/internal/ports/outbound"
	"github.com/nanicienta/api/pkg/domain"
	"github.com/nanicienta/api/pkg/ports/logging"
	"gorm.io/gorm"
)

func NewUserRepository(logger logging.Logger, db *gorm.DB) outbound.UserRepository {
	return &UserRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

type UserRepositoryImpl struct {
	logger logging.Logger
	db     *gorm.DB
}

func (r *UserRepositoryImpl) GetAll() ([]model.User, error) {
	var userEntities []entities.UserEntity
	if err := r.db.Find(&userEntities).Error; err != nil {
		r.logger.Error("Error getting users", "error", err)
		return nil, err
	}
	users := make([]model.User, len(userEntities))
	for i, userEntity := range userEntities {
		users[i] = userEntity.ToModel()
	}
	return users, nil
}

func (r *UserRepositoryImpl) GetPage(pageNumber int, pageSize int) (
	domain.Page[model.User],
	error,
) {
	panic("implement me")
}
func (r *UserRepositoryImpl) GetByID(id string) (*model.User, error) {
	panic("implement me")
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
	panic("implement me")
}
func (r *UserRepositoryImpl) Create(user *model.User) (*model.User, error) {
	panic("implement me")
}
func (r *UserRepositoryImpl) Update(user *model.User) (*model.User, error) {
	panic("implement me")
}
func (r *UserRepositoryImpl) Delete(id string) error {
	panic("implement me")
}
