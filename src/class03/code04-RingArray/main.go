package main

import "errors"

type RingArray struct {
	arr       []int
	size      int
	Limit     int
	pushIndex int
	pollIndex int
}

func NewRingArray(limit int) *RingArray {
	return &RingArray{
		arr:   make([]int, limit, limit),
		Limit: limit,
	}
}

func (r *RingArray) Push(value int) error {
	if r.IsFull() {
		return errors.New("array is full")
	}

	r.arr[r.nextIndex(r.pushIndex)] = value
	r.pushIndex = r.nextIndex(r.pushIndex)
	r.size++
	return nil
}

func (r *RingArray) Poll() (int, error) {
	if r.IsEmpty() {
		return 0, errors.New("array is empty")
	}

	value := r.arr[r.nextIndex(r.pollIndex)]
	r.pollIndex = r.nextIndex(r.pollIndex)
	r.size--
	return value, nil
}

func (r *RingArray) IsEmpty() bool {
	return r.size == 0
}

func (r *RingArray) IsFull() bool {
	return r.size == r.Limit
}

func (r *RingArray) nextIndex(index int) int {
	if index < r.Limit {
		return index + 1
	}
	return 0
}
