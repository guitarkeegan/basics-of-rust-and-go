package stack

import "fmt"

type Node struct {
	value int
	prev  *Node
}

type Stack struct {
	Length int
	head   *Node
}

func (s *Stack) Peek() {
	if s.head != nil {
		fmt.Printf("%d\n", s.head.value)
	} else {
		fmt.Printf("nil\n")
	}
}

func NewStack() Stack {
	return Stack{Length: 0, head: nil}
}

func (s *Stack) Pop() *int {
	switch s.Length {
	case 0:
		{
			return nil
		}
	case 1:
		{
			s.Length--
			temp := s.head
			s.head = nil
			return &temp.value
		}
	default:
		s.Length--
		temp := s.head
		s.head = s.head.prev
		temp.prev = nil
		return &temp.value
	}
}

func (s *Stack) Push(item int) {
	s.Length++
	node := Node{value: item}
	if s.head == nil {
		s.head = &node
	} else {
		temp := s.head
		node.prev = temp
		s.head = &node
	}
}
