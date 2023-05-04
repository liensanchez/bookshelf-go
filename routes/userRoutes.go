package routes

import (
	"bookshelf-go/handler"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(users fiber.Router, db *sql.DB) {

	users.Get("/", handler.GetUserHandler(db))

	users.Post("/create", handler.CreateUserHandler(db))
}
