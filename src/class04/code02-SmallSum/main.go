package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func smallSum(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return process(arr, 0, len(arr)-1)
}

func process(arr []int, L int, R int) int {
	if L == R {
		return 0
	}
	M := L + (R-L)>>1
	return process(arr, L, M) + process(arr, M+1, R) + merge(arr, L, M, R)
}

func merge(arr []int, L int, M int, R int) int {
	p1, p2 := L, M+1
	index := 0
	help := make([]int, R-L+1)
	sum := 0
	for p1 <= M && p2 <= R {
		if arr[p1] < arr[p2] {
			help[index] = arr[p1]
			sum += arr[p1] * (R - p2 + 1)
			p1++
		} else {
			help[index] = arr[p2]
			p2++
		}
		index++
	}
	for p1 <= M {
		help[index] = arr[p1]
		p1++
		index++
	}

	for p2 <= R {
		help[index] = arr[p2]
		p2++
		index++
	}
	for i := 0; i < len(help); i++ {
		arr[L] = help[i]
		L++
	}
	return sum
}

func generateRandomArray(maxSize, maxValue int) []int {
	maxSize = rand.Intn(maxSize + 1)
	arr := make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		arr[i] = rand.Intn(maxValue + 1)
	}
	return arr
}

func directlyCalculating(arr []int) int {
	sum := 0
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				sum += arr[j]
			}
		}
	}
	sort.Ints(arr)
	return sum
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

func main() {
	testTimes := 100000
	maxSize := 4
	maxVale := 20

	for i := 0; i < testTimes; i++ {
		arr1 := generateRandomArray(maxSize, maxVale)
		//arr1 := []int{1, 2, 3, 4}
		arr2 := copyArray(arr1)
		ans1 := smallSum(arr1)
		ans2 := directlyCalculating(arr2)
		if ans1 != ans2 || !isEqual(arr1, arr2) {
			fmt.Println("something wrong")
			fmt.Println(arr1)
			fmt.Println(ans1)
			fmt.Println(arr2)
			fmt.Println(ans2)
			return
		}
	}
	fmt.Println("Nice algorithm")
}
