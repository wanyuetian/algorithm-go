package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// arr中，只有一种数，出现奇数次
func findOddTimesNum1(arr []int) (int, error) {
	if arr == nil || len(arr) < 1 {
		return -1, errors.New("invalid arr")
	}
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor = eor ^ arr[i]
	}
	return eor, nil
}

// arr中，有两种数，出现奇数次
func findOddTimesNum2(arr []int) (int, int, error) {
	if arr == nil || len(arr) < 2 {
		return -1, -1, errors.New("invalid arr")
	}
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor = eor ^ arr[i]
	}
	rightOne := eor & (-eor)
	onlyOne := eor
	for i := 0; i < len(arr); i++ {
		if arr[i]&rightOne != 0 {
			onlyOne = onlyOne ^ arr[i]
		}
	}
	return onlyOne, onlyOne ^ eor, nil
}

func test1(arr []int) (int, error) {
	if arr == nil || len(arr) < 1 {
		return -1, errors.New("invalid arr")
	}
	cache := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := cache[arr[i]]; ok {
			cache[arr[i]]++
		} else {
			cache[arr[i]] = 1
		}
	}
	for k, v := range cache {
		if v%2 != 0 {
			return k, nil
		}
	}
	return -1, errors.New("invalid arr")
}

func test2(arr []int) (int, int, error) {
	if arr == nil || len(arr) < 2 {
		return -1, -1, errors.New("invalid arr")
	}
	cache := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := cache[arr[i]]; ok {
			cache[arr[i]]++
		} else {
			cache[arr[i]] = 1
		}
	}
	ans := make([]int, 0)
	for k, v := range cache {
		if v%2 != 0 {
			ans = append(ans, k)
			if len(ans) == 2 {
				return ans[0], ans[1], nil
			}
		}
	}
	return -1, -1, errors.New("invalid arr")
}

func generateRandomArray1(maxValue, maxTimes, maxKinds int) []int {
	arr := make([]int, 0)
	target1 := rand.Intn(maxValue + 1)

	mKinds := rand.Intn(maxKinds + 1)
	kTimes := rand.Intn(maxTimes) + 1
	for kTimes == 0 || kTimes%2 == 0 {
		kTimes = rand.Intn(maxTimes) + 1
	}
	mNums := make([]int, 0)
	for i := 0; i < mKinds; i++ {
		num := rand.Intn(maxValue)
		for num == target1 || contains(mNums, num) {
			num = rand.Intn(maxValue)
		}
		mNums = append(mNums, num)
	}
	for i := 0; i < len(mNums); i++ {
		if rand.Float32() > 0.5 {
			mNums[i] *= -1
		}
	}
	if rand.Float32() > 0.5 {
		target1 *= -1
	}

	for i := 0; i < kTimes; i++ {
		arr = append(arr, target1)
	}

	for i := 0; i < mKinds; i++ {
		mTimes := rand.Intn(maxTimes) + 2
		for mTimes%2 != 0 {
			mTimes = rand.Intn(maxTimes) + 2
		}
		for j := 0; j < mTimes; j++ {
			arr = append(arr, mNums[i])
		}
	}
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })

	return arr
}

func contains(s []int, value int) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == value {
			return true
		}
	}
	return false
}

func generateRandomArray2(maxValue, maxTimes, maxKinds int) []int {
	arr := make([]int, 0)
	target1 := rand.Intn(maxValue + 1)
	target2 := rand.Intn(maxValue + 1)
	for target2 == target1 {
		target2 = rand.Intn(maxValue + 1)
	}
	mKinds := rand.Intn(maxKinds + 1)
	kTimes := rand.Intn(maxTimes) + 1
	for kTimes == 0 || kTimes%2 == 0 {
		kTimes = rand.Intn(maxTimes) + 1
	}
	mNums := make([]int, 0)
	for i := 0; i < mKinds; i++ {
		num := rand.Intn(maxValue)
		for num == target1 || num == target2 || contains(mNums, num) {
			num = rand.Intn(maxValue)
		}
		mNums = append(mNums, num)
	}
	for i := 0; i < len(mNums); i++ {
		if rand.Float32() > 0.5 {
			mNums[i] *= -1
		}
	}
	if rand.Float32() > 0.5 {
		target1 *= -1
	}
	if rand.Float32() > 0.5 {
		target2 *= -1
	}
	for i := 0; i < kTimes; i++ {
		arr = append(arr, target1)
	}
	for i := 0; i < kTimes; i++ {
		arr = append(arr, target2)
	}
	for i := 0; i < mKinds; i++ {
		mTimes := rand.Intn(maxTimes) + 2
		for mTimes%2 != 0 {
			mTimes = rand.Intn(maxTimes) + 2
		}
		for j := 0; j < mTimes; j++ {
			arr = append(arr, mNums[i])
		}
	}
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })

	return arr
}

func main() {
	rand.Seed(time.Now().UnixNano())
	testTimes := 100000
	maxValue := 100
	maxTimes := 4
	maxKinds := 4
	for i := 0; i < testTimes; i++ {
		arr1 := generateRandomArray1(maxValue, maxTimes, maxKinds)
		arr2 := generateRandomArray2(maxValue, maxTimes, maxKinds)
		t1, _ := test1(arr1)
		t2, t3, _ := test2(arr2)
		ans1, _ := findOddTimesNum1(arr1)
		ans2, ans3, _ := findOddTimesNum2(arr2)
		if t1 != ans1 {
			fmt.Println("findOddTimesNum1 wrong")
			return
		}
		if !((t2 == ans2 && t3 == ans3) || (t2 == ans3 && t3 == ans2)) {
			fmt.Println("findOddTimesNum2 wrong")
			return
		}

	}
	fmt.Println("Nice")
}
