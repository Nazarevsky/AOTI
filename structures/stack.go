package structures

import (
	"container/list"
	"errors"
	"fmt"
)

type Stack struct {
	list *list.List
}

func NewStack() Stack {
	stackList := list.New()
	stackList.PushBack("A")
	return Stack{
		list: stackList,
	}
}

func (s *Stack) Push(nterminal string) {
	s.list.PushBack(nterminal)
}

func (s *Stack) Pop() (string, error) {
	if s.list.Len() == 0 {
		return "", errors.New("the stack is empty")
	}

	elem := s.list.Back()
	s.list.Remove(elem)

	return elem.Value.(string), nil
}

func (s *Stack) LastElement() string {
	if s.list.Len() != 0 {
		return s.list.Back().Value.(string)
	}
	return ""
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) AddArray(elements []string) {
	for i := len(elements) - 1; i >= 0; i-- {
		s.Push(elements[i])
	}
}

func (s *Stack) Print() {
	output := ""
	for e := s.list.Back(); e != nil; e = e.Prev() {
		output += fmt.Sprintf("%v ", e.Value)
	}
	fmt.Printf("[%s]\n", output)
}
