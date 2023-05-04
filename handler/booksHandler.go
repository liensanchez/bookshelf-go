package handler

import (
	"bookshelf-go/services"

	"github.com/gofiber/fiber/v2"
)

func GetBooksHandler(c *fiber.Ctx) error {
	response, err := services.GetBooks()
	if err != nil {
		return c.SendString("Cannot get all books")
	}
	return c.JSON(response)

}

func OneBookHandler(c *fiber.Ctx) error {

	idBook := c.Params("id")

	response, err := services.GetOneBook(idBook)
	if err != nil {
		return c.SendString("Cannot get Book" + idBook)
	}

	return c.JSON(response)
}
