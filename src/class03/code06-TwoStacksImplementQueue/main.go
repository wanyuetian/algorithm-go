package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Stack []int
}

func NewStack() *Stack {
	return &Stack{
		Stack: make([]int, 0),
	}
}

func (s *Stack) Push(value int) {
	s.Stack = append(s.Stack, value)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	v := s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[0 : len(s.Stack)-1]
	return v, nil
}

func (s *Stack) Peak() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.Stack[len(s.Stack)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.Stack) == 0
}

type Queue struct {
	PushStack *Stack
	PopStack  *Stack
}

func NewQueue() *Queue {
	return &Queue{
		PushStack: NewStack(),
		PopStack:  NewStack(),
	}
}
func (q *Queue) Add(value int) {
	q.PushStack.Push(value)
	q.pushToPop()
}

func (q *Queue) Poll() (int, error) {
	q.pushToPop()
	if q.PopStack.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return q.PopStack.Pop()
}

func (q *Queue) Peak() (int, error) {
	q.pushToPop()
	if !q.PopStack.IsEmpty() {
		return q.PopStack.Peak()
	}
	return 0, errors.New("queue is empty")
}

func (q *Queue) pushToPop() {
	if !q.PopStack.IsEmpty() {
		return
	}

	for !q.PushStack.IsEmpty() {
		if v, err := q.PushStack.Pop(); err != nil {
			break
		} else {
			q.PopStack.Push(v)
		}
	}
}

func main() {
	q := NewQueue()
	peakPrint(q)
	pollPrint(q)
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	q.Add(5)
	peakPrint(q)
	pollPrint(q)
	pollPrint(q)
	pollPrint(q)
	pollPrint(q)
	pollPrint(q)
	pollPrint(q)
	peakPrint(q)
}

func peakPrint(q *Queue) {
	if v, err := q.Peak(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}

func pollPrint(q *Queue) {
	if v, err := q.Poll(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}
