package main

import "fmt"

/*

 */

type UnionFind[T comparable] struct {
	Nodes   map[T]*Node[T]
	Parents map[*Node[T]]*Node[T]
	SizeMap map[*Node[T]]int
}

type Node[T comparable] struct {
	Value T
}

func NewUnionFind[T comparable](values []T) *UnionFind[T] {
	nodes := make(map[T]*Node[T])
	parents := make(map[*Node[T]]*Node[T])
	sizeMap := make(map[*Node[T]]int)
	for i := range values {
		node := &Node[T]{
			Value: values[i],
		}
		nodes[values[i]] = node
		parents[node] = node
		sizeMap[node] = 1
	}
	return &UnionFind[T]{
		Nodes:   nodes,
		Parents: parents,
		SizeMap: sizeMap,
	}
}

func (u *UnionFind[T]) findFather(cur *Node[T]) *Node[T] {
	path := make([]*Node[T], 0)
	for cur != u.Parents[cur] {
		path = append(path, cur)
		cur = u.Parents[cur]
	}
	for len(path) != 0 {
		node := path[0]
		path = path[1:]
		u.Parents[node] = cur
	}
	return cur
}

func (u *UnionFind[T]) IsSameSet(a *Node[T], b *Node[T]) bool {
	return u.findFather(a) == u.findFather(b)
}

func (u *UnionFind[T]) Union(a *Node[T], b *Node[T]) {
	aHead := u.findFather(a)
	bHead := u.findFather(b)
	if aHead != bHead {
		aSetSize := u.SizeMap[aHead]
		bSetSize := u.SizeMap[bHead]
		var big, small *Node[T]
		if aSetSize >= bSetSize {
			big = aHead
			small = bHead
		} else {
			big = bHead
			small = aHead
		}
		u.Parents[small] = big
		u.SizeMap[big] = aSetSize + bSetSize
		delete(u.SizeMap, small)
	}
}

func (u *UnionFind[T]) Sets() int {
	return len(u.SizeMap)
}

func main() {
	u := NewUnionFind[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	u.Union(&Node[int]{Value: 1}, &Node[int]{Value: 2})
	fmt.Println(u.Sets())
}
