package linkedlists

import (
	"strconv"
	"strings"
)

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
