package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// Node represents one node in the inverted Merkle tree structure with a pointer to its parent in the tree
type Node struct {
	hash   string
	parent *Node
}

// Tree represents an inverted Merkle tree where level 0 represents the individual public Lamport keys
// and [:][0] represents the public key
type Tree [][]Node

// NewNode will create a node from a string-encoded hash
func newNode(hash string) Node {
	return Node{hash: hash, parent: nil}
}

func makeLevel(nodes []Node) ([]Node, error) {
	if !isPowerOf2(len(nodes)) {
		return []Node{}, errors.New("Number of nodes must be a power of 2")
	}
	var level []Node

	return level, nil
}

func MakeTree([]PublicKey) (Tree, error) {

}

func isPowerOf2(n int) bool {
	switch {
	case n == 1:
		return true
	case n == 0 || n%2 != 0:
		return false
	default:
		return isPowerOf2(n / 2)
	}
}
