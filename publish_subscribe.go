package main

import (
	"errors"
	"time"
)

var (
	ErrTopicClosed = errors.New("Topic has been closed")
)

// Message
type Message string

// Topic
type Topic struct {
	Subscribers    []Authentication
	MessageHistory []struct {
		Author    string
		Message   Message
		Timestamp time.Time
	}
}

// Subscribe
func (t *Topic) Subscribe(Authentication) (Subscription, error) {
	// Implementation
}

// Unsubscribe
func (t *Topic) Unsubscribe(Subscription) error {
	// Implementation
}

// Delete
func (t *Topic) Delete() error {
	// Implementation
}

type Subscription struct {
	ch chan<- Message

	Inbox chan Message
}

// Publish
func (s *Subscription) Publish(msg Message) error {
	if _, ok := ch; !ok {
		return ErrTopicClosed
	}

	ch <- msg

	return nil
}

// Authentication
type Authentication struct {
}

func main() {
}
