package main

import "fmt"

type stack struct {
	data []byte
	cur int
}

func NewStack() *stack {
	data := make([]byte, 1000)
	return &stack{
		data: data,
		cur: -1,
	}
}

func (s *stack) Push(n byte) {
	s.cur++
	s.data[s.cur] = n
}

func (s *stack) Pop() (error, byte) {
	if s.IsEmpty() {
		return fmt.Errorf("stack is empty"), 0
	}
	data := s.data[s.cur]
	s.cur--
	return nil, data
}

func (s *stack) IsEmpty() bool {
	return s.cur == -1
}