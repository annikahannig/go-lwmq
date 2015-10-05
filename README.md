
# Light Weight Message Queue

This implements a simple pubsub model.
Messages which can not be delivered are stored in a
message store.

Stores may be in memory (like the default store) or 
may persist the message in a database.

# Example

```go
// Create a new consumer
type LogConsumer struct{}

func (self *LogConsumer) HandleMessage(msg *Message) error {
  log.Println("[", msg.Queue.Id, "] @ ", msg.Timestamp, ": ", msg.Payload)
  return nil
}
```

```go
// Create queue
q := lwmq.NewQueue("pb_23")

// Post some messages
q.AddMessage(NewMessage("This is a test."))
q.AddMessage(NewMessage("This is another test."))

// Add message consumers
err := q.AddConsumer(LogConsumer)
if err != nil {
  log.Fatal("Could not add consumer:", err)
}

// Send message to consumer
q.Deliver()

// This message will be instantly delivered
q.AddMessage(NewMessage("This is a test."))

// Remove consumer
q.RemoveConsumer(countConsumer)

// This message will be saved in the store
q.AddMessage(NewMessage("This is yet another test."))
```

# License
MIT
