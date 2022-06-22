package main

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

type DoubleEndsQueue struct {
	Head *Node
	Tail *Node
}

func NewDoubleEndsQueue() *DoubleEndsQueue {
	return &DoubleEndsQueue{}
}

func (q *DoubleEndsQueue) addFromHead(value int) {
	cur := NewNode(value)
	if q.Head == nil {
		q.Head = cur
		q.Tail = cur
	} else {
		cur.Next = q.Head
		q.Head.Prev = cur
		q.Head = cur
	}
}

func (q *DoubleEndsQueue) addFromBottom(value int) {
	cur := NewNode(value)
	if q.Head == nil {
		q.Head = cur
		q.Tail = cur
	} else {
		cur.Prev = q.Tail
		q.Tail.Next = cur
		q.Tail = cur
	}
}

func (q *DoubleEndsQueue) popFromHead() *Node {
	if q.Head == nil {
		return nil
	}
	node := q.Head
	if q.Head == q.Tail {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = q.Head.Next
		node.Next = nil
		q.Head.Prev = nil
	}

	return node

}

func (q *DoubleEndsQueue) popFromBottom() *Node {
	if q.Head == nil {
		return nil
	}
	node := q.Tail
	if q.Head == q.Tail {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Tail = q.Tail.Prev
		node.Prev = nil
		q.Tail.Next = nil
	}

	return node
}

func (q *DoubleEndsQueue) isEmpty() bool {
	return q.Head == nil
}

type Stack struct {
	Q *DoubleEndsQueue
}

func (s *Stack) Push(value int) {
	s.Q.addFromBottom(value)
}

func (s *Stack) Pop() *Node {
	return s.Q.popFromBottom()
}

func (s *Stack) Peak() *Node {
	return s.Q.Tail
}

func (s *Stack) IsEmpty() bool {
	return s.Q.isEmpty()
}

type Queue struct {
	Q *DoubleEndsQueue
}

func (q *Queue) Push(value int) {
	q.Q.addFromBottom(value)
}

func (q *Queue) Poll(value int) *Node {
	return q.Q.popFromHead()
}

func (q *Queue) isEmpty() bool {
	return q.Q.isEmpty()
}
