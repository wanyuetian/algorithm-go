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
	for i := 0; i < N-1; i++ {
		for j := 0; j < N-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func generateRandomArray(maxSize, maxVal int) []int {
	size := rand.Intn(maxSize + 1)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(maxVal + 1)
	}
	return arr
}

func copyArr(arr []int) []int {
	newArr := make([]int, 0)
	newArr = append(newArr, arr...)
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
	maxSize := 10
	maxValue := 200
	for i := 0; i < testTimes; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		arr1 := copyArr(arr)
		arr2 := copyArr(arr1)
		bubbleSort(arr1)
		sort.Ints(arr2)
		if !isEqual(arr1, arr2) {
			fmt.Printf("Something wrong with test data: %v\n", arr)
			return
		}
	}

	fmt.Println("good job")

}
