package app

import (
	"bookshelf-go/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	routes.RouterContainer(app)

	port := os.Getenv("PORT")

	app.Listen(port)
}
