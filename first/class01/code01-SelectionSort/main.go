package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func selectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	N := len(arr)

	for i := 0; i < N-1; i++ {
		minIndex := i
		for j := i + 1; j < N; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
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
		selectionSort(arr1)
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
