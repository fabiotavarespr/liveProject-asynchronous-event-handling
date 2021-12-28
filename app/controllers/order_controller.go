package controllers

import (
	"time"

	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/events"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/models"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/topics"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger/attributes"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/utils"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/platform/producers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PostOrder func for post a new order.
// @Description Create a new order.
// @Summary create a new order
// @Tags Order
// @Accept json
// @Produce json
// @Param order body models.Order true "Order"
// @Success 201 {object} models.Order
// @Router /v1/order [post]
func PostOrder(c *fiber.Ctx) error {
	// Create new Order struct
	order := &models.Order{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(order); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Order model.
	validate := utils.NewValidator()

	// Set initialized default data for order:
	order.ID = uuid.New()

	// Validate order fields.
	if err := validate.Struct(order); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	var event = events.OrderReceived{
		EventBase: events.BaseEvent{
			EventID:        uuid.New(),
			EventTimestamp: time.Now(),
		},
		EventBody: *order,
	}

	if err := producers.ProducerEvent(event, topics.TopicOrderReceived); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	details := attributes.New()
	details["order"] = order

	logger.Info(c.Context(), "Handler - POST - Order", details)

	// Return status 201 Created.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"order": order,
	})
}
