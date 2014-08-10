package linkedlists

import (
	"strconv"
	"strings"
	"testing"
)

func TestEdgeCasesDeleteDupes(t *testing.T) {
	list := &SingleLinkedList{}
	deleteDupes(list)

	if NumericListToString(list) != "[]" {
		t.Error("Something wrong happened")
	}
}

func TestAllElementsDifferent(t *testing.T) {
	list := &SingleLinkedList{}
	for i := 0; i < 10; i++ {
		list.Push(i)
	}
	deleteDupes(list)

	if NumericListToString(list) != "[9,8,7,6,5,4,3,2,1,0]" {
		t.Error("Problems with deleting when all elements are different: %v", NumericListToString(list))
	}
}

func TestAllElementsEqual(t *testing.T) {
	list := &SingleLinkedList{}
	for i := 0; i < 10; i++ {
		list.Push(1)
	}
	deleteDupes(list)

	if NumericListToString(list) != "[1]" {
		t.Error("Problems with deleting when all elements are equal: %v", NumericListToString(list))
	}
}

func TestSomeDupesWithNoOrder(t *testing.T) {
	list := &SingleLinkedList{}
	list.Push(5)
	list.Push(2)
	list.Push(3)
	list.Push(2)
	list.Push(6)
	list.Push(3)
	list.Push(5)
	list.Push(1)
	list.Push(4)
	list.Push(4)

	deleteDupes(list)

	if NumericListToString(list) != "[4,1,5,3,6,2]" {
		t.Error("Problems with deleting when all elements are equal: %v", NumericListToString(list))
	}
}

func NumericListToString(s *SingleLinkedList) string {
	out := "["
	if !s.IsEmpty() {
		current := s.top
		for current != nil {
			out += strconv.Itoa(current.item.(int)) + ","
			current = current.next
		}
	}

	out = strings.TrimSuffix(out, ",")
	out += "]"

	return out
}
