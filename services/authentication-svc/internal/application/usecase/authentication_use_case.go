// Package usecase provides the implementation of the use case interface for the application.
package usecase

import (
	"github.com/nanicienta/api/authentication-svc/internal/application/command"
	"github.com/nanicienta/api/authentication-svc/internal/application/result"
	"github.com/nanicienta/api/authentication-svc/internal/domain/model"
	"github.com/nanicienta/api/authentication-svc/internal/ports/outbound/repository"
	"github.com/nanicienta/api/pkg/domain/errors"
	"github.com/nanicienta/api/pkg/ports/logging"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticationUseCase is the use case for authentication-related operations
type AuthenticationUseCase struct {
	logger                 logging.Logger
	authIdentityRepository repository.AuthIdentityRepository
	refreshTokenRepository repository.RefreshTokenRepository
}

// NewAuthenticationUseCase creates a new AuthenticationUseCase
func NewAuthenticationUseCase(
	logger logging.Logger,
	authIdentityRepository repository.AuthIdentityRepository,
	tokenRepository repository.RefreshTokenRepository,
) AuthenticationUseCase {
	return AuthenticationUseCase{
		logger:                 logger,
		authIdentityRepository: authIdentityRepository,
		refreshTokenRepository: tokenRepository,
	}
}

// PasswordLogin handles the password login use case
func (c AuthenticationUseCase) PasswordLogin(cmd command.PasswordLoginCommand) (
	*result.LoginResult,
	*errors.AppError,
) {
	authModel, err := c.authIdentityRepository.GetByEmailAndOrganizationID(
		cmd.Email,
		cmd.OrganizationID,
	)
	if err != nil {
		c.logger.Error("failed to get auth identity", "error", err)
		return nil, errors.NewInternalServerError(
			"failed to get auth identity", "error",
			errors.InternalServerErrorGettingAuthIdentity,
		).AppError
	}
	if authModel == nil || !verifyPassword(authModel.Hash, cmd.Password) || authModel.IsLocked {
		if authModel != nil {
			err = c.updateBadAttemptCount(authModel)
			if err != nil {
				c.logger.Error("failed to update bad attempt count", "error", err)
				return nil, errors.NewInternalServerError(
					"failed to get auth identity", "error",
					errors.InternalServerErrorCode,
				).AppError
			}
		}
		c.logger.Warn("invalid credentials", "email", cmd.Email)
		return nil, errors.NewUnauthorizedError(
			"Invalid credentials",
			"The username or password are invalid",
			errors.InvalidCredentials,
		).AppError
	}

	return &result.LoginResult{}, nil
}

// Logout handles the logout use case
func (c AuthenticationUseCase) Logout(cmd command.LogoutCommand) *errors.AppError {
	err := c.refreshTokenRepository.DeleteRefreshTokenByUserIDAndDeviceID(cmd.UserID, cmd.DeviceID)
	if err != nil {
		c.logger.Error("failed to delete refresh token", "error", err)
		return errors.NewInternalServerError(
			"failed to delete refresh token",
			"error",
			errors.InternalServerErrorCode,
		).AppError
	}
	return nil
}

func (c AuthenticationUseCase) updateBadAttemptCount(authModel *model.AuthIdentityModel) error {
	if authModel.IsLocked {
		return nil
	}
	authModel.BadLoginAttempts++
	if authModel.BadLoginAttempts >= 5 {
		authModel.IsLocked = true
	}
	return c.authIdentityRepository.UpdateBadAttemptCount(authModel)
}

func (c AuthenticationUseCase) RefreshToken(cmd command.RefreshTokenCommand) (
	*result.LoginResult,
	*errors.AppError,
) {

	refreshToken, err := c.refreshTokenRepository.GetRefreshTokenByDeviceIDAndToken(cmd.DeviceID)

	return nil, nil
}

func verifyPassword(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
