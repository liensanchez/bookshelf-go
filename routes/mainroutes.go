package routes

import "github.com/gofiber/fiber/v2"

func RouterContainer(app *fiber.App) {

	api := app.Group("/api")

	users := api.Group("/users")

	UserRoutes(users)

}
