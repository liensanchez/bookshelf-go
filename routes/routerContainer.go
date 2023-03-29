package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func RouterContainer(app *fiber.App, db *sql.DB) {

	api := app.Group("/api")

	users := api.Group("/users")

	shelf := api.Group("/shelf")

	books := api.Group("/books")

	UserRoutes(users, db)

	ShelfRoutes(shelf)

	BooksRoutes(books)

}
