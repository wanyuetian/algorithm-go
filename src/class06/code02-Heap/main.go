package main

import (
	"errors"
	"fmt"
)

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

func main() {
	hp := NewMyHeap(10)
	v, err := hp.Pop()
	fmt.Println(v, err)
	err = hp.Push(8)
	fmt.Println(err)
	err = hp.Push(10)
	fmt.Println(err)
	err = hp.Push(6)
	fmt.Println(err)
	err = hp.Push(6)
	fmt.Println(err)
	err = hp.Push(16)
	fmt.Println(err)
	err = hp.Push(36)
	fmt.Println(err)
	err = hp.Push(36)
	fmt.Println(err)
	err = hp.Push(36)
	fmt.Println(err)
	err = hp.Push(136)
	fmt.Println(err)
	err = hp.Push(236)
	fmt.Println(err)
	err = hp.Push(236)
	fmt.Println(err)
	err = hp.Push(16)
	fmt.Println(err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
	v, err = hp.Pop()
	fmt.Println(v, err)
}
