package main

import (
	"testing"
)

func TestNewMember(t *testing.T) {
	member, err := NewMember("Test Member")
	if err != nil {
		t.Fatalf("Failed to create new member: %v", err)
	}
	if member.Name != "Test Member" {
		t.Errorf("Member name is incorrect, got: %s, want: %s", member.Name, "Test Member")
	}
	if member.Keys == nil {
		t.Errorf("Member keys should not be nil")
	}
}

func TestEncryptDecryptMessage(t *testing.T) {
	member, _ := NewMember("Test Member")
	message := "Hello 42Snippets"
	encryptedMessage, err := member.EncryptMessage(message)
	if err != nil {
		t.Fatalf("Failed to encrypt message: %v", err)
	}
	decryptedMessage, err := member.DecryptMessage(encryptedMessage)
	if err != nil {
		t.Fatalf("Failed to decrypt message: %v", err)
	}
	if decryptedMessage != message {
		t.Errorf("Decrypted message is incorrect, got: %s, want: %s", decryptedMessage, message)
	}
}
