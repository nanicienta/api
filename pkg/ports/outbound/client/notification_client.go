// Package client provides the client interfaces for outbound communication with external services.
package client

import "github.com/nanicienta/api/pkg/domain/command/notification"

// NotificationClient is an interface that defines the methods for sending notifications.
type NotificationClient interface {
	// SendNotification sends a notification to the user.
	SendNotification(request notification.SendNotificationCommand) (string, error)
}
