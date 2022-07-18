package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

// need n extra space
func isPalindrome1(head *Node) bool {
	stack := make([]*Node, 0)
	cur := head
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Next
	}
	cur = head
	index := len(stack) - 1
	for cur != nil {
		if cur.Val != stack[index].Val {
			return false
		}
		cur = cur.Next
		index--
	}
	return true
}

// need n/2 extra space
func isPalindrome2(head *Node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	stack := make([]*Node, 0)
	for slow != nil {
		stack = append(stack, slow)
		slow = slow.Next
	}
	//cur := head
	for i := len(stack) - 1; i >= 0; i-- {
		if head.Val != stack[i].Val {
			return false
		}
		head = head.Next
	}
	return true
}

// need O(1) extra space
func isPalindrome3(head *Node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	rightPartFirst := slow.Next
	slow.Next = nil
	pre := slow
	for rightPartFirst != nil {
		next := rightPartFirst.Next
		rightPartFirst.Next = pre
		pre = rightPartFirst
		rightPartFirst = next
	}
	last := pre
	first := head
	res := true

	for first != nil && last != nil {
		if first.Val != last.Val {
			res = false
			break
		}
		first = first.Next
		last = last.Next
	}

	// recover linklist
	cur := slow
	pre = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return res
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
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	ans1 := isPalindrome1(head)
	ans2 := isPalindrome2(head)
	ans3 := isPalindrome2(head)
	fmt.Println(ans1, ans2, ans3)
	head = &Node{
		Val: 0,
		Next: &Node{
			Val: 1,
			Next: &Node{
				Val: 2,
				Next: &Node{
					Val: 3,
					Next: &Node{
						Val: 2,
						Next: &Node{
							Val: 1,
							Next: &Node{
								Val:  0,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	ans1 = isPalindrome1(head)
	ans2 = isPalindrome2(head)
	ans3 = isPalindrome2(head)
	fmt.Println(ans1, ans2, ans3)
}
