package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 在线测试链接 : https://leetcode.com/problems/largest-bst-subtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Info struct {
	MaxBSTSubtreeSize int
	AllSize           int
	Max               int
	Min               int
}

func largestBSTSubtree(head *TreeNode) int {
	if head == nil {
		return 0
	}
	return process(head).MaxBSTSubtreeSize
}

func process(head *TreeNode) *Info {
	if head == nil {
		return nil
	}
	leftInfo := process(head.Left)
	rightInfo := process(head.Right)

	var maxBSTSubtreeSize, allSize, max, min = 0, 1, head.Val, head.Val
	if leftInfo != nil {
		allSize += leftInfo.AllSize
		min = int(math.Min(float64(min), float64(leftInfo.Min)))
		max = int(math.Max(float64(max), float64(leftInfo.Max)))
	}
	if rightInfo != nil {
		allSize += rightInfo.AllSize
		min = int(math.Min(float64(min), float64(rightInfo.Min)))
		max = int(math.Max(float64(max), float64(rightInfo.Max)))
	}
	var isLeftBST, isRightBST bool
	if leftInfo == nil || leftInfo.AllSize == leftInfo.MaxBSTSubtreeSize {
		isLeftBST = true
	}
	if rightInfo == nil || rightInfo.AllSize == rightInfo.MaxBSTSubtreeSize {
		isRightBST = true
	}
	if isLeftBST && isRightBST {
		var leftMaxLessX, rightMinMoreX = true, true
		if leftInfo != nil && leftInfo.Max >= head.Val {
			leftMaxLessX = false
		}
		if rightInfo != nil && rightInfo.Min <= head.Val {
			rightMinMoreX = false
		}
		maxBSTSubtreeSize = 1
		if leftMaxLessX && rightMinMoreX {
			if leftInfo != nil {
				maxBSTSubtreeSize += leftInfo.MaxBSTSubtreeSize
			}
			if rightInfo != nil {
				maxBSTSubtreeSize += rightInfo.MaxBSTSubtreeSize
			}
		} else {
			if leftInfo != nil {
				maxBSTSubtreeSize = int(math.Max(float64(maxBSTSubtreeSize), float64(leftInfo.MaxBSTSubtreeSize)))
			}
			if rightInfo != nil {
				maxBSTSubtreeSize = int(math.Max(float64(maxBSTSubtreeSize), float64(rightInfo.MaxBSTSubtreeSize)))
			}
		}
	} else {
		maxBSTSubtreeSize = 1
		if leftInfo != nil {
			maxBSTSubtreeSize = int(math.Max(float64(maxBSTSubtreeSize), float64(leftInfo.MaxBSTSubtreeSize)))
		}
		if rightInfo != nil {
			maxBSTSubtreeSize = int(math.Max(float64(maxBSTSubtreeSize), float64(rightInfo.MaxBSTSubtreeSize)))
		}
	}
	return &Info{
		MaxBSTSubtreeSize: maxBSTSubtreeSize,
		AllSize:           allSize,
		Max:               max,
		Min:               min,
	}
}

func right(head *TreeNode) int {
	if head == nil {
		return 0
	}
	size := getBSTSize(head)
	if size != 0 {
		return size
	}
	return int(math.Max(float64(right(head.Left)), float64(right(head.Right))))
}

func getBSTSize(head *TreeNode) int {
	if head == nil {
		return 0
	}
	arr := make([]int, 0)
	var in func(node *TreeNode)
	in = func(node *TreeNode) {
		if node == nil {
			return
		}
		in(node.Left)
		arr = append(arr, node.Val)
		in(node.Right)
	}

	in(head)
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			return 0
		}
	}
	return len(arr)
}

func generateRandomBST(maxLevel, maxValue int) *TreeNode {
	return generate(1, maxLevel, maxValue)
}

func generate(level int, maxLevel int, maxValue int) *TreeNode {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &TreeNode{Val: rand.Intn(maxValue) + 1}
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
		r1 := largestBSTSubtree(head)
		r2 := right(head)
		if r1 != r2 {
			fmt.Println("Oops")
			largestBSTSubtree(head)
		}
	}
	fmt.Println("Nice")
}
