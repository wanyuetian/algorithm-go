package main

import (
	"errors"
	"fmt"
)

// https://segmentfault.com/a/1190000041634906  Go 1.18 泛型全面讲解：一篇讲清泛型的全部

type GreaterHeap[T comparable] struct {
	Data       []T
	Limit      int
	Size       int
	Map        map[T]int
	Comparator func(a, b T) int
}

func NewGreaterHeap[T comparable](comparator func(a, b T) int, limit int) *GreaterHeap[T] {
	return &GreaterHeap[T]{
		Data:       make([]T, limit),
		Limit:      limit,
		Size:       0,
		Comparator: comparator,
		Map:        make(map[T]int),
	}
}

func (h *GreaterHeap[T]) Contains(obj T) bool {
	_, ok := h.Map[obj]
	return ok
}

func (h *GreaterHeap[T]) IsEmpty() bool {
	return h.Size == 0
}

func (h *GreaterHeap[T]) IsFull() bool {
	return h.Limit == h.Size
}

func (h *GreaterHeap[T]) Push(obj T) error {
	if h.IsFull() {
		return errors.New("heap is full")
	}
	index := h.Size
	h.Data[index] = obj
	h.Map[obj] = index
	h.Size++
	return h.heapInsert(index)
}

func (h *GreaterHeap[T]) Pop() (T, error) {
	var target T
	if h.IsEmpty() {
		return target, errors.New("heap is empty")
	}
	target = h.Data[0]

	//h.Data[h.Size-1], h.Data[0] = h.Data[0], h.Data[h.Size-1]
	h.swap(h.Size-1, 0)
	h.Size--
	delete(h.Map, target)
	var err error
	if h.Size != 0 {
		err = h.heapify(0)
	}
	return target, err
}

func (h *GreaterHeap[T]) Remove(obj T) {
	var replace T
	replace = h.Data[h.Size-1]
	index, ok := h.Map[obj]
	if !ok {
		return
	}
	h.Data = append(h.Data[:index], h.Data[index+1:]...)
	h.Size--
	if obj != replace {
		h.Data[index] = replace
		h.Map[replace] = index
		h.Resign(replace)
	}
}

func (h *GreaterHeap[T]) Peek() (T, error) {
	var target T
	if h.IsEmpty() {
		return target, errors.New("heap is empty")
	}
	return h.Data[0], nil
}

func (h *GreaterHeap[T]) heapInsert(index int) error {
	for h.Comparator(h.Data[index], h.Data[(index-1)/2]) > 0 { // h.Data[index] > h.Data[(index-1)/2]
		//h.Data[index], h.Data[(index-1)/2] = h.Data[(index-1)/2], h.Data[index]
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
	return nil
}

func (h *GreaterHeap[T]) heapify(index int) error {
	if index >= h.Size {
		return errors.New("index out of range")
	}
	left := 2*index + 1
	right := 2*index + 2
	var bigger int
	for {
		if right < h.Size {
			if h.Comparator(h.Data[left], h.Data[right]) > 0 { // h.Data[left] > h.Data[right]
				bigger = left
			} else {
				bigger = right
			}
		} else if left < h.Size {
			bigger = left
		} else {
			break
		}
		if h.Comparator(h.Data[bigger], h.Data[index]) > 0 { // h.Data[bigger] > h.Data[index]
			//h.Data[bigger], h.Data[index] = h.Data[index], h.Data[bigger]
			h.swap(bigger, index)
			index = bigger
		} else {
			break
		}
		left = 2*index + 1
		right = 2*index + 2
	}
	return nil
}

func (h *GreaterHeap[T]) Resign(obj T) {
	index, ok := h.Map[obj]
	if ok {
		h.heapInsert(index)
		h.heapify(index)
	}
}

func (h *GreaterHeap[T]) swap(i, j int) {
	o1 := h.Data[i]
	o2 := h.Data[j]
	h.Data[i] = o2
	h.Data[j] = o1
	h.Map[o1] = j
	h.Map[o2] = i
}

func main() {
	hp := NewGreaterHeap[int](func(a, b int) int {
		return a - b
	}, 10)
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

	hp1 := NewGreaterHeap[float64](func(a, b float64) int {
		if a > b {
			return 1
		} else if a == b {
			return 0
		}
		return -1
	}, 10)
	v1, err := hp1.Pop()
	fmt.Println(v1, err)
	err = hp1.Push(8.0)
	fmt.Println(err)
	err = hp1.Push(10.0)
	fmt.Println(err)
	err = hp1.Push(6.1)
	fmt.Println(err)
	err = hp1.Push(6.3)
	fmt.Println(err)
	err = hp1.Push(16.32)
	v1, err = hp1.Pop()
	fmt.Println(v1, err)
	v1, err = hp1.Pop()
	fmt.Println(v1, err)
	v1, err = hp1.Pop()
	fmt.Println(v1, err)

	type Person struct {
		Name string
		Age  int
	}
	hp2 := NewGreaterHeap(func(a, b Person) int {
		return a.Age - b.Age
	}, 10)
	v2, err := hp2.Pop()
	fmt.Println(v2, err)
	err = hp2.Push(Person{
		Name: "zhangfei",
		Age:  18,
	})
	fmt.Println(err)
	err = hp2.Push(Person{
		Name: "liubei",
		Age:  22,
	})
	fmt.Println(err)
	fmt.Println(err)
	err = hp2.Push(Person{
		Name: "guanyu",
		Age:  20,
	})
	fmt.Println(err)

	v2, err = hp2.Pop()
	fmt.Println(v2, err)
	v2, err = hp2.Pop()
	fmt.Println(v2, err)
	v2, err = hp2.Pop()
	fmt.Println(v2, err)
}
