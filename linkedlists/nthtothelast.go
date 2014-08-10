package linkedlists

func getNToTheLast(n int, s *SingleLinkedList) interface{} {
	if s.top == nil || n < 1 {
		return nil
	}

	// We have n1 pointer pointing to the top and n2 to N nodes below it
	n1, n2 := s.top, s.top

	for i := 0; i < n-1; i++ {
		if n2 == nil {
			return nil
		}
		n2 = n2.next
	}

	for n2.next != nil {
		n1 = n1.next
		n2 = n2.next
	}

	return n1.item
}
