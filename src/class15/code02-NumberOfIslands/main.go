package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://leetcode.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	islands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				islands++
				infect(grid, i, j)
			}
		}
	}
	return islands
}

func infect(grid [][]byte, i, j int) {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = 0
	infect(grid, i-1, j)
	infect(grid, i+1, j)
	infect(grid, i, j-1)
	infect(grid, i, j+1)
}

type Dot struct{}

type UnionFind1[T comparable] struct {
	Nodes   map[*T]*Node[T]
	Parents map[*Node[T]]*Node[T]
	SizeMap map[*Node[T]]int
}

type Node[T comparable] struct {
	Value *T
}

func NewUnionFind1[T comparable](values []*T) *UnionFind1[T] {
	nodes := make(map[*T]*Node[T])
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
	return &UnionFind1[T]{
		Nodes:   nodes,
		Parents: parents,
		SizeMap: sizeMap,
	}
}

func (u *UnionFind1[T]) findFather(cur *Node[T]) *Node[T] {
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

func (u *UnionFind1[T]) IsSameSet(a *Node[T], b *Node[T]) bool {
	return u.findFather(a) == u.findFather(b)
}

func (u *UnionFind1[T]) Union(a *T, b *T) {
	aHead := u.findFather(u.Nodes[a])
	bHead := u.findFather(u.Nodes[b])
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

func (u *UnionFind1[T]) Sets() int {
	return len(u.SizeMap)
}

func numIslands1(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	dots := make([][]*Dot, row)
	dotsList := make([]*Dot, 0)
	for i := 0; i < row; i++ {
		dots[i] = make([]*Dot, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				dots[i][j] = &Dot{}
				dotsList = append(dotsList, dots[i][j])
			}
		}
	}
	uf := NewUnionFind1(dotsList)
	for i := 1; i < row; i++ {
		if grid[i-1][0] == '1' {
			uf.Union(dots[i-1][0], dots[i][0])
		}
	}
	for i := 1; i < col; i++ {
		if grid[0][i-1] == '1' {
			uf.Union(dots[0][i-1], dots[0][i])
		}
	}
	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				if grid[i-1][j] == '1' {
					uf.Union(dots[i][j], dots[i-1][j])
				}
				if grid[i][j-1] == '1' {
					uf.Union(dots[i][j], dots[i][j-1])
				}
			}
		}
	}
	return uf.Sets()
}

func generateRandomMatrix(row, col int) [][]byte {
	grid := make([][]byte, row)
	for i := 0; i < row; i++ {
		grid[i] = make([]byte, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if rand.Float64() > 0.5 {
				grid[i][j] = '1'
			} else {
				grid[i][j] = '0'
			}
		}
	}
	return grid
}

func copyRandomMatrix(grid [][]byte) [][]byte {
	if grid == nil || len(grid) == 0 {
		return nil
	}
	row := len(grid)
	col := len(grid[0])
	copyGrid := make([][]byte, row)
	for i := 0; i < row; i++ {
		copyGrid[i] = make([]byte, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			copyGrid[i][j] = grid[i][j]
		}
	}
	return copyGrid
}

func main() {
	row := 10000
	col := 10000
	var grid, grid1 [][]byte
	grid = generateRandomMatrix(row, col)
	grid1 = copyRandomMatrix(grid)

	start := time.Now().UnixMilli()
	fmt.Printf("感染方法的运行结果: %d\n", numIslands(grid))
	end := time.Now().UnixMilli()
	fmt.Printf("感染方法的运行时间: %d ms\n", end-start)

	start = time.Now().UnixMilli()
	fmt.Printf("并查集(map实现)的运行结果: %d\n", numIslands(grid1))
	end = time.Now().UnixMilli()
	fmt.Printf("并查集(map实现)的运行时间: %d ms\n", end-start)

}
