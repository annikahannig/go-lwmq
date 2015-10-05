package lwmq

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMessage(t *testing.T) {

	// New message should create a new message
	// from a payload string, with current time as timestamp
	t0 := time.Now()
	payload := "test"

	msg := NewMessage(payload)

	if msg.Payload != payload {
		t.Error("Message payload is not the expected payload")
	}

	if msg.Timestamp.Sub(t0) > 1*time.Millisecond {
		t.Error("Expected message not to be older than a millisecond")
	}

	// Test serialization
	data, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(data))
}
