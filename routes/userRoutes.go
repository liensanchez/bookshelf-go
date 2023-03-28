package routes

import (
	"bookshelf-go/services"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(users fiber.Router) {

	users.Get("/", services.Test())
}
