
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

func (self *LogConsumer) HandleMessage(msg *lwmq.Message) error {
  log.Println("[", msg.Queue.Id, "] @ ", msg.Timestamp, ": ", msg.Payload)
  return nil
}
```

```go
// Create queue
q := lwmq.NewQueue("pb_23")

// Post some messages
q.AddMessage(lwmq.NewMessage("This is a test."))
q.AddMessage(lwqm.NewMessage("This is another test."))

// Add message consumer
logConsumer := &LogConsumer{}
err := q.AddConsumer(logConsumer)
if err != nil {
  log.Fatal("Could not add consumer:", err)
}

// Send messages to consumer
q.Deliver()

// This message will be instantly delivered
q.AddMessage(lwmq.NewMessage("This is a test."))

// Remove consumer
q.RemoveConsumer(logConsumer)

// This message will be saved in the store
q.AddMessage(lwmq.NewMessage("This is yet another test."))
```

# License
MIT
