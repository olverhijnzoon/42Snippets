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
	group, err := NewGroup("TestGroup")
	if err != nil {
		fmt.Printf("Failed to create group: %v\n", err)
		return
	}

	// Create a new member
	member, err := NewMember("Alice")
	if err != nil {
		fmt.Printf("Failed to create member: %v\n", err)
		return
	}

	// Add the member to the group
	err = group.AddMember(member)
	if err != nil {
		fmt.Printf("Failed to add member to group: %v\n", err)
		return
	}

	// Print a success message
	fmt.Println("Successfully created a group and added a member!")
}
