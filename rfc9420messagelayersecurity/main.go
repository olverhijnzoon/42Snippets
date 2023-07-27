package main

import (
	"fmt"
)

func main() {
	fmt.Println("# 42Snippets")
	fmt.Println("## RFC9420 Message Layer Security")

	/*
		Message Layer Security (MLS) is a security protocol designed to provide end-to-end encryption for group communications. It is being developed by the IETF (Internet Engineering Task Force) with contributions from several organizations.

		The main goal of MLS is to ensure that data shared within a group remains confidential and authentic, even in the face of compromises of group members or the infrastructure. It is designed to be used in applications where group members can dynamically join and leave, and where the group size can be large.

		In this snippet, we try to explore some of the aspects of MLS.
	*/

	// Create a new group
	group, err := NewGroup("42Group")
	if err != nil {
		fmt.Printf("Failed to create group: %v\n", err)
		return
	}

	// Create a new member
	member1, err := NewMember("Alice")
	if err != nil {
		fmt.Printf("Failed to create member: %v\n", err)
		return
	}

	// Add the member to the group
	err = group.AddMember(member1)
	if err != nil {
		fmt.Printf("Failed to add member to group: %v\n", err)
		return
	}

	// Print a success message
	fmt.Println("Successfully created a group and added a member!")

	// Test encryption and decryption
	message := "MLS 42Snippets"
	encryptedMessage, err := member1.EncryptMessage(message)
	if err != nil {
		fmt.Printf("Failed to encrypt message: %v\n", err)
		return
	}
	fmt.Printf("Encrypted message: %s\n", encryptedMessage)

	decryptedMessage, err := member1.DecryptMessage(encryptedMessage)
	if err != nil {
		fmt.Printf("Failed to decrypt message: %v\n", err)
		return
	}
	fmt.Printf("Decrypted message: %s\n", decryptedMessage)

	// Create a new node
	node1 := &TreeNode{
		Member: member1,
		Parent: nil,
		Left:   nil,
		Right:  nil,
	}

	// Add the node to the group's tree
	group.Root = node1

	// Update the group's epoch
	group.Epoch++

	// Print the group's state
	fmt.Printf("Group epoch: %d\n", group.Epoch)
	fmt.Printf("Group root member: %s\n", group.Root.Member.Name)

	// Create a new member
	member2, err := NewMember("Bob")
	if err != nil {
		fmt.Printf("Failed to create member: %v\n", err)
		return
	}

	// Add the member to the group
	err = group.AddMember(member2)
	if err != nil {
		fmt.Printf("Failed to add member to group: %v\n", err)
		return
	}

	// Create a new node
	node2 := &TreeNode{
		Member: member2,
		Parent: nil,
		Left:   nil,
		Right:  nil,
	}

	// Update the group's tree
	group.Root.Right = node2

	// Update the group's epoch
	group.Epoch++

	// Print the group's state
	fmt.Printf("Group epoch: %d\n", group.Epoch)
	fmt.Printf("Group root member: %s\n", group.Root.Member.Name)
	fmt.Printf("Group right member: %s\n", group.Root.Right.Member.Name)

	decryptedMessage2, err := member2.DecryptMessage(encryptedMessage)
	fmt.Printf("Earlier decrypted message for Bob: %s\n", decryptedMessage2)

	decryptedMessage10, err := member1.DecryptMessage(encryptedMessage)
	if err != nil {
		fmt.Printf("Failed to decrypt message: %v\n", err)
	}
	fmt.Printf("Decrypted message for Alice: %s\n", decryptedMessage10)

	// Remove the first member from the group
	err = group.RemoveMember(member1)
	if err != nil {
		fmt.Printf("Failed to remove member from group: %v\n", err)
		return
	}

	// Update the group's tree
	group.Root = node2

	// Update the group's epoch
	group.Epoch++

	// Print the group's state
	fmt.Printf("Group epoch: %d\n", group.Epoch)
	fmt.Printf("Group root member: %s\n", group.Root.Member.Name)
}
