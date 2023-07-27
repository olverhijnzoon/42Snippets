package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"sync/atomic"
)

var idCounter int32

type Member struct {
	ID   int32
	Name string
	Keys *KeyPair
}

func NewMember(name string) (*Member, error) {
	// Generate a new RSA key pair
	keys, err := GenerateKeyPair()
	if err != nil {
		return nil, fmt.Errorf("failed to generate RSA key pair: %v", err)
	}

	// Generate a unique ID for the member
	id := atomic.AddInt32(&idCounter, 1)

	// Create a new Member
	member := &Member{
		ID:   id,
		Name: name,
		Keys: keys,
	}

	return member, nil
}

func (m *Member) EncryptMessage(message string) (string, error) {
	// Convert the message to a byte slice
	messageBytes := []byte(message)

	// Encrypt the message using the member's public key
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, m.Keys.PublicKey, messageBytes)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt message: %v", err)
	}

	// Convert the ciphertext to a base64-encoded string
	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)

	return ciphertextBase64, nil
}

func (m *Member) DecryptMessage(message string) (string, error) {
	// Convert the base64-encoded message to a byte slice
	ciphertext, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 message: %v", err)
	}

	// Decrypt the message using the member's private key
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, m.Keys.PrivateKey, ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt message: %v", err)
	}

	// Convert the plaintext byte slice to a string
	plaintextString := string(plaintext)

	return plaintextString, nil
}
