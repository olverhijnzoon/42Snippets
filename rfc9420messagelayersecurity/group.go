package main

import (
	"crypto/rsa"
	"errors"
	"fmt"
)

type AddProposal struct {
	Member *Member
}

type Group struct {
	Name    string
	Members []*Member
	Epoch   uint64
	Root    *TreeNode
}

type TreeNode struct {
	Member *Member
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}

type UpdateProposal struct {
	MemberID     int32
	NewPublicKey *rsa.PublicKey
}

type PublicKey = *rsa.PublicKey

func NewGroup(name string) (*Group, error) {
	if name == "" {
		return nil, errors.New("group name cannot be empty")
	}
	return &Group{
		Name:    name,
		Members: make([]*Member, 0),
	}, nil
}

func (g *Group) AddMember(member *Member) error {
	if member == nil {
		return errors.New("member cannot be nil")
	}

	// Create an Add proposal message
	proposal := &AddProposal{
		Member: member,
	}

	// Process the proposal
	err := g.ProcessAddProposal(proposal)
	if err != nil {
		return fmt.Errorf("failed to process add proposal: %v", err)
	}

	// If the proposal is accepted, add the member to the group
	g.Members = append(g.Members, member)

	// Increment the group's epoch
	g.Epoch++

	return nil
}

func (g *Group) RemoveMember(member *Member) error {
	for i, m := range g.Members {
		if m == member {
			// Remove the member from the group
			g.Members = append(g.Members[:i], g.Members[i+1:]...)

			// Increment the epoch
			g.Epoch++

			return nil
		}
	}
	return errors.New("member not found in group")
}

func (g *Group) UpdateMemberKeys(member *Member, newPublicKey PublicKey) error {
	// Create an Update proposal message
	updateProposal := &UpdateProposal{
		MemberID:     member.ID,
		NewPublicKey: newPublicKey,
	}

	// Process the proposal
	accepted, err := g.ProcessUpdateProposal(updateProposal)
	if err != nil {
		return err
	}

	// If accepted, update the member's keys
	if accepted {
		member.Keys.PublicKey = newPublicKey
	} else {
		return errors.New("Update proposal not accepted")
	}

	return nil
}

// This function would typically involve some kind of consensus algorithm to decide whether the proposal should be accepted or not. For simplicity, let's assume that the proposal is accepted if the member is currently part of the group.
func (g *Group) ProcessUpdateProposal(proposal *UpdateProposal) (bool, error) {
	// Check if the member is part of the group
	for _, member := range g.Members {
		if member.ID == proposal.MemberID {
			// If the member is part of the group, accept the proposal
			return true, nil
		}
	}

	// If the member is not part of the group, reject the proposal
	return false, nil
}

func (g *Group) ProcessAddProposal(proposal *AddProposal) error {
	// Here you would implement the logic to process the proposal
	// For example, you could check if the member is already in the group
	for _, member := range g.Members {
		if member == proposal.Member {
			return fmt.Errorf("member %s is already in the group", member.Name)
		}
	}

	// If the member is not in the group, the proposal is accepted
	return nil
}
