package usecase

import (
	"github.com/nanicienta/api/account-svc/internal/domain/model"
	"github.com/nanicienta/api/pkg/domain/command/notification"
	"github.com/nanicienta/api/pkg/domain/errors"
	"github.com/nanicienta/api/pkg/ports/logging"
	"github.com/nanicienta/api/pkg/ports/outbound/client"
)

// NotificationUseCase is the use case for sending notifications
type NotificationUseCase struct {
	log                logging.Logger
	notificationClient client.NotificationClient
}

// NewNotificationUseCase Creates a new NotificationUseCase
func NewNotificationUseCase(
	log logging.Logger, notificationClient client.
		NotificationClient,
) *NotificationUseCase {
	return &NotificationUseCase{
		log:                log,
		notificationClient: notificationClient,
	}
}

// SendAccountCreatedNotification Sends a notification when an account is created
func (c *NotificationUseCase) SendAccountCreatedNotification(_ *model.Account) (
	string,
	*errors.AppError,
) {
	notificationID, err := c.notificationClient.SendNotification(
		notification.
			SendNotificationCommand{},
	)
	if err != nil {
		c.log.Error("Error sending notification", "error", err)
		return "", errors.NewInternalServerError(
			"Error sending notification",
			"Internal server error while sending notification",
			errors.InternalServerErrorSendingNotification,
		).AppError
	}
	return notificationID, nil
}
