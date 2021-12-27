package controllers

import (
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/src/app/models"
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

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"health": health,
	})
}
