package main

import (
	"testing"
)

func TestMemberCreation(t *testing.T) {
	member, err := NewMember("Alice")
	if err != nil {
		t.Fatalf("Failed to create member: %v", err)
	}
	if member.Name != "Alice" {
		t.Errorf("Member name is incorrect, got: %s, want: %s", member.Name, "Alice")
	}
}

func TestGroupCreation(t *testing.T) {
	group, err := NewGroup("TestGroup")
	if err != nil {
		t.Fatalf("Failed to create group: %v", err)
	}
	if group.Name != "TestGroup" {
		t.Errorf("Group name is incorrect, got: %s, want: %s", group.Name, "TestGroup")
	}
}

func TestAddMember(t *testing.T) {
	group, _ := NewGroup("TestGroup")
	member, _ := NewMember("Alice")
	err := group.AddMember(member)
	if err != nil {
		t.Fatalf("Failed to add member to group: %v", err)
	}
	if len(group.Members) != 1 {
		t.Errorf("Number of group members is incorrect, got: %d, want: %d", len(group.Members), 1)
	}
}

func TestMessageEncryptionDecryption(t *testing.T) {
	member, _ := NewMember("Alice")
	message := "Hello, World!"
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

func TestRemoveMember(t *testing.T) {
	group, _ := NewGroup("TestGroup")
	member, _ := NewMember("Alice")
	group.AddMember(member)

	err := group.RemoveMember(member)
	if err != nil {
		t.Fatalf("Failed to remove member from group: %v", err)
	}

	if group.Root != nil {
		t.Errorf("Expected group root to be nil, got '%v'", group.Root)
	}

	if group.Epoch != 2 {
		t.Errorf("Expected group epoch to be 2, got '%d'", group.Epoch)
	}
}
