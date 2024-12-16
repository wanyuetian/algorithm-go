package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func nearestIndex(arr []int, value int) int {
	if arr == nil || len(arr) < 1 {
		return -1
	}
	L := 0
	R := len(arr) - 1
	index := -1
	for L <= R {
		mid := L + (R-L)>>1
		if arr[mid] >= value {
			index = mid
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return index
}

func search(arr []int, value int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] >= value {
			index = i
			break
		}
	}
	return index
}

func generateRandomArray(maxSize, maxVal int) []int {
	maxSize = rand.Intn(maxSize + 1)
	arr := make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		arr[i] = rand.Intn(maxVal + 1)
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
		//arr := []int{10, 15, 43, 60, 80, 83}

		sort.Ints(arr)
		value := rand.Intn(maxVal + 1)
		//num := 60
		pos1 := nearestIndex(arr, value)
		pos2 := search(arr, value)
		if pos1 != pos2 && arr[pos1] != arr[pos2] {
			fmt.Println("Something wrong...")
			fmt.Println(pos1, pos2)
			fmt.Println(arr)
			fmt.Println(value)
			return
		}
	}
	fmt.Println("Nice algorithm...")
}