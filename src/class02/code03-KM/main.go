package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 输入一定能够保证，数组中所有的数都出现了M次，只有一种数出现了K次
// 1 <= K < M
// 返回这种数
func onlyKTimes(arr []int, k, m int) int {
	help := make([]int, 64)
	hashes := make(map[int]int)
	initHashMap(hashes)

	for i := 0; i < len(arr); i++ {
		num := arr[i]
		for num != 0 {
			rightOne := num & (-num)
			help[hashes[rightOne]]++
			num ^= rightOne
		}
	}

	ans := 0
	for i := 0; i < 64; i++ {
		if help[i]%m != 0 {
			ans |= 1 << i
		}
	}
	return ans
}

func km(arr []int, k, m int) int {
	help := make([]int, 64)
	for i := 0; i < len(arr); i++ {
		num := arr[i]
		for j := 0; j < 64; j++ {
			help[j] += (num >> j) & 1
		}
	}
	ans := 0
	for i := 0; i < 64; i++ {
		help[i] = help[i] % m
		if help[i] != 0 {
			ans |= 1 << i
		}
	}
	return ans
}

func test(arr []int, k, m int) int {
	hashes := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		hashes[arr[i]]++
	}
	ans := 0
	for key, val := range hashes {
		if val == k {
			ans = key
			break
		}
	}
	return ans
}

func initHashMap(hashes map[int]int) {
	value := 1
	for i := 0; i < 64; i++ {
		hashes[value] = i
		value <<= 1
	}
}

func randomArray(maxKinds, ranges, k, m int) []int {
	kTimesNum := randomNumber(ranges)
	// 真命天子出现的次数
	times := k
	numKinds := rand.Intn(maxKinds+1) + 3
	arr := make([]int, times+(numKinds-1)*m)
	index := 0
	for ; index < times; index++ {
		arr[index] = kTimesNum
	}
	numKinds--
	hashes := make(map[int]struct{})
	hashes[kTimesNum] = struct{}{}
	for numKinds != 0 {
		curNum := randomNumber(ranges)
		for {
			_, ok := hashes[curNum]
			if ok {
				curNum = randomNumber(ranges)
			} else {
				hashes[curNum] = struct{}{}
				break
			}
		}
		numKinds--
		for i := 0; i < m; i++ {
			arr[index] = curNum
			index++
		}
	}
	for i := 0; i < len(arr); i++ {
		j := rand.Intn(len(arr))
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// [-range, +range]
func randomNumber(ranges int) int {
	negative := rand.Intn(2)
	num := rand.Intn(ranges + 1)
	if negative == 0 {
		return num * -1
	}
	return num
}

func main() {
	c1 := []int{-1, -1, -2, -2, -1, -1, 3, 3, 3, 3, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64}
	fmt.Println(onlyKTimes(c1, 2, 4))
	fmt.Println(km(c1, 2, 4))
	fmt.Println(test(c1, 2, 4))

	rand.Seed(time.Now().UnixNano())
	maxKinds := 5
	ranges := 100
	testTimes := 100000
	max := 9
	for i := 0; i < testTimes; i++ {
		k := rand.Intn(max + 1)
		m := rand.Intn(max + 1)
		if k == m {
			m++
		}
		if k > m {
			k, m = m, k
		}
		arr := randomArray(maxKinds, ranges, k, m)
		ans1 := onlyKTimes(arr, k, m)
		ans2 := km(arr, k, m)
		ans3 := test(arr, k, m)
		if ans1 != ans2 || ans1 != ans3 {
			fmt.Println("Something wrong...")
			fmt.Println(arr)
			fmt.Println(ans1, ans2, ans3)
		}
	}
	fmt.Println("Nice algorithm...")
}
