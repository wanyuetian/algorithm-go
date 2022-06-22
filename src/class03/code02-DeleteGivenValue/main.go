package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
	Prev  *ListNode
}

func removeValue(head *ListNode, num int) *ListNode {
	// head来到第一个不需要删除的位置
	for head != nil {
		if head.Value != num {
			break
		}
		head = head.Next
	}
	cur, prev := head, head
	for cur != nil {
		if cur.Value == num {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}
	return head
}

func main() {
	head := &ListNode{
		Value: 1,
		Next: &ListNode{
			Value: 2,
			Next:  &ListNode{Value: 3},
		},
	}
	head = removeValue(head, 2)
	fmt.Println(head)
}
