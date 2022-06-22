package main

import (
	"fmt"
	"math/rand"
	"time"
)

func netherlandsFlag(nums []int, L, R int) []int {
	if L > R {
		return []int{-1, -1}
	}
	if L == R {
		return []int{L, R}
	}
	lessIndex := L - 1
	moreIndex := R
	index := L
	res := make([]int, 2)
	for index < moreIndex {
		if nums[index] < nums[R] {
			nums[index], nums[lessIndex+1] = nums[lessIndex+1], nums[index]
			lessIndex++
			index++
		} else if nums[index] > nums[R] {
			nums[index], nums[moreIndex-1] = nums[moreIndex-1], nums[index]
			moreIndex--
		} else {
			index++
		}
	}

	nums[moreIndex], nums[R] = nums[R], nums[moreIndex]
	res[0] = lessIndex + 1
	res[1] = moreIndex
	return res
}

func quickSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	process(nums, 0, len(nums)-1)
}

func process(nums []int, L int, R int) {
	if L >= R {
		return
	}
	M := netherlandsFlag(nums, L, R)
	process(nums, L, M[0]-1)
	process(nums, M[1]+1, R)
}

func quickSortUnRecursive(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}

	res := netherlandsFlag(nums, 0, len(nums)-1)
	stack := NewStack()
	stack.Push(NewOp(0, res[0]-1))
	stack.Push(NewOp(res[1]+1, len(nums)-1))

	for !stack.IsEmpty() {
		op := stack.Pop()
		if op.Left < op.Right {
			randomPos := rand.Intn(op.Right-op.Left+1) + op.Left
			nums[op.Right], nums[randomPos] = nums[randomPos], nums[op.Right]
			res = netherlandsFlag(nums, op.Left, op.Right)
			stack.Push(NewOp(op.Left, res[0]-1))
			stack.Push(NewOp(res[1]+1, op.Right))
		}
	}
}

type Op struct {
	Left  int
	Right int
}

func NewOp(L, R int) *Op {
	return &Op{
		Left:  L,
		Right: R,
	}
}

type Stack struct {
	Ops []*Op
}

func NewStack() *Stack {
	return &Stack{Ops: make([]*Op, 0)}
}

func (s *Stack) Push(op *Op) {
	s.Ops = append(s.Ops, op)
}

func (s *Stack) Pop() *Op {
	op := s.Ops[len(s.Ops)-1]
	s.Ops = s.Ops[:len(s.Ops)-1]
	return op
}

func (s *Stack) IsEmpty() bool {
	return len(s.Ops) == 0
}

func copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	copyArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		copyArr[i] = arr[i]
	}
	return copyArr
}

func isEqual(arr1, arr2 []int) bool {
	if arr1 == nil && arr2 == nil {
		return true
	}
	if (arr1 == nil && arr2 != nil) || (arr1 != nil && arr2 == nil) {
		return false
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func generateRandomArray(maxSize, maxValue int) []int {
	maxSize = rand.Intn(maxSize + 1)
	arr := make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		arr[i] = rand.Intn(maxValue + 1)
	}
	return arr
}

func main() {
	rand.Seed(time.Now().UnixNano())
	testTimes := 100000
	maxSize := 10
	maxVal := 100
	for i := 0; i < testTimes; i++ {
		arr := generateRandomArray(maxSize, maxVal)
		arr1 := copyArray(arr)
		arr2 := copyArray(arr)
		quickSort(arr1)
		quickSortUnRecursive(arr2)
		if !isEqual(arr1, arr2) {
			fmt.Println("something wrong")
			fmt.Println(arr)
			fmt.Println(arr1)
			fmt.Println(arr2)
			return
		}
	}

	fmt.Println("Nice algorithm")
}
