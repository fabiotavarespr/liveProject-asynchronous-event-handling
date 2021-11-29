package routes

import (
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("health", controllers.Health)
}