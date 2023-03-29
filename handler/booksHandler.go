package handler

import (
	"bookshelf-go/services"

	"github.com/gofiber/fiber/v2"
)

func BookHandler(c *fiber.Ctx) error {
	response, err := services.GetBooks()
	if err != nil {
		return c.SendString("ESTA PORONGA SE ROMPIO!!!")
	}
	return c.JSON(response)

}
