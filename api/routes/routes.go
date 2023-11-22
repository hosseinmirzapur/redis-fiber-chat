package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	registerAPIRoutes(app)
	registerWSRoutes(app)
}
