package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type SearchController struct {
	Nodes *NodePool
}

type EnrollWordsRequest struct {
	Words []string `json:"words"`
}

type SearchTrieRequest struct {
	Prefix string `json:"prefix"`
}

func CreateSearchController() *SearchController {
	return &SearchController{
		Nodes: CreateNodePool(),
	}
}

func (cntrl *SearchController) CreateTrie(ctx *fiber.Ctx) error {
	id := cntrl.Nodes.AddNodeToPool()
	return ctx.JSON(fiber.Map{
		"message": "Node Created",
		"id": id,
	})
}

func (cntrl *SearchController) EnrollWords (ctx *fiber.Ctx) error {
	request := EnrollWordsRequest{}
	ctx.BodyParser(&request)
	id := ctx.Params("id")
	err := cntrl.Nodes.AddWordsToTrie(id, request.Words)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	msg := fmt.Sprintf("%d words have been added", len(request.Words))
	return ctx.JSON(fiber.Map{ "message": msg })
}

func (cntrl *SearchController) AutoComplete (ctx *fiber.Ctx) error {
	request := SearchTrieRequest{}
	ctx.BodyParser(&request)
	id := ctx.Params("id")
	suggestions, err := cntrl.Nodes.SearchTrie(id, request.Prefix)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{ "suggestions": suggestions })
}

