package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
	Prev  *ListNode
}

func removeValue(head *ListNode, num int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil && cur.Value == num {
		cur = cur.Next
	}
	newHead := cur
	pre := cur
	for cur != nil {
		next := cur.Next
		if cur.Value == num {
			pre.Next = next
		} else {
			pre = cur
		}
		cur = next
	}

	return newHead
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
