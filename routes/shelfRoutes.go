package routes

import (
	"bookshelf-go/services"

	"github.com/gofiber/fiber/v2"
)

func ShelfRoutes(users fiber.Router) {

	users.Get("/", services.TestShelf())
}
