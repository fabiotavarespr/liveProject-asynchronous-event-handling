package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/src/pkg/configs"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/src/pkg/middleware"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/src/pkg/routes"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/src/pkg/utils"

	_ "github.com/fabiotavarespr/liveProject-asynchronous-event-handling/docs"
	_ "github.com/joho/godotenv/autoload"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.HealthRoute(app)   // Register a health routes.
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.StartServer(app)
}
