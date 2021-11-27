package main

import (
	"unicode/utf8"
)

type Node struct {
	children	map[string]*Node
	end 		bool
}

func CreateNode() *Node {
	return &Node{
		children: make(map[string]*Node),
		end: false,
	}
}

func InsertText(root *Node, text string) {
	if root == nil {
		return
	}
	if text == "" {
		root.end = true
		return
	}

	key := text[0:1]
	if _, ok := root.children[key]; !ok {
		root.children[key] = CreateNode()
	}
	_, i := utf8.DecodeRuneInString(text)
	InsertText(root.children[key], text[i:])
}

func InsertMultipleTexts(root *Node, texts []string) {
	for _, text := range texts {
		InsertText(root, text)
	}
}

func FindSubTrieByPrefix(root *Node, prefix string) *Node {
	if prefix == "" {
		return root
	}
	if root == nil {
		return nil
	}
	key := prefix[0:1]
	if _, exists := root.children[key]; !exists {
		return root
	}
	_, i := utf8.DecodeRuneInString(prefix)
	return FindSubTrieByPrefix(root.children[key], prefix[i:])
}

func SearchSubTree(root *Node, texts *[]string, prefix string) {
	if root == nil {
		return
	}
	for k, node := range root.children {
		var newPrefix string = prefix + k
		if node.end {
			*texts = append(*texts, newPrefix)
		}
		SearchSubTree(node, texts, newPrefix)
	}
}

func AutoCompletePrefix(root *Node, prefix string) []string {
	subTrie := FindSubTrieByPrefix(root, prefix)
	var prefixes []string
	SearchSubTree(subTrie, &prefixes, prefix)
	return prefixes
}