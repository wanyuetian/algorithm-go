package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func HeapSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	h := NewMyHeap(len(arr))
	//for i := 0; i < h.Limit; i++ {
	//	h.Data[i] = arr[i]
	//	h.heapInsert(i)
	//
	//}
	h.Size = len(h.Data)
	for i := h.Limit - 1; i >= 0; i-- {
		h.Data[i] = arr[i]
		h.heapify(i)
	}
	h.Size = len(h.Data)
	h.Data[0], h.Data[h.Size-1] = h.Data[h.Size-1], h.Data[0]
	h.Size--
	for h.Size > 0 {
		h.heapify(0)
		h.Data[0], h.Data[h.Size-1] = h.Data[h.Size-1], h.Data[0]
		h.Size--
	}
	return h.Data
}

type MyHeap struct {
	Data  []int
	Limit int
	Size  int
}

func NewMyHeap(limit int) *MyHeap {
	return &MyHeap{
		Data:  make([]int, limit),
		Limit: limit,
		Size:  0,
	}
}

func (h *MyHeap) IsEmpty() bool {
	return h.Size == 0
}

func (h *MyHeap) IsFull() bool {
	return h.Limit == h.Size
}

func (h *MyHeap) Push(value int) error {
	if h.IsFull() {
		return errors.New("heap is full")
	}
	index := h.Size
	h.Data[index] = value
	h.Size++
	return h.heapInsert(index)
}

func (h *MyHeap) Pop() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("heap is empty")
	}
	target := h.Data[0]
	h.Data[h.Size-1], h.Data[0] = h.Data[0], h.Data[h.Size-1]
	h.Size--
	var err error
	if h.Size != 0 {
		err = h.heapify(0)
	}
	return target, err
}

func (h *MyHeap) heapInsert(index int) error {
	for h.Data[index] > h.Data[(index-1)/2] {
		h.Data[index], h.Data[(index-1)/2] = h.Data[(index-1)/2], h.Data[index]
		index = (index - 1) / 2
	}
	return nil
}

func (h *MyHeap) heapify(index int) error {
	if index >= h.Size {
		return errors.New("index out of range")
	}
	left := 2*index + 1
	right := 2*index + 2
	var bigger int
	for {
		if right < h.Size {
			if h.Data[left] > h.Data[right] {
				bigger = left
			} else {
				bigger = right
			}
		} else if left < h.Size {
			bigger = left
		} else {
			break
		}
		if h.Data[bigger] > h.Data[index] {
			h.Data[bigger], h.Data[index] = h.Data[index], h.Data[bigger]
			index = bigger
		} else {
			break
		}

		left = 2*index + 1
		right = 2*index + 2
	}
	return nil
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

func generateRandomArray(maxSize, maxValue int) []int {
	maxSize = rand.Intn(maxSize + 1)
	arr := make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		arr[i] = rand.Intn(maxValue + 1)
	}
	return arr
}

func main() {
	rand.Seed(time.Now().UnixNano())
	testTimes := 100000
	maxSize := 100
	maxVal := 1000
	for i := 0; i < testTimes; i++ {
		arr1 := generateRandomArray(maxSize, maxVal)
		arr2 := copyArray(arr1)
		arr1 = HeapSort(arr1)
		sort.Ints(arr2)
		if !isEqual(arr1, arr2) {
			fmt.Println(arr1)
			fmt.Println(arr2)
			fmt.Println("something wrong")
			return
		}
	}
	fmt.Println("Nice algorithm")
}
