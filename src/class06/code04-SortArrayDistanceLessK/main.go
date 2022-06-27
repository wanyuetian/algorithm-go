package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func SortArrayDistanceLessK(arr []int, k int) {
	if k == 0 {
		return
	}
	h := &Heap{}
	heap.Init(h)
	index := 0
	for ; index <= min(len(arr)-1, k-1); index++ {
		heap.Push(h, arr[index])
	}
	i := 0
	for ; index < len(arr); index++ {
		heap.Push(h, arr[index])
		arr[i] = heap.Pop(h).(int)
		i++
	}
	for !h.IsEmpty() {
		arr[i] = heap.Pop(h).(int)
		i++
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] <= (*h)[j]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	length := len(*h)
	res := (*h)[length-1]
	*h = (*h)[:length-1]
	return res
}

func (h *Heap) IsEmpty() bool {
	return len(*h) == 0
}

func copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	copyArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		copyArr[i] = arr[i]
	}
	return copyArr
}

func isEqual(arr1, arr2 []int) bool {
	if arr1 == nil && arr2 == nil {
		return true
	}
	if (arr1 == nil && arr2 != nil) || (arr1 != nil && arr2 == nil) {
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

func generateRandomArray(maxSize, maxValue, k int) []int {
	maxSize = rand.Intn(maxSize + 1)
	arr := make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		arr[i] = rand.Intn(maxValue + 1)
	}
	sort.Ints(arr)
	isSwap := make([]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		j := min(i+rand.Intn(k+1), len(arr)-1)
		if !isSwap[i] && !isSwap[j] {
			isSwap[i] = true
			isSwap[j] = true
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

func main() {
	rand.Seed(time.Now().UnixNano())
	testTimes := 100000
	maxSize := 5
	maxVal := 100

	for i := 0; i < testTimes; i++ {
		k := rand.Intn(maxSize + 1)
		//k = 1
		arr1 := generateRandomArray(maxSize, maxVal, k)
		//arr1 = []int{61, 96, 48}
		arr2 := copyArray(arr1)
		SortArrayDistanceLessK(arr1, k)
		sort.Ints(arr2)
		if !isEqual(arr1, arr2) {
			fmt.Println(k)
			fmt.Println(arr1)
			fmt.Println(arr2)
			fmt.Println("something wrong")
			return
		}
	}
	fmt.Println("Nice algorithm")
}
