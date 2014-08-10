package linkedlists

type SingleLinkedList struct {
	top *singlelinkednode
}

func (s *SingleLinkedList) IsEmpty() bool {
	return s.top == nil
}

func (s *SingleLinkedList) Push(item interface{}) {
	n := &singlelinkednode{item: item}
	n.next = s.top
	s.top = n
}

func (s *SingleLinkedList) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	old := s.top
	s.top = old.next

	return old.item
}
