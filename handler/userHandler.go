package handler

import (
	"bookshelf-go/services"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func UserHandler(db *sql.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		response, err := services.GetUsers(db)
		if err != nil {
			return err
		}

		return c.JSON(response)
	}
}
