package routes

import "github.com/gofiber/fiber/v2"

func registerAPIRoutes(app *fiber.App) {
	app.Group("/api")
}
