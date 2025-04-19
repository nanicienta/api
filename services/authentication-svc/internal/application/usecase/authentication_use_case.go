// Package usecase provides the implementation of the authentication use cases.
package usecase

import (
	"encoding/hex"
	"time"

	"github.com/nanicienta/api/authentication-svc/internal/application/command"
	"github.com/nanicienta/api/authentication-svc/internal/application/result"
	"github.com/nanicienta/api/authentication-svc/internal/domain/model"
	"github.com/nanicienta/api/authentication-svc/internal/ports/outbound/repository"
	"github.com/nanicienta/api/authentication-svc/internal/ports/outbound/service/password"
	"github.com/nanicienta/api/authentication-svc/internal/ports/outbound/service/token"
	"github.com/nanicienta/api/pkg/domain/errors"
	"github.com/nanicienta/api/pkg/ports/logging"
)

// AuthenticationUseCase orchestrates all authentication‑related operations.
// It stays “pure” and never touches HTTP, Gorm or crypto primitives directly.
type AuthenticationUseCase struct {
	logger                 logging.Logger
	authIdentityRepository repository.AuthIdentityRepository
	refreshTokenRepository repository.RefreshTokenRepository
	tokenService           token.Service
	passwordService        password.Service
}

// NewAuthenticationUseCase wires the dependencies for the use‑case layer.
func NewAuthenticationUseCase(
	logger logging.Logger,
	authIdentityRepository repository.AuthIdentityRepository,
	tokenRepository repository.RefreshTokenRepository,
	tokenService token.Service,
	passwordService password.Service,
) AuthenticationUseCase {
	return AuthenticationUseCase{
		logger:                 logger,
		authIdentityRepository: authIdentityRepository,
		refreshTokenRepository: tokenRepository,
		tokenService:           tokenService,
		passwordService:        passwordService,
	}
}

// PasswordLogin authenticates an identity with e‑mail + password
// and returns fresh access/refresh tokens if the credentials are valid.
func (uc AuthenticationUseCase) PasswordLogin(
	cmd command.PasswordLoginCommand,
) (*result.LoginResult, *errors.AppError) {

	identity, err := uc.authIdentityRepository.
		GetByEmailAndOrganizationID(cmd.Email, cmd.OrganizationID)
	if err != nil {
		uc.logger.Error("query auth identity", "err", err)
		return nil, errors.NewInternalServerError(
			"failed to get auth identity", "query error",
			errors.InternalServerErrorGettingAuthIdentity,
		).AppError
	}

	// Guard: identity exists, password matches, account not locked
	if identity == nil ||
		!uc.passwordService.VerifyPassword(identity.Hash, cmd.Password) ||
		identity.IsLocked {
		_ = uc.updateBadAttemptCount(identity)
		return nil, errors.NewUnauthorizedError(
			"Invalid credentials",
			"The username or password are invalid",
			errors.InvalidCredentials,
		).AppError
	}

	// Happy path → reset attempts + issue tokens
	identity.BadLoginAttempts = 0
	_ = uc.authIdentityRepository.UpdateBadAttemptCount(identity)

	return uc.generateTokens(identity, cmd.DeviceID)
}

// Logout invalidates (deletes) the refresh‑token for a given device.
func (uc AuthenticationUseCase) Logout(cmd command.LogoutCommand) *errors.AppError {
	if err := uc.refreshTokenRepository.
		DeleteRefreshTokenByUserIDAndDeviceID(cmd.UserID, cmd.DeviceID); err != nil {

		uc.logger.Error("delete refresh token", "err", err)
		return errors.NewInternalServerError(
			"failed to delete refresh token",
			"db error",
			errors.InternalServerErrorCode,
		).AppError
	}
	return nil
}

// RefreshToken rotates a valid refresh‑token and returns a new
// access/refresh pair. It also revokes the previous refresh‑token.
func (uc AuthenticationUseCase) RefreshToken(
	cmd command.RefreshTokenCommand,
) (*result.LoginResult, *errors.AppError) {

	hash, err := uc.tokenService.VerifyRefresh(cmd.RefreshToken)
	if err != nil {
		return nil, errors.NewUnauthorizedError(
			"invalid refresh token",
			"parse error",
			errors.InvalidRefreshToken,
		).AppError
	}

	rt, err := uc.refreshTokenRepository.
		GetRefreshTokenByDeviceIDAndToken(cmd.DeviceID, hex.EncodeToString(hash[:]))
	if err != nil || rt == nil || rt.IsRevoked() {
		return nil, errors.NewUnauthorizedError(
			"invalid refresh token",
			"token not found or revoked",
			errors.InvalidRefreshToken,
		).AppError
	}

	identity, _ := uc.authIdentityRepository.GetByID(rt.UserID)

	// revoke old token
	_ = uc.refreshTokenRepository.DeleteByID(rt.ID)

	return uc.generateTokens(identity, cmd.DeviceID)
}

// generateTokens creates a new access‑token (JWT) and a new opaque
// refresh‑token linked to the given device, persisting the latter.
func (uc AuthenticationUseCase) generateTokens(
	identity *model.AuthIdentityModel,
	deviceID string,
) (*result.LoginResult, *errors.AppError) {
	tks, err := uc.tokenService.Issue(identity, deviceID)
	if err != nil {
		uc.logger.Error("issue tokens", "err", err)
		return nil, errors.NewInternalServerError(
			"failed to issue tokens",
			"jwt/sign",
			errors.InternalServerErrorCode,
		).AppError
	}

	rtModel := &model.RefreshTokenModel{
		UserID:    identity.ID,
		Token:     tks.RefreshToken,
		DeviceID:  deviceID,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}
	if err = uc.refreshTokenRepository.Save(rtModel); err != nil {
		uc.logger.Error("save refresh token", "err", err)
		return nil, errors.NewInternalServerError(
			"failed to persist refresh token",
			"db error",
			errors.InternalServerErrorCode,
		).AppError
	}

	return &result.LoginResult{
		AccessToken:  tks.Access,
		RefreshToken: tks.RefreshToken,
		ExpiresIn:    int64(time.Until(tks.AccessExp).Seconds()),
	}, nil
}

// updateBadAttemptCount increments failed attempts and locks
// the identity after five consecutive failures.
func (uc AuthenticationUseCase) updateBadAttemptCount(
	identity *model.AuthIdentityModel,
) error {
	if identity == nil || identity.IsLocked {
		return nil
	}
	identity.BadLoginAttempts++
	if identity.BadLoginAttempts >= 5 {
		identity.IsLocked = true
	}
	return uc.authIdentityRepository.UpdateBadAttemptCount(identity)
}
