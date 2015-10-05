package lwmq

import (
	"testing"
)

func TestMessageQueueing(t *testing.T) {

	// Create new queue consumer
	logConsumer := &LogConsumer{}
	countConsumer := NewCountConsumer()

	q := NewQueue("pb_23")

	err := q.AddConsumer(logConsumer)
	if err != nil {
		t.Error("Could not add consumer:", err)
	}

	err = q.AddConsumer(countConsumer)
	if err != nil {
		t.Error("Could not add consumer:", err)
	}

	// Post some messages
	q.AddMessage(NewMessage("This is a test."))
	q.AddMessage(NewMessage("This is another test."))

	// Remove count consumer
	q.RemoveConsumer(countConsumer)

	q.AddMessage(NewMessage("This is yet another test."))

	countConsumer.PrintStats()

}
