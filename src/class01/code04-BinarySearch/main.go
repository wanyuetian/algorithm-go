package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func binarySearch(arr []int, num int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	L, R := 0, len(arr)-1
	var mid int
	for L <= R {
		mid = L + ((R - L) >> 1)
		if arr[mid] == num {
			return mid
		}
		if arr[mid] > num {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}

	if arr[mid] == num {
		return mid
	}
	return -1
}

func search(arr []int, num int) int {
	if arr == nil || len(arr) == 0 {
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
	maxSize := 10
	maxVal := 100

	for i := 0; i < testTimes; i++ {
		arr := generateRandomArray(maxSize, maxVal)
		//arr := []int{10, 15, 43, 60, 80, 83}

		sort.Ints(arr)
		num := rand.Intn(maxVal + 1)
		//num := 60
		pos1 := binarySearch(arr, num)
		pos2 := search(arr, num)
		if pos1 != pos2 && arr[pos1] != arr[pos2] {
			fmt.Println("Something wrong...")
			fmt.Println(pos1, pos2)
			fmt.Println(arr)
			fmt.Println(num)
			return
		}
	}
	fmt.Println("Nice algorithm...")
}
