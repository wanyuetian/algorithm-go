package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func countSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	max := math.MinInt32
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	help := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		help[arr[i]]++
	}
	index := 0
	for i := 0; i < len(help); i++ {
		for help[i] > 0 {
			arr[index] = i
			help[i]--
			index++
		}
	}
}

func generateRandomArray(maxVal int, maxSize int) []int {
	maxSize = rand.Intn(maxSize + 1)
	arr := make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		arr[i] = rand.Intn(maxVal + 1)
	}
	return arr
}

func copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	newArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		newArr[i] = arr[i]
	}
	return newArr
}

func isEqual(arr1, arr2 []int) bool {
	if arr1 == nil && arr2 == nil {
		return true
	}
	if arr1 == nil {
		return false
	}
	if arr2 == nil {
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
	maxVal := 200
	maxSize := 100
	for i := 0; i < testTimes; i++ {
		arr1 := generateRandomArray(maxVal, maxSize)
		arr2 := copyArray(arr1)
		countSort(arr1)
		sort.Ints(arr2)
		if !isEqual(arr1, arr2) {
			fmt.Println("Oops")
			return
		}
	}
	fmt.Println("Nice")
}
