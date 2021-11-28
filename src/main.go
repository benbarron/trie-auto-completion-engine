package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	var app *fiber.App = fiber.New()
	searchController := CreateSearchController()
	app.Get("/trie/initialize", searchController.CreateTrie)
	app.Post("/trie/enroll/:id", searchController.EnrollWords)
	app.Post("/trie/search/:id", searchController.SearchTrie)
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(app.Listen(port))
}

