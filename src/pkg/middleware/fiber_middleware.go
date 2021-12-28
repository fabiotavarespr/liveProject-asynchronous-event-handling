package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),

		// Add Request Id
		requestid.New(),

		// Add Recover
		recover.New(),

		// // Add simple logger.
		// logger.New(logger.Config{
		// 	Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		// 	TimeFormat: time.RFC3339,
		// }),
	)
}
