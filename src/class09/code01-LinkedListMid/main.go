package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func midOrUpMidNode(head *Node) *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func midOrDownMidNode(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func midOrUpMidPreNode(head *Node) *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow := head
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func midOrDownMidPreNode(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func right1(head *Node) *Node {
	if head == nil {
		return nil
	}

	arr := make([]*Node, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	return arr[(len(arr)-1)/2]
}

func right2(head *Node) *Node {
	if head == nil {
		return nil
	}

	arr := make([]*Node, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	return arr[len(arr)/2]
}

func right3(head *Node) *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	arr := make([]*Node, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	return arr[(len(arr)-3)/2]
}

func right4(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}

	arr := make([]*Node, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	return arr[(len(arr)-2)/2]
}

func main() {
	head := &Node{
		Val: 0,
		Next: &Node{
			Val: 1,
			Next: &Node{
				Val: 2,
				Next: &Node{
					Val: 3,
					Next: &Node{
						Val: 4,
						Next: &Node{
							Val: 5,
							Next: &Node{
								Val: 6,
								Next: &Node{
									Val: 7,
									Next: &Node{
										Val:  8,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	var ans1, ans2 *Node
	ans1 = midOrUpMidNode(head)
	ans2 = right1(head)
	fmt.Printf("ans1: %d\n", ans1)
	fmt.Printf("ans2: %d\n", ans2)

	ans1 = midOrDownMidNode(head)
	ans2 = right2(head)
	fmt.Printf("ans1: %d\n", ans1)
	fmt.Printf("ans2: %d\n", ans2)

	ans1 = midOrUpMidPreNode(head)
	ans2 = right3(head)
	fmt.Printf("ans1: %d\n", ans1)
	fmt.Printf("ans2: %d\n", ans2)

	ans1 = midOrDownMidPreNode(head)
	ans2 = right4(head)
	fmt.Printf("ans1: %d\n", ans1)
	fmt.Printf("ans2: %d\n", ans2)
}
