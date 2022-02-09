package consumers

import (
	"encoding/json"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/events"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/handlers"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/models"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/topics"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger/attributes"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/platform/database"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/platform/producers"
	"github.com/google/uuid"
	"time"
)

// OrderSubscribeAndListen will subscribe to a Kafka topic and start polling and listening for events
// Adpated from https://github.com/confluentinc/confluent-kafka-go#examples
func (c *Consumer) OrderSubscribeAndListen() error {

	kc, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     c.Broker,
		"broker.address.family": "v4",
		"group.id":              c.Group + "-inventory",
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest"})

	if err != nil {
		logger.Error("Failed to create consumer", attributes.New().WithError(err))

		return err
	}

	logger.Info("Created Consumer", attributes.New().WithField("consumer", kc))

	if err = kc.SubscribeTopics([]string{c.Topic}, nil); err != nil {
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

		var event events.OrderReceived
		if err = json.Unmarshal([]byte(string(msg.Value)), &event); err != nil {

			continue
		}

		var order models.Order
		if order, err = extractOrder(event); err != nil {
			logger.Error("an issue occurred trying to extract order information from the order recieved event", attributes.New().WithError(err))

			handlers.HandleError(event)
			continue
		}

		if err = processOrderEvent(event, order); err != nil {
			logger.Error("an issue occurred trying to process the event", attributes.New().WithError(err))

			handlers.HandleError(event)
			continue
		}

		if err = publishOrderConfirmedEvent(order); err != nil {
			logger.Error("an issue occurred trying to publish an order confirmed event", attributes.New().WithError(err))

			handlers.HandleError(event)
			continue
		}
	}
}

func extractOrder(event events.OrderReceived) (models.Order, error) {
	logger.Info("attempting to extract order from event", attributes.New().WithField("event", event))

	body := event.Body()
	order, ok := body.(models.Order)
	if !ok {
		return models.Order{}, errors.New("event body can't be cast as an order")
	}

	return order, nil
}

func processOrderEvent(event events.Event, order models.Order) error {
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

	// event hasn't been processed yet, decrement the inventory
	if err = handlers.DecrementInventory(order); err != nil {
		logger.Error("an issue occurred trying to decrement the inventory", attributes.New().WithError(err))
		return err
	}

	// mark the event as processed
	if err = db.CreateProcessedEvent(event); err != nil {
		logger.Error("an issue occurred trying to insert the event", attributes.New().WithError(err))
		return err
	}

	return nil
}

func publishOrderConfirmedEvent(o models.Order) error {
	// publish an order confirmed event
	e := translateOrderToEvent(o)

	logger.Info("transformed order to event", attributes.New().WithField("event", e))

	var err error
	if err = producers.ProducerEvent(e, topics.TopicOrderConfirmed); err != nil {
		return err
	}

	logger.Info("published event", attributes.New().WithField("event", e))

	return nil
}

func translateOrderToEvent(o models.Order) events.Event {
	var event = events.OrderConfirmed{
		EventBase: events.BaseEvent{
			EventID:        uuid.New(),
			EventName:      topics.TopicOrderConfirmed,
			EventTimestamp: time.Now(),
		},
		EventBody: o,
	}

	return event
}
