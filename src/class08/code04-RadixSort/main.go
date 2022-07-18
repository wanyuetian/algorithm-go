package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

// only for no-negative value

func radixSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	digit := maxBits(arr)
	radix := 10
	help := make([]int, len(arr))
	L, R := 0, len(arr)-1

	for d := 1; d <= digit; d++ {
		counter := make([]int, radix) // [0..9]
		for i := L; i <= R; i++ {
			j := getDigit(arr[i], d)
			counter[j]++
		}

		for i := 1; i < radix; i++ {
			counter[i] += counter[i-1]
		}

		for i := R; i >= L; i-- {
			j := getDigit(arr[i], d)
			help[counter[j]-1] = arr[i]
			counter[j]--
		}
		for i, j := L, 0; i <= R; i++ {
			arr[i] = help[j]
			j++
		}
	}

}

func getDigit(x int, d int) int {
	// (x / ((int) Math.pow(10, d - 1))) % 10
	return (x / (int)(math.Pow(10.0, float64(d-1)))) % 10
}

func maxBits(arr []int) int {
	max := math.MinInt
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	res := 0
	for max != 0 {
		res++
		max /= 10
	}
	return res
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
		radixSort(arr1)
		sort.Ints(arr2)
		if !isEqual(arr1, arr2) {
			fmt.Println("Oops")
			return
		}
	}
	fmt.Println("Nice")
}
