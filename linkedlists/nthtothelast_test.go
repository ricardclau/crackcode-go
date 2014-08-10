package linkedlists

import (
	"testing"
)

func TestEdgeCases(t *testing.T) {
	s := &SingleLinkedList{}
	if getNToTheLast(2, s) != nil {
		t.Error("Empty LinkedList -> nil not satisfied")
	}

	s.Push(1)
	s.Push(2)

	if getNToTheLast(0, s) != nil {
		t.Error("Nth < 1 -> nil not satisfied")
	}

	if getNToTheLast(4, s) != nil {
		t.Error("Asking for a Nth larger than List Size -> nil not satisfied")
	}
}

func TestNthToTheLast(t *testing.T) {
	s := &SingleLinkedList{}
	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	if nth := getNToTheLast(3, s); nth != 2 {
		t.Errorf("Unexpected value %d", nth)
	}

	if nth := getNToTheLast(1, s); nth != 0 {
		t.Errorf("Unexpected value %d", nth)
	}

	if nth := getNToTheLast(99, s); nth != 98 {
		t.Errorf("Unexpected value %d", nth)
	}
}
