package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func binarySearch(arr []int, num int) int {
	if arr == nil || len(arr) < 1 {
		return -1
	}
	L := 0
	R := len(arr) - 1

	return process(arr, L, R, num)
}

func process(arr []int, L, R, num int) int {
	if L > R {
		return -1
	}
	mid := L + (R-L)>>1
	if arr[mid] == num {
		return mid
	} else if arr[mid] > num {
		R = mid - 1
	} else {
		L = mid + 1
	}
	return process(arr, L, R, num)
}

func findNum(arr []int, num int) int {
	if arr == nil || len(arr) < 1 {
		return -1
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
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
	maxSize := 100
	maxVal := 100

	for i := 0; i < testTimes; i++ {
		arr := generateRandomArray(maxSize, maxVal)
		sort.Ints(arr)
		target := rand.Intn(maxVal)
		ans1 := binarySearch(arr, target)
		ans2 := findNum(arr, target)
		if ans1 != ans2 && arr[ans1] != arr[ans2] {
			fmt.Println("Something wrong...")
			fmt.Println(arr)
			return
		}
	}
	fmt.Println("Nice algorithm...")
}
