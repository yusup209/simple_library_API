package main

import (
	"fmt"
	"libraryapi/configs"
	"libraryapi/models"
	"libraryapi/routes"
)

func main() {
	configs.DB.Debug().AutoMigrate(&models.Users{}, &models.Category{}, &models.Bookshelf{}, &models.Book{}, &models.BookTransaction{}, &models.Guestbook{})
	fmt.Println("Helloooo")
	routes.Setup()
}
