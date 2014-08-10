package linkedlists

import (
	"testing"
)

func TestDeleteNodeEdgeCases(t *testing.T) {
	if deleteNode(nil) != false {
		t.Error("DeleteNode returning true when nil passed")
	}

	list := &SingleLinkedList{}
	list.Push(1)
	list.Push(2)

	lastNode := list.top.next

	if deleteNode(lastNode) != false {
		t.Error("Apparently we could delete the last node!")
	}
}

func TestDeleteFirstNode(t *testing.T) {
	list := &SingleLinkedList{}
	list.Push(1)
	list.Push(2)
	list.Push(3)

	toBeDeleted := list.top

	if deleteNode(toBeDeleted) != true {
		t.Error("Problem deleting first node")
	}

	if NumericListToString(list) != "[2,1]" {
		t.Error("Structure different than expected after deleting first node: %v", NumericListToString(list))
	}
}

func TestDelete3rdNodeInListOf5(t *testing.T) {
	list := &SingleLinkedList{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	list.Push(5)

	toBeDeleted := list.top.next.next

	if deleteNode(toBeDeleted) != true {
		t.Error("Problem deleting 3rd node")
	}

	if NumericListToString(list) != "[5,4,2,1]" {
		t.Error("Structure different than expected after deleting first node: %v", NumericListToString(list))
	}
}
