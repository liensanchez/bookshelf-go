package main

import (
	"bookshelf-go/app"
	"bookshelf-go/database"
)

func main() {
	db := database.StartDatabase(false)
	app.StartServer(db)
}
