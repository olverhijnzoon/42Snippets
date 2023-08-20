package main

import (
	"testing"
)

func TestNewConsensus(t *testing.T) {
	c := NewConsensus()
	if c.State != Follower {
		t.Errorf("Expected initial state to be Follower, got %v", c.State)
	}
	if c.CurrentTerm != 0 {
		t.Errorf("Expected initial term to be 0, got %v", c.CurrentTerm)
	}
	if c.VotedFor != -1 {
		t.Errorf("Expected initial VotedFor to be -1, got %v", c.VotedFor)
	}
}

func TestStartElection(t *testing.T) {
	c := NewConsensus()
	c.StartElection()

	if c.State != Candidate {
		t.Errorf("Expected state to be Candidate after starting election, got %v", c.State)
	}
	if c.CurrentTerm != 1 {
		t.Errorf("Expected term to increment after starting election, got %v", c.CurrentTerm)
	}
}

func TestHandleVoteRequest(t *testing.T) {
	c := NewConsensus()

	// Test with lower term
	result := c.HandleVoteRequest(0, 1)
	if result {
		t.Errorf("Expected to reject vote request with lower term")
	}

	// Test with same term and no prior vote
	result = c.HandleVoteRequest(1, 1)
	if !result {
		t.Errorf("Expected to accept vote request with same term and no prior vote")
	}

	// Test with same term but already voted
	result = c.HandleVoteRequest(1, 2)
	if result {
		t.Errorf("Expected to reject vote request with same term but already voted")
	}
}

func TestBecomeLeader(t *testing.T) {
	c := NewConsensus()
	c.becomeLeader()

	if c.State != Leader {
		t.Errorf("Expected state to be Leader after becoming leader, got %v", c.State)
	}
	if c.LeaderID != -1 { // Assuming -1 is the current node's ID
		t.Errorf("Expected LeaderID to be set to current node's ID after becoming leader, got %v", c.LeaderID)
	}
}
