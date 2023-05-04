package handler

import (
	"bookshelf-go/models"
	"bookshelf-go/services"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func GetUserHandler(db *sql.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		response, err := services.GetUsers(db)
		if err != nil {
			return err
		}

		return c.JSON(response)
	}
}

func CreateUserHandler(db *sql.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {

		newUser := new(models.User)

		if err := c.BodyParser(newUser); err != nil {
			return c.Status(405).SendString(err.Error())
		}

		response, err := services.CreateUser(db, newUser)
		if err != nil {
			return err
		}

		return c.JSON(response)
	}
}
