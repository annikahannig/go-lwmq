package lwmq

import (
	"time"
)

/**
 * A message can be passed to a number of consumers
 * or be persisted in a MessageStore for later delivery.
 */
type Message struct {
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
	Offline   bool        `json:"offline"`
	Queue     *Queue      `json:"-"`
}

/**
 * Create a new message with a payload
 * and current time
 */
func NewMessage(payload interface{}) *Message {
	m := &Message{
		Payload:   payload,
		Timestamp: time.Now(),
	}
	return m
}
