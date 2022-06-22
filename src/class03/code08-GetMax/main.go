package main

import (
	"errors"
	"fmt"
)

func getMax(arr []int) (int, error) {
	if arr == nil || len(arr) == 0 {
		return 0, errors.New("empty array")
	}
	return process(arr, 0, len(arr)-1), nil
}

func process(arr []int, L, R int) int {
	if L == R {
		return arr[L]
	}
	mid := L + (R-L)>>1
	leftMax := process(arr, L, mid)
	rightMax := process(arr, mid+1, R)
	return max(leftMax, rightMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 22, 3, 4, 12, 133, 412, 12}
	if maxV, err := getMax(arr); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(maxV)
	}
}
