package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	arr []int
}

func NewQueue() *Queue {
	return &Queue{arr: make([]int, 0)}
}

func (q *Queue) Add(v int) {
	q.arr = append(q.arr, v)
}

func (q *Queue) Poll() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	v := q.arr[0]
	q.arr = q.arr[1:]
	return v, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
}

type TwoQueueStack struct {
	Q    *Queue
	Help *Queue
}

func NewTwoQueueStack() *TwoQueueStack {
	return &TwoQueueStack{
		Q:    NewQueue(),
		Help: NewQueue(),
	}
}

func (s *TwoQueueStack) Push(v int) {
	s.Q.Add(v)
}

func (s *TwoQueueStack) Pop() (int, error) {
	if s.Q.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	length := len(s.Q.arr)
	for i := 0; i < length-1; i++ {
		v, err := s.Q.Poll()
		if err != nil {
			return 0, err
		}
		s.Help.Add(v)
	}
	v, err := s.Q.Poll()
	if err != nil {
		return 0, err
	}
	s.Help, s.Q = s.Q, s.Help
	return v, nil
}

func (s *TwoQueueStack) IsEmpty() bool {
	return s.Q.IsEmpty()
}

func main() {
	s := NewTwoQueueStack()
	printPop(s)
	fmt.Println(s.IsEmpty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Println(s.IsEmpty())
	printPop(s)
	printPop(s)
	printPop(s)
	printPop(s)
	printPop(s)
	printPop(s)
	printPop(s)
}

func printPop(s *TwoQueueStack) {
	v, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}
