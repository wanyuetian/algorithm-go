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

func isCBT1(head *Node) bool {
	if head == nil {
		return true
	}
	queue := []*Node{head}
	leaf := false
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		left := cur.Left
		right := cur.Right
		if (left == nil && right != nil) || (leaf && (left != nil || right != nil)) {
			return false
		}
		if left != nil {
			queue = append(queue, left)
		}
		if right != nil {
			queue = append(queue, right)
		}
		if left == nil || right == nil {
			leaf = true
		}
	}
	return true
}

type Info struct {
	IsFull, IsCBT bool
	Height        int
}

func isCBT2(head *Node) bool {
	if head == nil {
		return true
	}
	return process(head).IsCBT
}

func process(head *Node) *Info {
	if head == nil {
		return &Info{
			IsFull: true,
			IsCBT:  true,
			Height: 0,
		}
	}
	leftInfo := process(head.Left)
	rightInfo := process(head.Right)
	isFull, isCBT := false, false
	height := max(leftInfo.Height, rightInfo.Height) + 1
	if leftInfo.IsFull && rightInfo.IsFull && leftInfo.Height == rightInfo.Height {
		isFull = true
	}
	if isFull {
		return &Info{
			IsFull: true,
			IsCBT:  true,
			Height: height,
		}
	}

	if leftInfo.IsFull && rightInfo.IsFull && leftInfo.Height == rightInfo.Height+1 {
		isCBT = true
	}
	if leftInfo.IsFull && rightInfo.IsCBT && leftInfo.Height == rightInfo.Height {
		isCBT = true
	}
	if leftInfo.IsCBT && rightInfo.IsFull && leftInfo.Height == rightInfo.Height+1 {
		isCBT = true
	}
	return &Info{
		IsFull: isFull,
		IsCBT:  isCBT,
		Height: height,
	}
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
		r1 := isCBT1(head)
		r2 := isCBT2(head)
		if r1 != r2 {
			fmt.Println("Oops")
		}
	}
	fmt.Println("Nice")
}
