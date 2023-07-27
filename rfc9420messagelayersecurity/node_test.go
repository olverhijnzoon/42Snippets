package main

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func TestNode(t *testing.T) {
	// Generate a random RSA key pair
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Create a new node with the public key
	node := NewNode(&key.PublicKey)
	if node.PublicKey != &key.PublicKey {
		t.Fatalf("Node public key does not match input public key")
	}

	// Generate another random RSA key pair
	key2, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Insert the new public key into the tree
	node.Insert(&key2.PublicKey)
	if node.Left == nil || node.Left.PublicKey != &key2.PublicKey {
		t.Fatalf("Failed to insert new public key into tree")
	}
}
