package services

import (
	"bookshelf-go/models"
	"database/sql"
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
