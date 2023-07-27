package main

import (
	"os"
	"testing"
)

func TestHandleProposalMessage(t *testing.T) {
	// Create a new group
	group, err := NewGroup("TestGroup")
	if err != nil {
		t.Fatalf("Failed to create group: %v", err)
	}

	// Create a new protocol state
	ps := NewProtocolState(group)

	// Create a new member
	member, err := NewMember("Alice")
	if err != nil {
		t.Fatalf("Failed to create member: %v", err)
	}

	// Create a Proposal message to add the member
	proposalMessage := &ProtocolMessage{
		Type:    "Proposal",
		Member:  member,
		Content: "add",
	}

	// Handle the Proposal message
	err = ps.HandleProposalMessage(proposalMessage)
	if err != nil {
		t.Fatalf("Failed to handle Proposal message: %v", err)
	}

	// Check that the member was added to the group
	if len(group.Members) != 1 || group.Members[0] != member {
		t.Fatalf("Member was not added to group")
	}

	// Create a Proposal message to remove the member
	proposalMessage = &ProtocolMessage{
		Type:    "Proposal",
		Member:  member,
		Content: "remove",
	}

	// Handle the Proposal message
	err = ps.HandleProposalMessage(proposalMessage)
	if err != nil {
		t.Fatalf("Failed to handle Proposal message: %v", err)
	}

	// Check that the member was removed from the group
	if len(group.Members) != 0 {
		t.Fatalf("Member was not removed from group")
	}
}

func TestHandleCommitMessage(t *testing.T) {
	// Create a new group
	group, err := NewGroup("TestGroup")
	if err != nil {
		t.Fatalf("Failed to create group: %v", err)
	}

	// Create a new protocol state
	ps := NewProtocolState(group)

	// Create a new member
	member, err := NewMember("Alice")
	if err != nil {
		t.Fatalf("Failed to create member: %v", err)
	}

	// Create a Commit message to update the group state
	commitMessage := &ProtocolMessage{
		Type:    "Commit",
		Member:  member,
		Content: "newState", // This represents the new group state
	}

	// Handle the Commit message
	err = ps.HandleCommitMessage(commitMessage)
	if err != nil {
		t.Fatalf("Failed to handle Commit message: %v", err)
	}

}

func TestHandleWelcomeMessage(t *testing.T) {
	// Create a new group
	group, err := NewGroup("TestGroup")
	if err != nil {
		t.Fatalf("Failed to create group: %v", err)
	}

	// Create a new protocol state
	ps := NewProtocolState(group)

	// Create a new member
	member, err := NewMember("Alice")
	if err != nil {
		t.Fatalf("Failed to create member: %v", err)
	}

	// Read the private key from the file
	keyBytes, err := os.ReadFile("private.pem")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	// Convert the key bytes to a string
	groupSecretKey := string(keyBytes)
	welcomeMessage := &ProtocolMessage{
		Type:    "Welcome",
		Member:  member,
		Content: groupSecretKey,
	}

	// Handle the Welcome message
	err = ps.HandleWelcomeMessage(welcomeMessage)
	if err != nil {
		t.Fatalf("Failed to handle Welcome message: %v", err)
	}

	// Check that the member's key was updated
	if member.Keys.PrivateKey == nil {
		t.Fatalf("Member's key was not updated")
	}
}
