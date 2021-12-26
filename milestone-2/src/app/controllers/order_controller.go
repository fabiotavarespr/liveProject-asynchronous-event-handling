package controllers

import (
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/src/app/models"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/src/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PostOrder func for post a new order.
// @Description Create a new order.
// @Summary create a new order
// @Tags Order
// @Accept json
// @Produce json
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

	// Return status 201 Created.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"order": order,
	})
}
