package services

import "github.com/gofiber/fiber/v2"

func Test() fiber.Handler {

	return func(c *fiber.Ctx) error {
		return c.SendString("Hola esto funciona")
	}
}
