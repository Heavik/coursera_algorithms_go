package datastructs

import "errors"

type stack struct {
	top    int
	values []int
}

// Stack creates a new empty stack
func Stack() *stack {
	return &stack{top: -1}
}

func (s *stack) IsEmpty() bool {
	return s.top == -1
}

func (s *stack) Push(item int) {
	s.top++
	if s.top >= len(s.values) {
		s.values = append(s.values, item)
	} else {
		s.values[s.top] = item
	}
}

func (s *stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	value := s.values[s.top]
	s.top--
	return value, nil
}
