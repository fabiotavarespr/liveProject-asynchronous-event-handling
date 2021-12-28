package controllers

import (
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/models"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger/attributes"
	"github.com/gofiber/fiber/v2"
)

const (
	UP   = "UP"
	DOWN = "DOWN"
)

// GetHealth func check health.
// @Description Health check.
// @Summary health check
// @Produce json
// @Success 200 {array} models.Health
// @Router /health [get]
func GetHealth(c *fiber.Ctx) error {

	health := models.Health{
		Status: UP,
	}

	details := attributes.New()
	details["health"] = health

	logger.Info(c.Context(), "Handler - GET - Health", details)

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"health": health,
	})
}
