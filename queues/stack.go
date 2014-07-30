package queues

type node struct {
	item interface{}
	next *node
}

type Stack struct {
	top *node
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Push(item interface{}) {
	n := &node{item: item}
	n.next = s.top
	s.top = n
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	old := s.top
	s.top = old.next

	return old.item
}
