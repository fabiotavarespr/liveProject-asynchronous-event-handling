package handlers

import (
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/models"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger/attributes"
)

// DecrementInventory will decrement the inventory of produts by the specific quantity in the order
func DecrementInventory(order models.Order) error {
	logger.Info("attempting to decrement inventory from order", attributes.New().WithField("order.id", order.ID))

	// We are not actually connecting to an inventory system, so just log it for now
	for _, p := range order.Products {
		logger.Info("decrementing inventory for product",
			attributes.New().
				WithField("order.id", order.ID).
				WithField("product.code", p.ProductCode).
				WithField("product.quantity", p.Quantity))
	}

	return nil
}
