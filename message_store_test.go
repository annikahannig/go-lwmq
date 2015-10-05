package lwmq

import (
	"testing"
)

func TestMemoryMessageStore(t *testing.T) {
	store := NewMemoryMessageStore()

	q := NewQueue("foo")
	q.SetStore(store)

	q.AddMessage(NewMessage("foo1"))
	q.AddMessage(NewMessage("foo2"))
	q.AddMessage(NewMessage("bar1"))

	// Message store should now contain 3 messages
	msgs, err := store.PopMessages(q)

	if err != nil {
		t.Error(err)
	}

	if len(msgs) != 3 {
		t.Error("There should be 3 message in the store. Received only:", len(msgs))
	}

	msgs, err = store.PopMessages(q)
	if len(msgs) != 0 {
		t.Error("There should be no message in the store. Received:", len(msgs))
	}

}
