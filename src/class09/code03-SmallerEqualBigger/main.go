package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func listPartition1(head *Node, pivot int) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	var sH, sT, eH, eT, mH, mT *Node
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = nil
		if cur.Val < pivot {
			if sH != nil {
				sT.Next = cur
				sT = sT.Next
			} else {
				sH = cur
				sT = cur
			}
		} else if cur.Val == pivot {
			if eH != nil {
				eT.Next = cur
				eT = eT.Next
			} else {
				eH = cur
				eT = cur
			}
		} else {
			if mH != nil {
				mT.Next = cur
				mT = mT.Next
			} else {
				mH = cur
				mT = cur
			}
		}
		cur = next
	}
	var newHead *Node
	if sH != nil {
		newHead = sH
	}
	if newHead == nil && eH != nil {
		newHead = eH
	}
	if newHead == nil && mH != nil {
		newHead = mH
	}
	if sT != nil {
		sT.Next = eH
	}
	if eT != nil {
		eT.Next = mH
	}
	if sT != nil && eT == nil {
		sT.Next = mH
	}
	return newHead
}

func listPartition2(head *Node, pivot int) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	arr := make([]*Node, 0)
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = nil
		arr = append(arr, cur)
		cur = next
	}

	var sH, sT, eH, eT, mH, mT *Node
	for i := 0; i < len(arr); i++ {
		cur = arr[i]
		if cur.Val < pivot {
			if sH != nil {
				sT.Next = cur
				sT = sT.Next
			} else {
				sH = cur
				sT = cur
			}
		} else if cur.Val == pivot {
			if eH != nil {
				eT.Next = cur
				eT = eT.Next
			} else {
				eH = cur
				eT = cur
			}
		} else {
			if mH != nil {
				mT.Next = cur
				mT = mT.Next
			} else {
				mH = cur
				mT = cur
			}
		}
	}
	var newHead *Node
	if sH != nil {
		newHead = sH
	}
	if newHead == nil && eH != nil {
		newHead = eH
	}
	if newHead == nil && mH != nil {
		newHead = mH
	}
	if sT != nil {
		sT.Next = eH
	}
	if eT != nil {
		eT.Next = mH
	}
	if sT != nil && eT == nil {
		sT.Next = mH
	}
	return newHead
	return nil
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
	printLinklist(head)
	newHead1 := listPartition1(head, 0)
	newHead2 := listPartition2(head, 0)
	printLinklist(newHead1)
	printLinklist(newHead2)
	newHead1 = listPartition1(head, 2)
	newHead2 = listPartition2(head, 2)
	printLinklist(newHead1)
	printLinklist(newHead2)
	newHead1 = listPartition1(head, 3)
	newHead2 = listPartition2(head, 3)
	printLinklist(newHead1)
	printLinklist(newHead2)
}

func printLinklist(head *Node) {
	if head == nil {
		return
	}
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("--------")
}
