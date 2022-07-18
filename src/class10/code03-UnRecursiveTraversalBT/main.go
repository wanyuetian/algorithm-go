package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Stack[T any] struct {
	Data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(value T) {
	if s.IsEmpty() {
		s.Data = append(s.Data, value)
	} else {
		s.Data = append(s.Data, value)
	}
}

func (s *Stack[T]) Pop() (T, error) {
	var v T
	if s.IsEmpty() {
		return v, errors.New("stack is empty")
	}
	v = s.Data[len(s.Data)-1]
	s.Data = s.Data[0 : len(s.Data)-1]
	return v, nil
}

func (s *Stack[T]) Peak() (T, error) {
	var v T
	if s.IsEmpty() {
		return v, errors.New("stack is empty")
	}
	return s.Data[len(s.Data)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.Data) == 0
}

func pre(head *Node) {
	if head == nil {
		return
	}
	cur := head
	s := NewStack[*Node]()
	s.Push(cur)
	for !s.IsEmpty() {
		v, _ := s.Pop()
		fmt.Println(v.Val)
		if v.Right != nil {
			s.Push(v.Right)
		}
		if v.Left != nil {
			s.Push(v.Left)
		}
	}
}

func in(head *Node) {
	if head == nil {
		return
	}
	cur := head
	s := NewStack[*Node]()
	for cur != nil || !s.IsEmpty() {
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		} else {
			v, _ := s.Pop()
			fmt.Println(v.Val)
			cur = v.Right
		}
	}
}

func pos(head *Node) {
	if head == nil {
		return
	}
	cur := head
	s1 := NewStack[*Node]()
	s2 := NewStack[*Node]()
	s1.Push(cur)
	for !s1.IsEmpty() {
		v, _ := s1.Pop()
		s2.Push(v)
		if v.Left != nil {
			s1.Push(v.Left)
		}
		if v.Right != nil {
			s1.Push(v.Right)
		}
	}
	for !s2.IsEmpty() {
		v, _ := s2.Pop()
		fmt.Println(v.Val)
	}
}

func main() {
	head := &Node{Val: 1}
	head.Left = &Node{Val: 2}
	head.Right = &Node{Val: 3}
	head.Left.Left = &Node{Val: 4}
	head.Left.Right = &Node{Val: 5}
	head.Right.Left = &Node{Val: 6}
	head.Right.Right = &Node{Val: 7}
	pre(head)
	fmt.Println("------------")
	in(head)
	fmt.Println("------------")
	pos(head)
}
