package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Line struct {
	Start int
	End   int
}

func NewLine(start, end int) *Line {
	return &Line{
		Start: start,
		End:   end,
	}
}

type Lines []*Line

func NewLines(size int) Lines {
	return make([]*Line, size)
}

func (l Lines) Len() int {
	return len(l)
}

func (l Lines) Less(i, j int) bool {
	return l[i].Start <= l[j].Start
}

func (l Lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l Lines) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	l = append(l, x.(*Line))
}

func (l Lines) Pop() interface{} {
	length := len(l)
	res := l[length-1]
	tmp := &l
	*tmp = (*tmp)[:length-1]
	return res
}

func (l Lines) IsEmpty() bool {
	return len(l) == 0
}

func (l Lines) Peek() *Line {
	if !l.IsEmpty() {
		return l[len(l)-1]
	}
	return nil
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a >= b {
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

func (h *Heap) Peek() int {
	return (*h)[0]
}

func (h *Heap) IsEmpty() bool {
	return len(*h) == 0
}

func maxCover1(lines [][2]int) int {
	mostLeft, mostRight := math.MaxInt, math.MinInt

	for i := 0; i < len(lines); i++ {
		mostLeft = Min(mostLeft, lines[i][0])
		mostRight = Max(mostRight, lines[i][1])
	}

	cover := 0

	for start := float64(mostLeft) + 0.5; start < float64(mostRight); start++ {
		cur := 0
		for i := 0; i < len(lines); i++ {

			if start > float64(lines[i][0]) && start < float64(lines[i][1]) {
				cur++
			}
		}
		cover = Max(cover, cur)
	}

	return cover
}

func maxCover2(m [][2]int) int {
	lines := NewLines(len(m))
	for i := 0; i < len(m); i++ {
		lines[i] = NewLine(m[i][0], m[i][1])
	}
	sort.Sort(lines)
	h := &Heap{}
	heap.Init(h)
	max := 0
	for i := 0; i < len(lines); i++ {
		for !h.IsEmpty() && h.Peek() <= lines[i].Start {
			heap.Pop(h)
		}
		heap.Push(h, lines[i].End)
		max = Max(max, h.Len())
	}
	return max
}

func generateLines(N, L, R int) [][2]int {
	size := rand.Intn(N + 1)
	res := make([][2]int, size)

	for i := 0; i < size; i++ {
		a := L + int(rand.Float64()*float64(R-L+1))
		b := L + int(rand.Float64()*float64(R-L+1))
		if a == b {
			b = a + 1
		}
		res[i][0] = Min(a, b)
		res[i][1] = Max(a, b)
	}
	return res
}

func main() {
	rand.Seed(time.Now().UnixNano())
	N := 10
	L := 0
	R := 200
	testTimes := 100000

	for i := 0; i < testTimes; i++ {
		lines := generateLines(N, L, R)
		//lines = [][2]int{
		//	{68, 79},
		//	{120, 176},
		//	{41, 158},
		//	{21, 124},
		//	{42, 144},
		//	{89, 129},
		//}
		ans1 := maxCover1(lines)
		ans2 := maxCover2(lines)
		if ans1 != ans2 {
			fmt.Println(lines)
			fmt.Println("Oops")
			//return
		}
	}
	fmt.Println("Nice")
}
