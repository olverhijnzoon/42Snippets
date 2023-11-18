package main

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	node := NewNode(1)
	if node.ID != 1 {
		t.Errorf("Expected node ID to be 1, got %v", node.ID)
	}
	if node.State != Follower {
		t.Errorf("Expected initial node state to be Follower, got %v", node.State)
	}
}

func TestNodeTransitions(t *testing.T) {
	node := NewNode(1)

	// Test transition to Candidate
	node.TransitionToCandidate()
	if node.State != Candidate {
		t.Errorf("Expected node state to be Candidate after transition, got %v", node.State)
	}

	// Test transition back to Follower
	node.TransitionToFollower()
	if node.State != Follower {
		t.Errorf("Expected node state to be Follower after transition, got %v", node.State)
	}

	// Test transition to Leader
	node.TransitionToLeader()
	if node.State != Leader {
		t.Errorf("Expected node state to be Leader after transition, got %v", node.State)
	}
}
