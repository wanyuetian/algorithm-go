package main

import (
	"errors"
	"fmt"
)

// arr中，只有一种数，出现奇数次
func findOddTimesNum1(arr []int) (int, error) {
	if arr == nil || len(arr) == 0 {
		return 0, errors.New("invalid arr")
	}
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}
	return eor, nil
}

// arr中，有两种数，出现奇数次
func findOddTimesNum2(arr []int) (int, int, error) {
	if arr == nil || len(arr) < 2 {
		return 0, 0, errors.New("invalid arr")
	}

	eor := 0
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}

	// 提取最右侧的1
	rightOne := eor & (-eor)

	onlyOne := 0
	for i := 0; i < len(arr); i++ {
		if (arr[i] & rightOne) != 0 {
			onlyOne ^= arr[i]
		}
	}
	return onlyOne, onlyOne ^ eor, nil
}

func main() {
	arr1 := []int{3, 3, 2, 3, 1, 1, 1, 3, 1, 1, 1}
	num, _ := findOddTimesNum1(arr1)
	fmt.Println(num)

	arr2 := []int{4, 3, 4, 2, 2, 2, 4, 1, 1, 1, 3, 3, 1, 1, 1, 4, 2, 2}
	num1, num2, _ := findOddTimesNum2(arr2)
	fmt.Println(num1, num2)
}
