package queues

import (
	"testing"
)

func TestEmptyQueue(t *testing.T) {
	q := Queue{}
	if !q.IsEmpty() {
		t.Error("IsEmpty not working properly")
	}

	if q.Dequeue() != nil {
		t.Error("WTF?")
	}
}

func TestEnqueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)

	if q.IsEmpty() {
		t.Error("IsEmpty not working properly")
	}

	if a := q.Dequeue(); a != 1 {
		t.Error("Enqueue Dequeue failing for 1 elem")
	}

	if !q.IsEmpty() {
		t.Error("IsEmpty not working properly after enqueue -> dequeue")
	}

	if q.Dequeue() != nil {
		t.Error("WTF?")
	}
}

func TestEnqDeqManyElements(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue("aaaa")

	if a := q.Dequeue(); a != 1 {
		t.Errorf("Unexpected value %s", a)
	}

	if a := q.Dequeue(); a != 2 {
		t.Errorf("Unexpected value %s", a)
	}

	q.Enqueue(12.45)
	q.Enqueue(5)

	if a := q.Dequeue(); a != "aaaa" {
		t.Errorf("Unexpected value %s", a)
	}

	if a := q.Dequeue(); a != 12.45 {
		t.Errorf("Unexpected value %s", a)
	}

	if a := q.Dequeue(); a != 5 {
		t.Errorf("Unexpected value %s", a)
	}

	if !q.IsEmpty() {
		t.Error("Pop did not work properly")
	}

	if q.Dequeue() != nil {
		t.Error("WTF?")
	}
}
