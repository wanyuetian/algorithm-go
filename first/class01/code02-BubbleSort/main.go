package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	N := len(arr)
	for i := 1; i < N; i++ {
		for j := 0; j < N-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func generateRandomArray(maxSize, maxVal int) []int {
	maxSize = rand.Intn(maxSize + 1)
	arr := make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		arr[i] = rand.Intn(maxVal + 1)
	}
	return arr
}

func copyArray(arr []int) []int {
	newArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		newArr[i] = arr[i]
	}
	return newArr
}

func isEqual(arr1, arr2 []int) bool {
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
	rand.Seed(time.Now().UnixNano())
	testTimes := 100000
	maxSize := 100
	maxVal := 100

	for i := 0; i < testTimes; i++ {
		arr1 := generateRandomArray(maxSize, maxVal)
		arr2 := copyArray(arr1)
		bubbleSort(arr1)
		sort.Ints(arr2)
		if !isEqual(arr1, arr2) {
			fmt.Println("Something wrong...")
			fmt.Println(arr1)
			fmt.Println(arr2)
			return
		}
	}
	fmt.Println("Nice algorithm...")
}
