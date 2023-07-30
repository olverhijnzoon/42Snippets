package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

type Node struct {
	Left       *Node
	Right      *Node
	Parent     *Node
	IsLeaf     bool
	NodeData   string
	PrivateKey *rsa.PrivateKey
	PublicKey  rsa.PublicKey
}

type RatchetTree struct {
	Root *Node
}

func (t *RatchetTree) insert(data string, node *Node) *Node {
	// Generate a new RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	if t.Root == nil {
		t.Root = &Node{NodeData: data, IsLeaf: true, PrivateKey: privateKey, PublicKey: *publicKey}
		return t.Root
	}

	if node.IsLeaf {
		node.IsLeaf = false
		node.Left = &Node{Parent: node, NodeData: data, IsLeaf: true, PrivateKey: privateKey, PublicKey: *publicKey}
		return node.Left
	}

	if node.Right == nil {
		node.Right = &Node{Parent: node, NodeData: data, IsLeaf: true, PrivateKey: privateKey, PublicKey: *publicKey}
		return node.Right
	}

	if t.height(node.Left) > t.height(node.Right) {
		return t.insert(data, node.Right)
	} else {
		return t.insert(data, node.Left)
	}
}

func (t *RatchetTree) height(node *Node) int {
	if node == nil {
		return 0
	}
	leftHeight := t.height(node.Left)
	rightHeight := t.height(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func (t *RatchetTree) printTree(node *Node, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "   "
		}
		format += "x--| "
		level++
		t.printTree(node.Left, level)
		fmt.Printf(format+"%s, ", node.NodeData)
		publicKeyStr := fmt.Sprintf("%x", node.PublicKey.N.Bytes())
		fmt.Printf("%s\n", publicKeyStr[:42])
		t.printTree(node.Right, level)
	}
}

func main() {
	tree := &RatchetTree{}
	names := []string{"Neo", "Morpheus", "Trinity", "Cypher", "Tank", "Dozer", "Mouse", "Switch", "Apoc", "Seraph", "Niobe", "Persephone", "Link", "Zee", "Ghost", "Sparks"}

	for _, name := range names {
		tree.insert(name, tree.Root)
	}

	tree.printTree(tree.Root, 0)
}
