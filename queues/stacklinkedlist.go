package queues

type StackLinkedList struct {
	top *node
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.top == nil
}

func (s *StackLinkedList) Push(item interface{}) {
	n := &node{item: item}
	n.next = s.top
	s.top = n
}

func (s *StackLinkedList) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	old := s.top
	s.top = old.next

	return old.item
}
