package lwmq

import (
	"time"
)

/**
 * A message can be passed to a number of consumers
 * or be persisted in a MessageStore for later delivery.
 */
type Message struct {
	Payload   interface{}
	Timestamp time.Time
	Offline   bool
	Queue     *Queue
}

/**
 * Create a new message from a payload string
 */
func NewMessage(payload interface{}) *Message {
	m := &Message{
		Payload:   payload,
		Timestamp: time.Now(),
	}
	return m
}
