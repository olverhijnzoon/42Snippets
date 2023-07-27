package main

import (
	"errors"
	"sync"
)

type Storage struct {
	groups map[string]*Group
	mu     sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		groups: make(map[string]*Group),
	}
}

func (s *Storage) StoreGroup(group *Group) error {
	if group == nil {
		return errors.New("group cannot be nil")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.groups[group.Name] = group
	return nil
}

func (s *Storage) RetrieveGroup(name string) (*Group, error) {
	if name == "" {
		return nil, errors.New("group name cannot be empty")
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	group, ok := s.groups[name]
	if !ok {
		return nil, errors.New("group not found")
	}
	return group, nil
}
