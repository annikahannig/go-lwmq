package lwmq

import (
	"fmt"
)

/**
 * A queue can receive message and will broadcast the
 * message to the subscribed consumers.
 *
 * In case no consumers have subscribed to the queue,
 * the message will be persisted in a MessageStore
 */
type Queue struct {
	Id        string
	consumers []Consumer
	store     MessageStore
}

/**
 * Initialize a new queue with a given id.
 * Create a new memory message store for the queue.
 */
func NewQueue(id string) *Queue {
	queue := Queue{
		Id:    id,
		store: NewMemoryMessageStore(),
	}
	return &queue
}

func (self *Queue) SetStore(store MessageStore) {
	self.store = store
}

func (self *Queue) AddConsumer(c Consumer) error {
	self.consumers = append(self.consumers, c)

	return nil
}

func (self *Queue) RemoveConsumer(consumer Consumer) error {
	var (
		index int
		found bool
	)

	// Find consumer in queue
	for i, c := range self.consumers {
		if c == consumer {
			index = i
			found = true
		}
	}

	if found == false {
		err := fmt.Errorf("Could not find consumer in queue")
		return err
	}

	// Splice slice and remove consumer
	self.consumers[index] = self.consumers[len(self.consumers)-1]
	self.consumers = self.consumers[:len(self.consumers)-1]

	return nil
}

func (self *Queue) AddMessage(msg *Message) error {
	// Assign queue to message
	msg.Queue = self

	// Check if there are any consumers
	if len(self.consumers) == 0 {
		// Mark message as offline and deliver later
		msg.Offline = true
	}

	err := self.store.AddMessage(msg)
	if err != nil {
		return err
	}

	// Deliver message to consumers
	err = self.Deliver()

	return err
}

func (self *Queue) Deliver() error {
	// Don't do anything if there aren't any consumers
	if len(self.consumers) == 0 {
		return nil
	}

	// Get messages for distribution
	messages, err := self.store.PopMessages(self)
	if err != nil {
		return err
	}

	// Broadcast message
	for _, msg := range messages {
		for _, c := range self.consumers {
			c.HandleMessage(msg)
		}
	}
	return nil
}
