package routes

import (
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/src/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// SwaggerRoute func for describe group of API Docs routes.
func HealthRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/health")

	// Routes for GET method:
	route.Get("*", controllers.GetHealth) // get one user by ID
}
