package routes

import (
	"bookshelf-go/handler"

	"github.com/gofiber/fiber/v2"
)

func BooksRoutes(books fiber.Router) {

	books.Get("/", handler.GetBooksHandler)

	books.Get("/:id", handler.OneBookHandler)
}
