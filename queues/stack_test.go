package queues

import (
	"testing"
)

func TestEmptyStack(t *testing.T) {
	s := Stack{}
	if !s.IsEmpty() {
		t.Error("IsEmpty not working properly")
	}
}

func TestPushPopOneElement(t *testing.T) {
	s := Stack{}
	s.Push(1)

	if s.IsEmpty() {
		t.Error("Push did not work properly")
	}

	if a := s.Pop(); a != 1 {
		t.Errorf("Unexpected value %s", a)
	}

	if !s.IsEmpty() {
		t.Error("Pop did not work properly")
	}
}

func TestPushPopManyElements(t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push("aaaa")
	s.Push(12.45)
	s.Push(5)

	if a := s.Pop(); a != 5 {
		t.Errorf("Unexpected value %s", a)
	}

	if a := s.Pop(); a != 12.45 {
		t.Errorf("Unexpected value %s", a)
	}

	if a := s.Pop(); a != "aaaa" {
		t.Errorf("Unexpected value %s", a)
	}

	if a := s.Pop(); a != 2 {
		t.Errorf("Unexpected value %s", a)
	}

	if a := s.Pop(); a != 1 {
		t.Errorf("Unexpected value %s", a)
	}

	if !s.IsEmpty() {
		t.Error("Pop did not work properly")
	}
}
