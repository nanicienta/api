package usecases

import (
	"github.com/nanicienta/api/account-svc/internal/domain/model"
	"github.com/nanicienta/api/account-svc/internal/ports/outbound"
	"github.com/nanicienta/api/pkg/domain"
	"github.com/nanicienta/api/pkg/ports/logging"
)

type UserUseCases struct {
	log            logging.Logger
	userRepository outbound.UserRepository
}

func NewUserUseCases(logger logging.Logger, userRepository outbound.UserRepository) *UserUseCases {
	return &UserUseCases{
		log:            logger,
		userRepository: userRepository,
	}
}

func (u *UserUseCases) GetAll(page int, size int) (domain.Page[model.User], error) {
	users, err := u.userRepository.GetPage(page, size)
	if err != nil {
		u.log.Error("Error getting users", "error", err)
		return domain.Page[model.User]{}, err
	}
	return users, nil
}
