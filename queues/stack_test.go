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
	s.Push(&Node{Value: 1})

	if s.IsEmpty() {
		t.Error("Push did not work properly")
	}

	if a := s.Pop(); a.Value != 1 {
		t.Errorf("Unexpected value %s", a.Value)
	}

	if !s.IsEmpty() {
		t.Error("Pop did not work properly")
	}
}

func TestPushPopManyElements(t *testing.T) {
	s := Stack{}
	s.Push(&Node{Value: 1})
	s.Push(&Node{Value: 2})
	s.Push(&Node{Value: 3})
	s.Push(&Node{Value: 4})
	s.Push(&Node{Value: 5})

	if a := s.Pop(); a.Value != 5 {
		t.Errorf("Unexpected value %s", a.Value)
	}

	if a := s.Pop(); a.Value != 4 {
		t.Errorf("Unexpected value %s", a.Value)
	}

	if a := s.Pop(); a.Value != 3 {
		t.Errorf("Unexpected value %s", a.Value)
	}

	if a := s.Pop(); a.Value != 2 {
		t.Errorf("Unexpected value %s", a.Value)
	}

	if a := s.Pop(); a.Value != 1 {
		t.Errorf("Unexpected value %s", a.Value)
	}

	if !s.IsEmpty() {
		t.Error("Pop did not work properly")
	}
}
