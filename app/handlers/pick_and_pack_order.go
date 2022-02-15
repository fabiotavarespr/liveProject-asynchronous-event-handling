package handlers

import (
	"fmt"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/events"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/models"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/topics"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger/attributes"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/platform/producers"
	"github.com/google/uuid"
	"strings"
	"time"
)

// PickAndPackOrder will alert the warehouse personnel to pick and pack the customers order
func PickAndPackOrder(order models.Order) error {
	logger.Info("attempting to alert warehouse personnel to pick and pack order", attributes.New().WithField("order.id", order.ID))

	// We are not actually connecting to the warehouse system, so just log it for now
	for _, p := range order.Products {
		logger.Info("picking product to be packed for shipping",
			attributes.New().
				WithField("order.id", order.ID).
				WithField("product.code", p.ProductCode).
				WithField("product.quantity", p.Quantity))
	}

	// notify the customer the order is being picked and packed
	var b strings.Builder
	for _, p := range order.Products {
		fmt.Fprintf(&b, "%d of product [%s]", p.Quantity, p.ProductCode)
	}

	address := fmt.Sprintf("<div>Shipping to Address:</div><div>%s</div><div>%s %s, %s</div>", order.Customer.ShippingAddress.Line1, order.Customer.ShippingAddress.City, order.Customer.ShippingAddress.State, order.Customer.ShippingAddress.PostalCode)
	subject := fmt.Sprintf("Hello %s, your order has been received.", order.Customer.FirstName)
	body := fmt.Sprintf("<div>Your order has been received and we will be preparing it for shipping as soon as possible. Here is a review of the products in your order:</div><div>%s</div><div>%s</div>", b.String(), address)

	var err error
	event := events.Notification{
		EventBase: events.BaseEvent{
			EventID:        uuid.New(),
			EventTimestamp: time.Now(),
		},
		EventBody: models.Notification{
			Type:      models.Email,
			Recipient: order.Customer.EmailAddress,
			From:      "orders@ppe4all.com",
			Subject:   subject,
			Body:      body,
		},
	}

	if err = producers.ProducerEvent(event, topics.TopicNotification); err != nil {
		logger.Error("an issue ocurred publishing an error event to Kafka", attributes.New().WithError(err).WithField("topic", topics.TopicNotification))

		return err
	}

	return nil
}
