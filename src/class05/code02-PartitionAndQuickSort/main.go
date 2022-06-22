package main

import (
	"fmt"
	"math/rand"
	"time"
)

// nums[L..R]上，以arr[R]位置的数做划分值
func partition(nums []int, L, R int) int {
	if L > R {
		return -1
	}
	if L == R {
		return L
	}
	lessEqual := L - 1
	index := L
	for index < R {
		if nums[index] <= nums[R] {
			nums[index], nums[lessEqual+1] = nums[lessEqual+1], nums[index]
			lessEqual++
		}
		index++
	}
	nums[R], nums[lessEqual+1] = nums[lessEqual+1], nums[R]
	return lessEqual + 1
}

func netherlandsFlag(nums []int, L, R int) []int {
	if L > R {
		return []int{-1, -1}
	}
	if L == R {
		return []int{L, R}
	}

	res := make([]int, 2)
	lessIndex, moreIndex := L-1, R
	index := L
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

func quickSort1(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	process1(nums, 0, len(nums)-1)
}

func process1(nums []int, L, R int) {
	if L >= R {
		return
	}
	M := partition(nums, L, R)
	process1(nums, L, M-1)
	process1(nums, M+1, R)
}

func quickSort2(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	process2(nums, 0, len(nums)-1)
}

func process2(nums []int, L, R int) {
	if L >= R {
		return
	}
	M := netherlandsFlag(nums, L, R)
	process1(nums, L, M[0]-1)
	process1(nums, M[1]+1, R)
}

func quickSort3(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	process3(nums, 0, len(nums)-1)
}

// 1 - 3 0 - 2
func process3(nums []int, L, R int) {
	if L >= R {
		return
	}
	randomPos := rand.Intn(R-L+1) + L
	nums[randomPos], nums[R] = nums[R], nums[randomPos]
	M := netherlandsFlag(nums, L, R)
	process3(nums, L, M[0]-1)
	process3(nums, M[1]+1, R)
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
		arr1 := generateRandomArray(maxSize, maxVal)
		arr2 := copyArray(arr1)
		arr3 := copyArray(arr1)
		quickSort1(arr1)
		quickSort2(arr2)
		quickSort3(arr3)
		if !isEqual(arr1, arr2) || !isEqual(arr2, arr3) {
			fmt.Println("something wrong")
			fmt.Println(arr1)
			fmt.Println(arr2)
			fmt.Println(arr3)
			return
		}
	}

	fmt.Println("Nice algorithm")
}
