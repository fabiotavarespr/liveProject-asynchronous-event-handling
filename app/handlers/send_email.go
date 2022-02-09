package handlers

import (
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/models"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger/attributes"
)

// SendEmail will construct and send an email based on the supplied notification information
func SendEmail(notification models.Notification) error {
	logger.Info("attempting to send an email notification", attributes.New().WithField("notification.from", notification.From))

	logger.Info("would send email here",
		attributes.New().
			WithField("subject", notification.Subject).
			WithField("body", notification.Body).
			WithField("recipient", notification.Recipient).
			WithField("from", notification.From))

	return nil
}
