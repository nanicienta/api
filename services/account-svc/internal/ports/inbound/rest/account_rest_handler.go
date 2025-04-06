// Package rest handler for account-related HTTP requests
package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nanicienta/api/account-svc/internal/application/command"
	"github.com/nanicienta/api/account-svc/internal/application/usecase"
	"github.com/nanicienta/api/account-svc/internal/infrastructure/server/middleware"
	"github.com/nanicienta/api/pkg/domain/auth"
	"github.com/nanicienta/api/pkg/ports/logging"
	"net/http"
)

// AccountRestHandler is the handler for account-related HTTP requests
type AccountRestHandler struct {
	gin         *gin.Engine
	userUseCase *usecase.AccountUseCase
	logger      logging.Logger
}

// NewAccountRestHandler Creates a new AccountRestHandler
func NewAccountRestHandler(
	gin *gin.Engine,
	logger logging.Logger,
	userUseCase *usecase.AccountUseCase,
) AccountRestHandler {
	return AccountRestHandler{
		gin:         gin,
		logger:      logger,
		userUseCase: userUseCase,
	}
}

// InitRouter initializes the router for account-related endpoints
func (u *AccountRestHandler) InitRouter() {
	api := u.gin.Group("/api/v1/account")
	secured := api.Group("", middleware.Auth())
	{
		secured.POST("", u.createAccount)
		secured.GET("/me", u.me)
		secured.PUT("/me", u.updateAccount)
	}
}

func (u *AccountRestHandler) createAccount(c *gin.Context) {
	account, err := u.parseCreateAccountCommand(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, &gin.H{
				"message": "Invalid command",
			},
		)
		return
	}
	result, appError := u.userUseCase.CreateAccount(account)
	if appError != nil {
		c.AbortWithStatusJSON(appError.HTTPStatus, appError)
		return
	}
	c.JSON(http.StatusCreated, result)

}

func (u *AccountRestHandler) updateAccount(c *gin.Context) {
	userID := c.GetString(auth.ContextUserIDKey)
	updateProfileRequest, err := u.parseUpdateProfileRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, &gin.H{
				"message": "Invalid command",
			},
		)
		return
	}
	result, appError := u.userUseCase.UpdateAccount(
		userID,
		updateProfileRequest,
	)
	if appError != nil {
		c.AbortWithStatusJSON(appError.HTTPStatus, appError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (u *AccountRestHandler) me(c *gin.Context) {
	userID := c.GetString(auth.ContextUserIDKey)
	result, appError := u.userUseCase.GetAccountByID(userID)
	if appError != nil {
		c.AbortWithStatusJSON(appError.HTTPStatus, appError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (u *AccountRestHandler) parseUpdateProfileRequest(ctx *gin.Context) (
	command.UpdateProfileCommand,
	error,
) {
	var req command.UpdateProfileCommand
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return command.UpdateProfileCommand{}, err
	}
	return req, nil
}

func (u *AccountRestHandler) parseCreateAccountCommand(c *gin.Context) (
	command.CreateAccountCommand,
	error,
) {
	var req command.CreateAccountCommand
	if err := c.ShouldBindJSON(&req); err != nil {
		return command.CreateAccountCommand{}, err
	}
	return req, nil
}
