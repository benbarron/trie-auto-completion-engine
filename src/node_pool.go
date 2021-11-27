package main

import (
	"fmt"
	"github.com/google/uuid"
)

type NodePool struct {
	Nodes map[string]*Node
}

func CreateNodePool() *NodePool {
	return &NodePool{
		Nodes: make(map[string]*Node),
	}
}

func (np *NodePool) AddNodeToPool() string {
	var id string = uuid.New().String()
	np.Nodes[id] = CreateNode()
	return id
}

func (np *NodePool) AddWordsToTrie(id string, words []string) error {
	node, exists := np.Nodes[id]
	if !exists {
		return fmt.Errorf("node does not exist with id: %s", id)
	}
	InsertMultipleTexts(node, words)
	return nil
}

func (np *NodePool) SearchTrie(id string, prefix string) ([]string, error) {
	node, exists := np.Nodes[id]
	if !exists {
		return nil, fmt.Errorf("node does not exist with id: %s", id)
	}
	prefixes := AutoCompletePrefix(node, prefix)
	return prefixes, nil
}
