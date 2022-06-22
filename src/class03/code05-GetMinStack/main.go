package main

import (
	"errors"
	"fmt"
)

type MinStack struct {
	Stack []int
	Min   []int
}

func NewMinStack() *MinStack {
	return &MinStack{
		Stack: make([]int, 0),
		Min:   make([]int, 0),
	}
}

func (s *MinStack) Push(value int) {
	if s.IsEmpty() {
		s.Stack = append(s.Stack, value)
		s.Min = append(s.Min, value)
	} else {
		s.Stack = append(s.Stack, value)
		if value < s.Min[len(s.Min)-1] {
			s.Min = append(s.Min, value)
		} else {
			s.Min = append(s.Min, s.Min[len(s.Min)-1])
		}
	}
}

func (s *MinStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	v := s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[0 : len(s.Stack)-1]
	s.Min = s.Min[0 : len(s.Min)-1]
	return v, nil
}

func (s *MinStack) Peak() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.Stack[len(s.Stack)-1], nil
}

func (s *MinStack) GetMin() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.Min[len(s.Min)-1], nil
}

func (s *MinStack) IsEmpty() bool {
	return len(s.Stack) == 0
}

func main() {
	stack1 := NewMinStack()
	PrintMin(stack1)
	stack1.Push(3)
	PrintMin(stack1)
	stack1.Push(4)
	PrintMin(stack1)
	stack1.Push(1)
	PrintMin(stack1)
	stack1.Pop()
	PrintMin(stack1)

}

func PrintMin(stack *MinStack) {
	if v, err := stack.GetMin(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}

/*
	stack1.push(3);
	System.out.println(stack1.getmin());
	stack1.push(4);
	System.out.println(stack1.getmin());
	stack1.push(1);
	System.out.println(stack1.getmin());
	System.out.println(stack1.pop());
	System.out.println(stack1.getmin());

	System.out.println("=============");

	MyStack1 stack2 = new MyStack1();
	stack2.push(3);
	System.out.println(stack2.getmin());
	stack2.push(4);
	System.out.println(stack2.getmin());
	stack2.push(1);
	System.out.println(stack2.getmin());
	System.out.println(stack2.pop());
	System.out.println(stack2.getmin());
*/
