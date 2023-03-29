package app

import (
	"bookshelf-go/routes"
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func StartServer(db *sql.DB) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	//variables para ingresar a SQL
	PORT := os.Getenv("PORT")
	app := fiber.New()

	routes.RouterContainer(app, db)

	app.Listen(PORT)
}
