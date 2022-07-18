package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func f(head *Node) {
	if head == nil {
		return
	}
	// 1
	f(head.Left)
	// 2
	f(head.Right)
	// 3
}

func pre(head *Node) {
	if head == nil {
		return
	}
	// 1
	fmt.Println(head.Val)
	pre(head.Left)
	// 2
	pre(head.Right)
	// 3
}

func in(head *Node) {
	if head == nil {
		return
	}
	// 1
	in(head.Left)
	// 2
	fmt.Println(head.Val)
	in(head.Right)
	// 3
}

func pos(head *Node) {
	if head == nil {
		return
	}
	// 1
	pos(head.Left)
	// 2
	pos(head.Right)
	// 3
	fmt.Println(head.Val)
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
