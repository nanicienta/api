// Package rest handler for account-related HTTP requests
package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nanicienta/api/authentication-svc/internal/application/command"
	"github.com/nanicienta/api/authentication-svc/internal/application/usecase"
	"github.com/nanicienta/api/pkg/domain/auth"
	"github.com/nanicienta/api/pkg/domain/errors"
	"github.com/nanicienta/api/pkg/ports/logging"
	"net/http"
)

// AuthenticationRestHandler is the handler for authentication-related HTTP requests
type AuthenticationRestHandler struct {
	gin                   *gin.Engine
	authenticationUseCase *usecase.AuthenticationUseCase
	logger                logging.Logger
}

// NewAuthenticationRestHandler Creates a new AuthenticationRestHandler
func NewAuthenticationRestHandler(
	gin *gin.Engine,
	logger logging.Logger,
	authenticationUseCase *usecase.AuthenticationUseCase,
) AuthenticationRestHandler {
	return AuthenticationRestHandler{
		gin:                   gin,
		logger:                logger,
		authenticationUseCase: authenticationUseCase,
	}
}

// InitRouter initializes the router for account-related endpoints
func (arh *AuthenticationRestHandler) InitRouter() {
	api := arh.gin.Group("/api/v1/auth")
	public := api.Group("")
	{
		public.POST("/login", arh.login)
		public.POST("/refresh", arh.refresh)
		public.GET("/providers", arh.getProvidersByOrg)
	}
	private := api.Group("")
	{
		private.DELETE("/logout", arh.logout)
	}
}

func (arh *AuthenticationRestHandler) login(c *gin.Context) {
	cmd, err := parseLoginCommand(c)
	if err != nil {
		arh.logger.Error("failed to parse login command (Parsing the context)", "err", err)
		c.AbortWithStatusJSON(
			http.StatusBadRequest, errors.NewBadRequestError(
				"Failed to parse body",
				"Failed to parse body", errors.InvalidRequest,
			),
		)
		return
	}

	out, appErr := arh.authenticationUseCase.PasswordLogin(cmd)
	if appErr != nil {
		arh.logger.Warn("login failed", "err", appErr)
		c.AbortWithStatusJSON(appErr.HTTPStatus, appErr)
		return
	}

	c.JSON(http.StatusCreated, out)
}

func (arh *AuthenticationRestHandler) logout(c *gin.Context) {
	cmd := command.LogoutCommand{
		DeviceID: c.GetHeader(auth.DeviceIdHeaderName),
	}
	if err := arh.authenticationUseCase.Logout(cmd); err != nil {
		arh.logger.Warn("logout failed", "err", err)
		c.AbortWithStatusJSON(err.HTTPStatus, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (arh *AuthenticationRestHandler) refresh(c *gin.Context) {
	cmd, err := parseRefreshTokenCommand(c)
	if err != nil {
		arh.logger.Error("failed to parse refresh token command (Parsing the context)", "err", err)
		c.AbortWithStatusJSON(
			http.StatusUnauthorized, errors.AppError{
				Message:     "failed to parse refresh token command",
				HTTPStatus:  http.StatusUnauthorized,
				Description: "failed to parse refresh token command",
			},
		)
		return
	}
	result, appErr := arh.authenticationUseCase.RefreshToken(cmd)
	if appErr != nil {
		arh.logger.Error("failed to refresh token", "err", appErr)
		c.AbortWithStatusJSON(appErr.HTTPStatus, appErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (arh *AuthenticationRestHandler) getProvidersByOrg(context *gin.Context) {
	panic("implement me")
}

func parseRefreshTokenCommand(c *gin.Context) (command.RefreshTokenCommand, error) {
	var cmd command.RefreshTokenCommand
	if err := c.ShouldBindJSON(&cmd); err != nil {
		return cmd, err
	}
	cmd.DeviceID = c.GetHeader(auth.DeviceIdHeaderName)
	return cmd, nil
}

// HELPERS

func parseLoginCommand(c *gin.Context) (
	command.PasswordLoginCommand,
	error,
) {
	var cmd command.PasswordLoginCommand
	if err := c.ShouldBindJSON(&cmd); err != nil {
		return cmd, err
	}
	cmd.DeviceID = c.GetHeader(auth.DeviceIdHeaderName)
	return cmd, nil
}
