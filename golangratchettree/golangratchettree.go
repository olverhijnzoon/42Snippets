package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

type Node struct {
	Left           *Node
	Right          *Node
	Parent         *Node
	IsLeaf         bool
	NodeData       string
	PrivateKey     *rsa.PrivateKey
	PublicKey      rsa.PublicKey
	Credential     string // Only for leaf nodes
	UnmergedLeaves []int  // An ordered list of "unmerged" leaves
	ParentHash     string // A hash of certain information
	Index          int
}

type RatchetTree struct {
	Root      *Node
	LeafCount int
}

func (t *RatchetTree) insert(data string, node *Node) *Node {
	// Generate a new RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	if t.Root == nil {
		t.Root = &Node{NodeData: data, IsLeaf: true, PrivateKey: privateKey, PublicKey: *publicKey, Index: t.LeafCount}
		t.LeafCount++
		return t.Root
	}

	if node.IsLeaf {
		node.IsLeaf = false
		node.Left = &Node{Parent: node, NodeData: data, IsLeaf: true, PrivateKey: privateKey, PublicKey: *publicKey, Index: t.LeafCount}
		t.LeafCount++
		return node.Left
	}

	if node.Right == nil {
		node.Right = &Node{Parent: node, NodeData: data, IsLeaf: true, PrivateKey: privateKey, PublicKey: *publicKey, Index: t.LeafCount}
		t.LeafCount++
		return node.Right
	}

	if t.height(node.Left) > t.height(node.Right) {
		return t.insert(data, node.Right)
	} else {
		return t.insert(data, node.Left)
	}
}

func (t *RatchetTree) markUnmerged(leafIndex int, ancestorIndex int) {
	leaf := t.findNodeByIndex(leafIndex)
	if leaf == nil {
		fmt.Println("Leaf node not found")
		return
	}
	leaf.UnmergedLeaves = append(leaf.UnmergedLeaves, ancestorIndex)
}

func (t *RatchetTree) findNodeByIndex(index int) *Node {
	return t.findNodeByIndexRecursive(t.Root, index)
}

func (t *RatchetTree) findNodeByIndexRecursive(node *Node, index int) *Node {
	if node == nil {
		return nil
	}
	if node.Index == index {
		return node
	}
	foundNode := t.findNodeByIndexRecursive(node.Left, index)
	if foundNode != nil {
		return foundNode
	}
	return t.findNodeByIndexRecursive(node.Right, index)
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

	// Mark the leaf node at index 5 as "unmerged" relative to its ancestor at index 2
	tree.markUnmerged(6, 2)

	// Find the node at index 5 and print its data
	node := tree.findNodeByIndex(6)
	if node != nil {
		fmt.Println("Found node:", node.NodeData)
	}

	tree.printTree(tree.Root, 0)
}
