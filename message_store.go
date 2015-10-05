package lwmq

import (
	"fmt"
)

/**
 * A message store accept messages and
 * keeps them for delivery.
 *
 * This may be in memory or in a database.
 */
type MessageStore interface {

	// Retrieve messages from the store
	Messages(*Queue) ([]*Message, error)

	// Add messages to the store
	AddMessage(*Message) error

	// Retrieve all messages and remove from store
	PopMessages(*Queue) ([]*Message, error)
}

/**
 * This is the default store and will
 * just store messages in a map.
 */
type MemoryMessageStore struct {
	messages map[*Queue][]*Message
}

func NewMemoryMessageStore() *MemoryMessageStore {
	s := &MemoryMessageStore{
		messages: make(map[*Queue][]*Message),
	}
	return s
}

/*
 * Implement store interface
 */

/**
 * Add message to store
 */
func (self *MemoryMessageStore) AddMessage(msg *Message) error {
	if msg.Queue == nil {
		err := fmt.Errorf("Can not store message without a queue")
		return err
	}

	// Mark as offline message
	msg.Offline = true

	// Save message
	self.messages[msg.Queue] = append(self.messages[msg.Queue], msg)

	return nil
}

/**
 * Retrieve all messages (but do not remove from store)
 */
func (self *MemoryMessageStore) Messages(q *Queue) ([]*Message, error) {
	messages, ok := self.messages[q]
	if !ok {
		return nil, fmt.Errorf("Unknown queue: %s", q.Id)
	}

	return messages, nil
}

/**
 * Retrieve messages and remove from store
 */
func (self *MemoryMessageStore) PopMessages(q *Queue) ([]*Message, error) {
	messages, err := self.Messages(q)
	if err != nil {
		return nil, err
	}

	// Clear messages
	self.messages[q] = []*Message{}

	return messages, nil
}
