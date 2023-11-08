package code01_FriendCircles

// https://leetcode.com/problems/friend-circles/

func findCircleNum(isConnected [][]int) int {
	N := len(isConnected)
	uf := NewUnionFind(N)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if isConnected[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	return uf.Sets
}

type UnionFind struct {
	Parent []int
	Size   []int
	Help   []int
	Sets   int
}

func NewUnionFind(N int) *UnionFind {
	uf := &UnionFind{
		Parent: make([]int, N),
		Size:   make([]int, N),
		Help:   make([]int, N),
		Sets:   N,
	}
	for i := 0; i < N; i++ {
		uf.Parent[i] = i
		uf.Size[i] = 1
	}
	return uf
}

func (uf *UnionFind) find(i int) int {
	hi := 0
	for i != uf.Parent[i] {
		uf.Help[hi] = i
		hi++
		i = uf.Parent[i]
	}
	hi--
	for hi >= 0 {
		uf.Parent[uf.Help[hi]] = i
		hi--
	}
	return i
}

func (uf *UnionFind) union(i, j int) {
	f1 := uf.find(i)
	f2 := uf.find(j)
	if f1 != f2 {
		if uf.Size[f1] > uf.Size[f2] {
			uf.Size[f1] += uf.Size[f2]
			uf.Parent[f2] = f1
		} else {
			uf.Size[f2] += uf.Size[f1]
			uf.Parent[f1] = f2
		}
		uf.Sets--
	}
}
