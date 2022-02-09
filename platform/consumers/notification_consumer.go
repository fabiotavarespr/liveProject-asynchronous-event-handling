package consumers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/events"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/handlers"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/models"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger/attributes"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/platform/database"
)

// SubscribeAndListen will subscribe to a Kafka topic and start polling and listening for events
// Adpated from https://github.com/confluentinc/confluent-kafka-go#examples
func (c *Consumer) NotificationSubscribeAndListen() error {

	kc, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     c.Broker,
		"broker.address.family": "v4",
		"group.id":              c.Group + "-notification",
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest"})

	if err != nil {
		logger.Error("Failed to create consumer", attributes.New().WithError(err))

		return err
	}

	logger.Info("Created Consumer", attributes.New().WithField("consumer", kc))

	err = kc.SubscribeTopics([]string{c.Topic}, nil)
	if err != nil {
		logger.Error("Failed to subscribe to topic", attributes.New().WithField("topic", c.Topic).WithError(err))

		return err
	}

	for {
		msg, err := kc.ReadMessage(-1)
		if err != nil {
			// The client will automatically try to recover from all errors.
			logger.Error("Reading error", attributes.New().WithError(err))

			logger.Warn("Closing consumer...", nil)
			kc.Close()

			return err
		}

		logger.Info(string(msg.Value), attributes.New().WithField("topic", msg.TopicPartition))

		var event events.Notification
		if err = json.Unmarshal([]byte(string(msg.Value)), &event); err != nil {
			logger.Error("an issue occurred unmarshalling event from message received", attributes.New().WithField("error", err))

			continue
		}

		notification, err := extractNotification(event)
		if err != nil {
			logger.Error("an issue occurred trying to extract order information from the order recieved event", attributes.New().WithError(err))

			handlers.HandleError(event)
			continue
		}

		if err = processNotificationEvent(event, notification); err != nil {
			logger.Error("an issue occurred trying to process the event", attributes.New().WithError(err))

			handlers.HandleError(event)
			continue
		}
	}
}

func extractNotification(event events.Notification) (models.Notification, error) {
	logger.Info("attempting to extract order from event", attributes.New().WithField("event", event))

	body := event.Body()
	notification, ok := body.(models.Notification)
	if !ok {
		return models.Notification{}, errors.New("event body can't be cast as a Notification")
	}

	return notification, nil
}

func processNotificationEvent(event events.Event, notification models.Notification) error {
	var err error

	db, err := database.OpenDBConnection()
	if err != nil {
		logger.Error("an issue occurred trying to make a connection to the database", attributes.New().WithError(err))
		return err
	}

	// check to see if event has already been processed
	eventAlreadyProcessed, err := db.EventExists(event)
	if err != nil {
		logger.Error("an issue occurred trying to check if an event was already processed", attributes.New().WithError(err))
		return err
	}

	// if event has already been processed, nothing more to do
	if eventAlreadyProcessed {
		logger.Info("event was processed previously", attributes.New().WithField("event.id", event.ID()).
			WithField("event.name", event.Name()))
		return nil
	}

	// event hasn't been processed yet, send the notification
	switch notification.Type {
	case models.Email:
		if err = handlers.SendEmail(notification); err != nil {
			logger.Error("an issue occurred trying to send an email to the customer", attributes.New().WithError(err))

			return err
		}
	default:
		logger.Error("notification type is not supported at this time", attributes.New().WithField("notification.type", notification.Type))

		return fmt.Errorf("notification type, \"%s\" is not supported", notification.Type)
	}

	// mark the event as processed
	if err = db.CreateProcessedEvent(event); err != nil {
		logger.Error("an issue occurred trying to insert the event", attributes.New().WithError(err))
		return err
	}

	return nil
}
