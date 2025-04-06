// Package gorm contains all the stuff related to the gorm adapters
package gorm

import (
	"github.com/nanicienta/api/account-svc/internal/domain/model"
	"github.com/nanicienta/api/account-svc/internal/infrastructure/adapters/entity"
	"github.com/nanicienta/api/account-svc/internal/ports/outbound/repository"
	"github.com/nanicienta/api/pkg/domain"
	"github.com/nanicienta/api/pkg/ports/logging"
	"gorm.io/gorm"
	"math"
)

// NewAccountRepository creates a new instance of AccountRepository
func NewAccountRepository(logger logging.Logger, db *gorm.DB) repository.AccountRepository {
	return &AccountRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

// AccountRepositoryImpl is the implementation of repository.AccountRepository
type AccountRepositoryImpl struct {
	logger logging.Logger
	db     *gorm.DB
}

// GetAll retrieves all accounts
func (r *AccountRepositoryImpl) GetAll() ([]model.Account, error) {
	var accountEntities []entity.AccountEntity
	if err := r.db.Find(&accountEntities).Error; err != nil {
		r.logger.Error("Error getting accounts", "error", err)
		return nil, err
	}
	accounts := make([]model.Account, len(accountEntities))
	for i, accountEntity := range accountEntities {
		accounts[i] = accountEntity.ToModel()
	}
	return accounts, nil
}

// TODO move to the pkg module
func calculateTotalPages(totalRows int, pageSize int) int {
	if pageSize == 0 {
		return 0
	}
	return int(math.Ceil(float64(totalRows) / float64(pageSize)))
}

// TODO move to a mapper maybe?
func mapAccountEntitiesToModels(entities []entity.AccountEntity) []model.Account {
	models := make([]model.Account, len(entities))
	for i := range entities {
		models[i] = entities[i].ToModel()
	}
	return models
}

// GetPage retrieves a paginated list of accounts
func (r *AccountRepositoryImpl) GetPage(
	pageNumber int,
	pageSize int,
) (domain.Page[model.Account], error) {
	var totalRows int64
	r.db.Model(&entity.AccountEntity{}).Count(&totalRows)
	totalPages := calculateTotalPages(int(totalRows), pageSize)
	var users []entity.AccountEntity
	if err := r.db.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return domain.EmptyPage[model.Account](), err
	}
	pageBuilder := new(domain.PageBuilder[model.Account])
	return pageBuilder.SetItems(mapAccountEntitiesToModels(users)).
		SetTotal(int(totalRows)).
		SetPageSize(pageSize).
		SetPageNumber(pageNumber).
		SetTotalPages(totalPages).
		Build(), nil
}

// GetByID retrieves an account by ID
func (r *AccountRepositoryImpl) GetByID(id string) (*model.Account, error) {
	var accountEntity entity.AccountEntity
	if err := r.db.First(&accountEntity, "id = ?", id).Error; err != nil {
		r.logger.Error("Error getting account by ID", "error", err)
		return nil, err
	}
	accountModel := accountEntity.ToModel()
	return &accountModel, nil

}

// GetByEmail retrieves an account by email
func (r *AccountRepositoryImpl) GetByEmail(_ string) (*model.Account, error) {
	panic("implement me")
}

// Create creates a new account
func (r *AccountRepositoryImpl) Create(_ *model.Account) (*model.Account, error) {
	panic("implement me")
}

// Update updates an existing account
func (r *AccountRepositoryImpl) Update(_ *model.Account) (*model.Account, error) {
	panic("implement me")
}

// Delete removes an account by ID (Applies soft delete)
func (r *AccountRepositoryImpl) Delete(id string) error {
	tx := r.db.Update("is_deleted", true).Where("id = ?", id)
	if tx.Error != nil {
		r.logger.Error("Error deleting account", "error", tx.Error)
		return tx.Error
	}
	return nil
}
