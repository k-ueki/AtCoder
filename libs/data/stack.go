package data

type (
	Stack struct {
		top    *node
		length int
	}

	node struct {
		value interface{}
		prev  *node
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Push(v interface{}) {
	n := &node{v, s.top}
	s.top = n
	s.length++
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *Stack) Top() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) IsEmpty() bool {
	if s.length == 0 {
		return true
	}
	return false
}
