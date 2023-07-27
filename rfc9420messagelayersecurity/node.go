package main

import (
	"crypto/rsa"
)

/*
	The MLS protocol uses a binary tree structure, called a "ratchet tree", to manage the group and its members. Each leaf node in the tree represents a member of the group, and each non-leaf node (or "parent node") represents a shared secret that some subset of the group members know. The root of the tree represents a shared secret that all group members know.
*/

type Node struct {
	PublicKey *rsa.PublicKey
	Left      *Node
	Right     *Node
}

func NewNode(publicKey *rsa.PublicKey) *Node {
	return &Node{
		PublicKey: publicKey,
		Left:      nil,
		Right:     nil,
	}
}

func (n *Node) Insert(publicKey *rsa.PublicKey) {
	if n == nil {
		n = NewNode(publicKey)
	} else if n.Left == nil {
		n.Left = NewNode(publicKey)
	} else if n.Right == nil {
		n.Right = NewNode(publicKey)
	} else {
		n.Left.Insert(publicKey)
	}
}
