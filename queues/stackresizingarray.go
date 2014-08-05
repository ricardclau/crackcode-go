package queues

type StackResizingArray struct {
	elements []interface{}
	count    int
}

func NewStackResizingArray() *StackResizingArray {
	return &StackResizingArray{elements: make([]interface{}, 2)}
}

func (s *StackResizingArray) IsEmpty() bool {
	return s.count == 0
}

func (s *StackResizingArray) Push(item interface{}) {
	s.elements[s.count] = item
	s.count++
	if s.count == len(s.elements) {
		s.resize(s.count * 2)
	}
}

func (s *StackResizingArray) Pop() interface{} {
	s.count--
	elem := s.elements[s.count]
	s.elements[s.count] = nil
	if s.count == len(s.elements)/4 {
		s.resize(len(s.elements) / 2)
	}

	return elem
}

func (s *StackResizingArray) resize(newsize int) {
	tmp := make([]interface{}, newsize)
	for i := 0; i < s.count; i++ {
		tmp[i] = s.elements[i]
	}

	s.elements = tmp
}
