package structture

// Stack 栈
type Stack struct {
	elements []interface{}
}

func (s *Stack) Push(val interface{}) {
	s.elements = append(s.elements, val)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.elements) == 0 {
		return nil, false
	}

	l := len(s.elements)
	top := s.elements[l-1]
	s.elements = s.elements[:l-1]
	return top, true
}

func (s *Stack) Top() (interface{}, bool) {
	if len(s.elements) == 0 {
		return nil, false
	}

	return s.elements[len(s.elements)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}
