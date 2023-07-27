package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type ProtocolState struct {
	Group *Group
}

type ProtocolMessage struct {
	Type    string
	Member  *Member
	Content string
}

func NewProtocolState(group *Group) *ProtocolState {
	return &ProtocolState{
		Group: group,
	}
}

func (ps *ProtocolState) HandleProtocolMessage(message *ProtocolMessage) error {
	if message == nil {
		return errors.New("message cannot be nil")
	}
	switch message.Type {
	case "Proposal":
		return ps.HandleProposalMessage(message)
	case "Commit":
		return ps.HandleCommitMessage(message)
	case "Welcome":
		return ps.HandleWelcomeMessage(message)
	default:
		return errors.New("unknown message type")
	}
}

func (ps *ProtocolState) HandleProposalMessage(message *ProtocolMessage) error {
	// Check that the message is not nil
	if message == nil {
		return errors.New("message cannot be nil")
	}

	// Check that the message type is "Proposal"
	if message.Type != "Proposal" {
		return errors.New("message type must be 'Proposal'")
	}

	// Check that the member is not nil
	if message.Member == nil {
		return errors.New("member cannot be nil")
	}

	// Check that the content is not empty
	if message.Content == "" {
		return errors.New("content cannot be empty")
	}

	// Parse the content of the message
	// In a real implementation, the content would likely be a structured data type that you would parse into a Go struct.
	// For simplicity, let's assume that the content is a string that can be "add", "remove", or "update".
	switch message.Content {
	case "add":
		// Add the member to the group
		return ps.Group.AddMember(message.Member)
	case "remove":
		// Remove the member from the group
		return ps.Group.RemoveMember(message.Member)
	case "update":
		// Update the member's key
		// In a real implementation, the new key would be included in the message.
		newKey, err := GenerateKeyPair()
		if err != nil {
			return err
		}
		message.Member.Keys = newKey
		return nil
	default:
		return errors.New("unknown proposal type")
	}
}

func (ps *ProtocolState) HandleCommitMessage(message *ProtocolMessage) error {
	// Check that the message is not nil
	if message == nil {
		return errors.New("message cannot be nil")
	}

	// Check that the message type is "Commit"
	if message.Type != "Commit" {
		return errors.New("message type must be 'Commit'")
	}

	// Check that the member is not nil
	if message.Member == nil {
		return errors.New("member cannot be nil")
	}

	// Check that the content is not empty
	if message.Content == "" {
		return errors.New("content cannot be empty")
	}

	// Parse the content of the message
	// In a real implementation, the content would likely be a structured data type that you would parse into a Go struct.
	// For simplicity, let's assume that the content is a string that represents the new group state.
	newGroupState := message.Content

	// Update the group state
	// In a real implementation, you would perform a cryptographic operation to update the group state.
	ps.Group.State = newGroupState

	return nil
}

func (ps *ProtocolState) HandleWelcomeMessage(message *ProtocolMessage) error {
	// Check that the message is not nil
	if message == nil {
		return errors.New("message cannot be nil")
	}

	// Check that the message type is "Welcome"
	if message.Type != "Welcome" {
		return errors.New("message type must be 'Welcome'")
	}

	// Check that the member is not nil
	if message.Member == nil {
		return errors.New("member cannot be nil")
	}

	// Check that the content is not empty
	if message.Content == "" {
		return errors.New("content cannot be empty")
	}

	// Parse the content of the message
	// In a real implementation, the content would likely be a structured data type that you would parse into a Go struct.
	// For simplicity, let's assume that the content is a string that represents the group's secret key.
	groupSecretKey := message.Content

	// Deserialize the group's secret key
	block, _ := pem.Decode([]byte(groupSecretKey))
	if block == nil {
		return errors.New("failed to decode PEM block")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return errors.New("failed to cast to *rsa.PrivateKey")
	}

	// Set the member's key to the group's secret key
	message.Member.Keys.PrivateKey = rsaPrivateKey

	return nil
}
