package main

import (
	"fmt"
	"math/rand"
)

func mergeSort1(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	L := 0
	R := len(arr) - 1
	process(arr, L, R)
}

func process(arr []int, L, R int) {
	if L == R {
		return
	}
	mid := L + (R-L)/2
	process(arr, L, mid)
	process(arr, mid+1, R)
	merge(arr, L, mid, R)
}

func merge(arr []int, L int, mid int, R int) {
	help := make([]int, R-L+1)
	p1 := L
	p2 := mid + 1
	index := 0
	for p1 <= mid && p2 <= R {
		if arr[p1] <= arr[p2] {
			help[index] = arr[p1]
			p1++
		} else {
			help[index] = arr[p2]
			p2++
		}
		index++
	}
	for p1 <= mid {
		arr[index] = arr[p1]
		p1++
	}
	for p2 <= R {
		arr[p2] = arr[p2]
		p2++
	}
	for i := 0; i < len(help); i++ {
		arr[L] = help[i]
		L++
	}
}

func mergeSort2(arr []int) {
	if arr == nil || len(arr) == 0 {
		return
	}
	N := len(arr)
	mergeSize := 1
	for mergeSize < N {
		L := 0
		for L < N {
			M := L + mergeSize - 1
			if M >= N {
				break
			}
			R := M + min(mergeSize, N-M-1)
			merge(arr, L, M, R)
			L = R + 1
		}
		if mergeSize > N/2 {
			break
		}
		mergeSize <<= 1
	}

}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func generateRandomArray(maxVal, maxSize int) []int {
	maxSize = rand.Intn(maxSize + 1)
	arr := make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		arr[i] = rand.Intn(maxVal + 1)
	}
	return arr
}

func copyArr(arr []int) []int {
	if arr == nil {
		return nil
	}
	tmpArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		tmpArr[i] = arr[i]
	}
	return tmpArr
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

func main() {
	testTimes := 100000
	maxVal := 100
	maxSize := 10

	for i := 0; i < testTimes; i++ {
		arr := generateRandomArray(maxVal, maxSize)
		arr1 := generateRandomArray(maxVal, maxSize)
		arr2 := copyArr(arr1)
		mergeSort1(arr1)
		mergeSort2(arr2)
		if !isEqual(arr1, arr2) {
			fmt.Println("Something wrong...")
			fmt.Println(arr)
			fmt.Println(arr1)
			fmt.Println(arr2)
			return
		}
	}
	fmt.Println("Nice algorithm...")
}
