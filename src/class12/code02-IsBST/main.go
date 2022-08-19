package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func isBST1(head *Node) bool {
	if head == nil {
		return true
	}
	arr := make([]*Node, 0)
	var in func(head *Node)
	in = func(head *Node) {
		if head == nil {
			return
		}
		in(head.Left)
		arr = append(arr, head)
		in(head.Right)
	}

	in(head)
	if len(arr) == 1 {
		return true
	}
	for i := 1; i < len(arr); i++ {
		if arr[i-1].Val >= arr[i].Val {
			return false
		}
	}
	return true
}

type Info struct {
	IsB      bool
	Min, Max int
}

func isBST2(head *Node) bool {
	if head == nil {
		return true
	}
	return process(head).IsB
}

func process(node *Node) *Info {
	if node == nil {
		return nil
	}
	leftInfo := process(node.Left)
	rightInfo := process(node.Right)
	if (leftInfo != nil && !leftInfo.IsB) || (rightInfo != nil && !rightInfo.IsB) {
		return &Info{IsB: false}
	}

	if (leftInfo != nil && leftInfo.Max >= node.Val) || (rightInfo != nil && rightInfo.Min <= node.Val) {
		return &Info{IsB: false}
	}

	minNum := node.Val
	maxNum := node.Val
	if leftInfo != nil {
		minNum = min(leftInfo.Min, minNum)
		maxNum = max(leftInfo.Max, maxNum)
	}

	if rightInfo != nil {
		minNum = min(rightInfo.Min, minNum)
		maxNum = max(rightInfo.Max, maxNum)
	}

	return &Info{
		IsB: true,
		Min: minNum,
		Max: maxNum,
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func generateRandomBST(maxLevel, maxValue int) *Node {
	return generate(1, maxLevel, maxValue)
}

func generate(level int, maxLevel int, maxValue int) *Node {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &Node{Val: rand.Intn(maxValue) + 1}
	head.Left = generate(level+1, maxLevel, maxValue)
	head.Right = generate(level+1, maxLevel, maxValue)
	return head
}

func main() {
	rand.Seed(time.Now().UnixNano())
	maxLevel := 10
	maxValue := 100
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		head := generateRandomBST(maxLevel, maxValue)
		r1 := isBST1(head)
		r2 := isBST2(head)
		if r1 != r2 {
			fmt.Println("Oops")
		}
	}
	fmt.Println("Nice")
}
