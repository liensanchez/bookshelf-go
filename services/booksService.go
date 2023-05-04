package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Book struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Authors []struct {
		Name string `json:"name"`
	} `json:"authors"`
	Copyright bool `json:"copyright"`
	Status    string
}

type Response struct {
	Results []Book `json:"results"`
}

func GetBooks() (Response, error) {

	url := "http://gutendex.com/books?sort=ascending"

	responseURL, err := http.Get(url)
	if err != nil {
		fmt.Println("Error en solicitud")
	}
	defer responseURL.Body.Close()

	body, err := io.ReadAll(responseURL.Body)
	if err != nil {
		fmt.Println("Error en lectura")
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error occurred while unmarshaling response:", err)
		return response, err
	}

	return response, nil
}

func GetOneBook(id string) (Book, error) {

	url := fmt.Sprintf("http://gutendex.com/books/%v/", id)

	responseURL, err := http.Get(url)
	if err != nil {
		fmt.Println("Error en solicitud")
	}
	defer responseURL.Body.Close()

	body, err := io.ReadAll(responseURL.Body)
	if err != nil {
		fmt.Println("Error en lectura")
	}

	var response Book
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error occurred while unmarshaling response:", err)
		return response, err
	}

	return response, nil
}
