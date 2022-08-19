package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Info struct {
	Height int
	IsB    bool
}

func isBalancedTree1(head *Node) bool {
	if head == nil {
		return true
	}
	ans := true
	leftHeight := process1(head.Left, &ans)
	rightHeight := process1(head.Right, &ans)
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		ans = false
	}

	return ans
}

func process1(node *Node, ans *bool) int {
	if !*ans || node == nil {
		return -1
	}
	leftHeight := process1(node.Left, ans)
	rightHeight := process1(node.Right, ans)
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		*ans = false
	}
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

func isBalancedTree2(head *Node) bool {
	if head == nil {
		return true
	}
	return process2(head).IsB
}

func process2(node *Node) *Info {
	if node == nil {
		return &Info{
			Height: 0,
			IsB:    true,
		}
	}
	leftInfo := process2(node.Left)
	rightInfo := process2(node.Right)
	isB := false
	var height int
	if (leftInfo.IsB && rightInfo.IsB) && math.Abs(float64(leftInfo.Height-rightInfo.Height)) < 2 {
		isB = true
	}
	height = int(math.Max(float64(leftInfo.Height), float64(rightInfo.Height))) + 1
	return &Info{
		Height: height,
		IsB:    isB,
	}
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
		r1 := isBalancedTree1(head)
		r2 := isBalancedTree2(head)
		if r1 != r2 {
			fmt.Println("Oops")
		}
	}
	fmt.Println("Nice")
}
