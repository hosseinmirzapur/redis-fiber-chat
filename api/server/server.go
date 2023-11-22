package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hosseinmirzapur/rthnk/api/routes"
)

func RunServer() error {
	app := fiber.New()

	app.Use(cors.New())

	routes.RegisterRoutes(app)

	return app.Listen(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
}
