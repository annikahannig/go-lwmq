package lwmq

import (
	"fmt"
	"log"
	"testing"
)

type LogConsumer struct{}

func (self *LogConsumer) HandleMessage(msg *Message) error {
	log.Println("[", msg.Queue.Id, "] @ ", msg.Timestamp, ": ", msg.Payload)
	return nil
}

type CountConsumer struct {
	count map[string]uint64
}

func NewCountConsumer() *CountConsumer {
	c := &CountConsumer{
		count: make(map[string]uint64),
	}
	return c
}

func (self *CountConsumer) HandleMessage(msg *Message) error {
	self.count[msg.Queue.Id] += 1
	return nil
}

func (self *CountConsumer) PrintStats() {
	for id, count := range self.count {
		fmt.Println(id, ":", count)
	}
}

func TestMessageConsuming(t *testing.T) {

	// Create new queue consumer
	logConsumer := &LogConsumer{}
	countConsumer := NewCountConsumer()

	q := NewQueue("pb_23")

	// Post some messages
	q.AddMessage(NewMessage("This is a test."))
	q.AddMessage(NewMessage("This is another test."))

	// Add message consumers
	err := q.AddConsumer(logConsumer)
	if err != nil {
		log.Fatal("Could not add consumer:", err)
	}

	err = q.AddConsumer(countConsumer)
	if err != nil {
		log.Fatal("Could not add consumer:", err)
	}

	q.Deliver()

	// Remove count consumer
	q.RemoveConsumer(countConsumer)

	q.AddMessage(NewMessage("This is yet another test."))

	countConsumer.PrintStats()

}
