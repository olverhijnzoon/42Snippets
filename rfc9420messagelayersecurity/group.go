package main

import (
	"errors"
)

type Group struct {
	Name    string
	Members []*Member
	State   string
}

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
	g.Members = append(g.Members, member)
	return nil
}

func (g *Group) RemoveMember(member *Member) error {
	for i, m := range g.Members {
		if m == member {
			g.Members = append(g.Members[:i], g.Members[i+1:]...)
			return nil
		}
	}
	return errors.New("member not found in group")
}
