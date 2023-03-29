package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Shelf    int    `json:"shelf"`
}
