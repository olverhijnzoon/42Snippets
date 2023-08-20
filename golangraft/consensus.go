package main

import (
	"math/rand"
	"sync"
	"time"
)

// Raft states
const (
	Follower = iota
	Candidate
	Leader
)

// Node represents a node in the Raft cluster.
type Node struct {
	ID   int
	Addr string // Address of the node for communication
}

// LogEntry represents an entry in the Raft log.
type LogEntry struct {
	Term    int
	Command string
}

// Consensus represents the core Raft consensus module.
type Consensus struct {
	mu                sync.Mutex
	State             int
	CurrentTerm       int
	VotedFor          int
	LeaderID          int
	ElectionTimeout   time.Duration
	HeartbeatInterval time.Duration
	voteCount         int
}

// Assuming a list of all nodes in the cluster:
var nodes = []Node{
	// TODO: Add nodes here
	// Example: {ID: 1, Addr: "localhost:8001"},
	//          {ID: 2, Addr: "localhost:8002"},
}

func NewConsensus() *Consensus {
	return &Consensus{
		State:             Follower,
		CurrentTerm:       0,
		VotedFor:          -1,
		ElectionTimeout:   time.Duration(rand.Intn(150)+150) * time.Millisecond,
		HeartbeatInterval: 100 * time.Millisecond,
	}
}

// StartElection starts the leader election process.
func (c *Consensus) StartElection() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.State = Candidate
	c.CurrentTerm++
	c.VotedFor = -1
	c.voteCount = 0

	// Request votes from other nodes
	// TODO: Implement logic to send vote requests to other nodes

	// If received majority votes, become leader
	if c.voteCount > len(nodes)/2 { // Assuming nodes is a list of all nodes in the cluster
		c.becomeLeader()
	}
}

// HandleVoteRequest handles incoming vote requests from other nodes.
func (c *Consensus) HandleVoteRequest(term int, candidateID int) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if term <= c.CurrentTerm {
		return false
	}

	if c.VotedFor == -1 || c.VotedFor == candidateID {
		c.VotedFor = candidateID
		return true
	}

	return false
}

// becomeLeader sets the current node as the leader.
func (c *Consensus) becomeLeader() {
	c.State = Leader
	c.LeaderID = -1 // Set to the current node's ID
	// TODO: Start sending heartbeats to other nodes
}

// AppendEntries handles log replication from the leader.
// Placeholder for now.
func (c *Consensus) AppendEntries(term int, leaderID int, entries []LogEntry) bool {
	// TODO: Implement log replication logic
	return false
}
