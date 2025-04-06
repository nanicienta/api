// Package usecase provides the implementation of the use case interface for the application.
package usecase

import (
	"github.com/nanicienta/api/account-svc/internal/application/command"
	"github.com/nanicienta/api/account-svc/internal/domain/model"
	"github.com/nanicienta/api/account-svc/internal/ports/outbound/repository"
	"github.com/nanicienta/api/pkg/domain"
	"github.com/nanicienta/api/pkg/domain/errors"
	"github.com/nanicienta/api/pkg/ports/logging"
)

// AccountUseCase is the use case for managing accounts
type AccountUseCase struct {
	log                 logging.Logger
	accountRepository   repository.AccountRepository
	notificationService NotificationUseCase
}

// NewAccountUseCase Creates a new AccountUseCase (Like user service)
func NewAccountUseCase(
	logger logging.Logger,
	accountRepository repository.AccountRepository,
) *AccountUseCase {
	return &AccountUseCase{
		log:               logger,
		accountRepository: accountRepository,
	}
}

// GetAll Returns all users paginated
func (u *AccountUseCase) GetAll(page int, size int) (domain.Page[model.Account], *errors.AppError) {
	users, err := u.accountRepository.GetPage(page, size)
	if err != nil {
		u.log.Error("Error getting accounts", "error", err)
		return domain.Page[model.Account]{}, errors.NewInternalServerError(
			"Error getting accounts",
			"Internal server error while getting accounts",
			errors.InternalServerErrorGettingAccounts,
		).AppError
	}
	return users, nil
}

// UpdateAccount Updates the account with the information on the command.UpdateProfileCommand
func (u *AccountUseCase) UpdateAccount(_ string, _ command.UpdateProfileCommand) (
	*model.Account,
	*errors.AppError,
) {
	return nil, nil
}

// CreateAccount Creates a new account with the information on the command.CreateAccountCommand
func (u *AccountUseCase) CreateAccount(command command.CreateAccountCommand) (
	*model.Account,
	*errors.AppError,
) {
	account := model.Account{
		Name:     command.Name,
		Email:    command.Email,
		Timezone: "America/Los_Angeles",
		Type:     command.Type,
		Source:   command.Source,
		IsActive: true,
		IsLocked: false,
	}
	persistedAccount, err := u.accountRepository.Create(&account)
	if err != nil {
		u.log.Error("Error creating account", "error", err)
		return nil, errors.NewInternalServerError(
			"Error creating account",
			"Internal server error while creating account",
			errors.InternalServerErrorCreatingAccount,
		).AppError
	}
	notificationID, appError := u.notificationService.SendAccountCreatedNotification(
		persistedAccount,
	)
	if appError != nil {
		u.log.Error("Error sending notification", "error", appError.Error())
		return nil, errors.NewInternalServerError(
			"Error sending notification",
			"Internal server error while sending notification",
			errors.InternalServerErrorSendingNotification,
		).AppError
	}
	//TODO add notification to the account
	u.log.Info("Notification sent", "notificationID", notificationID)
	return persistedAccount, nil
}

// GetAccountByID Returns the account with the given id
func (u *AccountUseCase) GetAccountByID(id string) (*model.Account, *errors.AppError) {
	account, err := u.accountRepository.GetByID(id)
	if err != nil {
		u.log.Error("Error getting user", "error", err)
		return nil, errors.NewInternalServerError(
			"Error getting account",
			"Internal server error while getting account",
			errors.InternalServerErrorGettingAccounts,
		).AppError
	}
	return account, nil
}
