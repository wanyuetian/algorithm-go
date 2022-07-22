package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func levelTraversal(root *Node) {
	if root == nil {
		return
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Println(cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
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
	levelTraversal(head)
}
