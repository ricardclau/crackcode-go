package queues

import (
	"testing"
)

func TestEmptyStackRA(t *testing.T) {
	s := NewStackResizingArray()
	if !s.IsEmpty() {
		t.Error("IsEmpty not working properly")
	}
}

func TestPushPopOneElementRA(t *testing.T) {
	s := NewStackResizingArray()
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

func TestPushPopManyElementsRA(t *testing.T) {
	s := NewStackResizingArray()
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

func TestResizeHappensAsExpected(t *testing.T) {
	s := NewStackResizingArray()
	if len(s.elements) != 2 {
		t.Errorf("Size of elements is not 8 but %v", len(s.elements))
	}

	s.Push(1)
	s.Push(2)
	if len(s.elements) != 4 {
		t.Errorf("Size of elements is not 8 but %v", len(s.elements))
	}

	s.Push(3)
	s.Push(4)
	if len(s.elements) != 8 {
		t.Errorf("Size of elements is not 8 but %v", len(s.elements))
	}

	s.Pop()
	if len(s.elements) != 8 {
		t.Errorf("Size of elements is not 8 but %v", len(s.elements))
	}

	s.Pop()
	if len(s.elements) != 4 {
		t.Errorf("Size of elements is not 8 but %v", len(s.elements))
	}
}
