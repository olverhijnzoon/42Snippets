package main

import (
	"sync"
	"time"
)

const HeartbeatInterval = 100 * time.Millisecond

type Node struct {
	mu    sync.Mutex
	ID    int
	State int
	Consensus
	// Channel to signal stopping of heartbeat
	stopHeartbeat chan bool
}

func NewNode(id int) *Node {
	return &Node{
		ID:        id,
		State:     Follower,
		Consensus: *NewConsensus(),
	}
}

func (n *Node) TransitionToCandidate() {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.StartElection()
	if n.State != Candidate {
		n.State = Candidate
	}
}

func (n *Node) SendHeartbeat() {
	// TODO: Implement sending of AppendEntries RPC to all followers.
}

func (n *Node) StartHeartbeatLoop() {
	ticker := time.NewTicker(HeartbeatInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			n.SendHeartbeat()
		// Exit the loop when a signal is received on stopHeartbeat channel
		case <-n.stopHeartbeat:
			return
		}
	}
}

// Send a signal to stop the heartbeat loop
func (n *Node) StopHeartbeat() {
	n.stopHeartbeat <- true
}

func (n *Node) TransitionToLeader() {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.becomeLeader()
	if n.State != Leader {
		n.State = Leader
		n.stopHeartbeat = make(chan bool) // Initialize the channel
		go n.StartHeartbeatLoop()         // Start sending heartbeats
	}
}

func (n *Node) TransitionToFollower() {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.State == Leader {
		n.StopHeartbeat() // Stop sending heartbeats if the node was previously a leader
	}
	n.State = Follower
}
