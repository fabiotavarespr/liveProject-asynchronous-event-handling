package routes

import (
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/src/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/v1")

	// Routes for GET method:
	route.Post("/order", controllers.PostOrder) // post new order
}
