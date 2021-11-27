package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	searchController := CreateSearchController()
	app.Get("/trie/initialize", searchController.CreateTrie)
	app.Post("/trie/enroll/:id", searchController.EnrollWords)
	app.Post("/trie/search/:id", searchController.AutoComplete)
	log.Fatal(app.Listen(":3000"))
}


