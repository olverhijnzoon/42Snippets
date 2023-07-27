package main

import (
	"errors"
)

type Message struct {
	Content string
	Sender  *Member
}

func NewMessage(content string, sender *Member) (*Message, error) {
	if content == "" {
		return nil, errors.New("message content cannot be empty")
	}
	if sender == nil {
		return nil, errors.New("message sender cannot be nil")
	}
	return &Message{
		Content: content,
		Sender:  sender,
	}, nil
}
