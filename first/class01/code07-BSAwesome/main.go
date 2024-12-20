package main

import (
	"fmt"
	"math/rand"
)

// 局部最小值问题，相邻两数不相等
func getLessIndex(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	N := len(arr)
	if N == 1 || arr[1] > arr[0] {
		return 0
	}
	if arr[N-2] > arr[N-1] {
		return N - 1
	}
	L, R := 0, N-1
	for L <= R {
		mid := L + (R-L)>>1
		if mid > 0 && arr[mid] > arr[mid-1] {
			R = mid - 1
		} else if arr[mid] > arr[mid+1] {
			L = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func generateRandomArray(maxSize, maxVal int) []int {
	maxSize = rand.Intn(maxSize + 1)
	arr := make([]int, maxSize)
	if maxSize > 0 {
		arr[0] = rand.Intn(maxVal + 1)
	}
	for i := 1; i < maxSize; i++ {
		value := rand.Intn(maxVal + 1)
		for value == arr[i-1] {
			value = rand.Intn(maxVal + 1)
		}
		arr[i] = value
	}
	return arr
}

func isRight(arr []int, index int) bool {
	if arr == nil || len(arr) == 0 {
		return index == -1
	}
	if len(arr) == 1 {
		return index == 0
	}
	if index == 0 {
		return arr[index] < arr[index+1]
	}
	if index == len(arr)-1 {
		return arr[index] < arr[index-1]
	}
	return arr[index] < arr[index-1] && arr[index] < arr[index+1]
}

func main() {
	testTimes := 100000
	maxVal := 1000
	maxSize := 100

	for i := 0; i < testTimes; i++ {
		arr := generateRandomArray(maxSize, maxVal)
		index := getLessIndex(arr)
		if !isRight(arr, index) {
			fmt.Println("Something wrong...")
			fmt.Println(arr)
			fmt.Println(index)
			return
		}
	}
	fmt.Println("Nice algorithm...")
}
