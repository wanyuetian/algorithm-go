package main

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Value int
	Next  *ListNode
	Prev  *ListNode
}

func NewListNode(value int) *ListNode {
	return &ListNode{Value: value}
}

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseDoubleLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		cur.Prev = next
		pre = cur
		cur = next
	}
	return pre
}

func testReverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	arr[0].Next = nil
	for i := 1; i < len(arr); i++ {
		arr[i].Next = arr[i-1]
	}
	return arr[len(arr)-1]
}

func testReverseDoubleLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	arr[0].Next = nil
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		cur.Prev = nil
		cur.Next = prev
		prev.Prev = cur
		prev = cur
	}
	return arr[len(arr)-1]
}

func generateRandomLinkedList(len, value int) *ListNode {
	size := rand.Intn(len + 1)
	if size == 0 {
		return nil
	}
	head := NewListNode(rand.Intn(value + 1))
	prev := head
	size--
	for size != 0 {
		cur := NewListNode(rand.Intn(value + 1))
		prev.Next = cur
		prev = cur
		size--
	}
	return head
}

func generateRandomDoubleLinkedList(len, value int) *ListNode {
	size := rand.Intn(len + 1)
	if size == 0 {
		return nil
	}
	head := NewListNode(rand.Intn(value + 1))
	prev := head
	size--
	for size != 0 {
		cur := NewListNode(rand.Intn(value + 1))
		prev.Next = cur
		cur.Prev = prev
		prev = cur
		size--
	}
	return head
}

func getLinkedListOriginOrder(head *ListNode) []int {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Value)
		head = head.Next
	}
	return arr
}

func checkLinkedListReverse(head *ListNode, origin []int) bool {
	for i := len(origin) - 1; i >= 0; i-- {
		if origin[i] != head.Value {
			return false
		}
		head = head.Next
	}
	return true
}

func getDoubleLinkedListOriginOrder(head *ListNode) []int {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Value)
		head = head.Next
	}
	return arr
}

func checkDoubleLinkedListReverse(head *ListNode, origin []int) bool {
	var end *ListNode
	for i := len(origin) - 1; i >= 0; i-- {
		if origin[i] != head.Value {
			return false
		}
		end = head
		head = head.Next
	}
	for i := 0; i < len(origin); i++ {
		if origin[i] != end.Value {
			return false
		}
		end = end.Prev
	}
	return true
}

func main() {
	length := 50
	value := 100
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		head1 := generateRandomLinkedList(length, value)
		origin1 := getLinkedListOriginOrder(head1)
		head1 = reverseLinkedList(head1)
		if !checkLinkedListReverse(head1, origin1) {
			fmt.Println("Oops1")
		}

		head2 := generateRandomLinkedList(length, value)
		origin2 := getLinkedListOriginOrder(head2)
		head2 = testReverseLinkedList(head2)
		if !checkLinkedListReverse(head2, origin2) {
			fmt.Println("Oops2")
		}

		head3 := generateRandomDoubleLinkedList(length, value)
		origin3 := getDoubleLinkedListOriginOrder(head3)
		head3 = reverseDoubleLinkedList(head3)
		if !checkDoubleLinkedListReverse(head3, origin3) {
			fmt.Println("Oops3")
		}

		head4 := generateRandomDoubleLinkedList(length, value)
		origin4 := getDoubleLinkedListOriginOrder(head4)
		head4 = testReverseDoubleLinkedList(head4)
		if !checkDoubleLinkedListReverse(head4, origin4) {
			fmt.Println("Oops4")
		}

	}
	fmt.Println("Nice algorithm")
}
