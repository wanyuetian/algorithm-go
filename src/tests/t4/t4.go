package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if nums == nil || len(nums) == 0 {
		return res
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if nums[len(nums)-1] < 0 || nums[0] > 0 {
		return res
	}

	lastI := math.MinInt
	for i := 0; i < len(nums)-2; i++ {
		if lastI == nums[i] {
			continue
		}
		lastJ := math.MinInt
		for j := i + 1; j < len(nums)-1; {
			if nums[j] == lastJ {
				j++
				continue
			}
			tmp := 0 - nums[i] - nums[j]
			left := j + 1
			right := len(nums) - 1
			for nums[left] < tmp && left < right {
				left++
			}
			for nums[right] > tmp && left < right {
				right--
			}
			if nums[left] == tmp {
				res = append(res, []int{nums[i], nums[j], tmp})
			}

			lastJ = nums[j]
			j++
		}
		lastI = nums[i]
	}
	return res
}

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
	m := make(map[int]struct{})
	m1 := &m
	fmt.Println(len(*m1))
}
