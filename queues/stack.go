package queues

type Node struct {
	Value interface{}
	Next  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Push(n *Node) {
	n.Next = s.top
	s.top = n
}

func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		return nil
	}

	oldTop := s.top
	s.top = oldTop.Next

	return oldTop
}
