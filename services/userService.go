package services

import (
	"bookshelf-go/models"
	"database/sql"
	"fmt"
)

func GetUsers(db *sql.DB) ([]models.User, error) {
	querySearch := "SELECT * FROM users"

	var results []models.User

	rows, err := db.Query(querySearch)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var usersFound models.User

		err = rows.Scan(&usersFound.Id, &usersFound.Name, &usersFound.LastName, &usersFound.Shelf)
		if err != nil {
			return nil, err
		}

		results = append(results, usersFound)
	}

	return results, nil
}

func CreateUser(db *sql.DB, user *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO users (name, lastname) VALUES ('%v', '%v')  RETURNING id, name, lastname, shelf_id", user.Name, user.LastName)

	err := db.QueryRow(query).Scan(&user.Id, &user.Name, &user.LastName, &user.Shelf)
	if err != nil {
		return nil, err
	}

	return user, nil
}
